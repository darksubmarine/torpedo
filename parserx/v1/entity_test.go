package v1_test

import (
	_ "embed"
	"github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

//go:embed post.entity.test.yaml
var postYaml []byte

func TestEntity(t *testing.T) {

	var myEntity v1.RootEntity
	assert.Nil(t, yaml.Unmarshal(postYaml, &myEntity))

	//-- Entity data --

	assert.Equal(t, myEntity.Kind, vx.KindEntity)
	assert.Equal(t, myEntity.Entity.Name, "post")
	assert.Equal(t, myEntity.Entity.Plural, "posts")
	assert.Equal(t, myEntity.Entity.Description, "The blog post entity")
	assert.Equal(t, myEntity.Entity.Doc, ".torpedo/entities/docs/post.md")

	//-- Entity schema --

	assert.Equal(t, myEntity.Entity.Schema.Reserved.Id.Type, v1.ULID)
	assert.Len(t, myEntity.Entity.Schema.Fields, 5)

	assert.Equal(t, myEntity.Entity.Schema.Fields[0].Name, "title")
	assert.Equal(t, myEntity.Entity.Schema.Fields[0].Type, v1.String)
	assert.Equal(t, myEntity.Entity.Schema.Fields[0].Description, "The post title")
	assert.Equal(t, myEntity.Entity.Schema.Fields[0].Doc, "The post title to be used as part of the post page.\nAlso as part of the SEO indexes.\n")

	// assert encrypted field
	assert.True(t, myEntity.Entity.Schema.Fields[1].Encrypted)

	// assert optional field
	assert.EqualValues(t, false, myEntity.Entity.Schema.Fields[3].Optional.Default)

	//-- Entity Relationships --
	assert.Len(t, myEntity.Entity.Relationships, 2)

	assert.Equal(t, "comments", myEntity.Entity.Relationships[0].Name)
	assert.Equal(t, "$rel", myEntity.Entity.Relationships[0].Type)
	assert.Equal(t, ".torpedo/entities/comment.yaml", myEntity.Entity.Relationships[0].Ref)
	assert.Equal(t, "hasMany", myEntity.Entity.Relationships[0].Cardinality)
	assert.Equal(t, "nested", myEntity.Entity.Relationships[0].Load.Type)
	assert.EqualValues(t, 100, myEntity.Entity.Relationships[0].Load.Metadata["maxItems"])

	assert.Equal(t, "authors", myEntity.Entity.Relationships[1].Name)
	assert.Empty(t, myEntity.Entity.Relationships[1].Type) // By default, should be $rel
	assert.Equal(t, ".torpedo/entities/author.yaml", myEntity.Entity.Relationships[1].Ref)
	assert.Equal(t, "hasMany", myEntity.Entity.Relationships[1].Cardinality)

	//-- Entity Adapters --
	assert.Len(t, myEntity.Entity.Adapters.Input, 1)
	assert.Len(t, myEntity.Entity.Adapters.Output, 6)

	assert.Equal(t, "http", myEntity.Entity.Adapters.Input[0].Type)
	assert.EqualValues(t, "works-post", myEntity.Entity.Adapters.Input[0].Metadata["resourceName"])

	assert.Equal(t, "memory", myEntity.Entity.Adapters.Output[0].Type)

	assert.Equal(t, "mongodb", myEntity.Entity.Adapters.Output[1].Type)
	assert.EqualValues(t, "posts", myEntity.Entity.Adapters.Output[1].Metadata["collection"])

	assert.Equal(t, "redis", myEntity.Entity.Adapters.Output[2].Type)
	assert.EqualValues(t, 30000, myEntity.Entity.Adapters.Output[2].Metadata["ttl"])

	assert.Equal(t, "sql", myEntity.Entity.Adapters.Output[3].Type)
	assert.EqualValues(t, "posts", myEntity.Entity.Adapters.Output[3].Metadata["table"])

	assert.Equal(t, "redis+mongodb", myEntity.Entity.Adapters.Output[4].Type)

	assert.Equal(t, "redis+sql", myEntity.Entity.Adapters.Output[5].Type)

}
