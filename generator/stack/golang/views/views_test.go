package views_test

import (
	"fmt"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views/data"
	"testing"
)

func TestEntityBase_View(t *testing.T) {
	t.Skip()
	meta := data.EntityViewMeta{
		Package:    "bitbucket.org/darksubmarine/torpedo/build",
		EntityPath: "/domain/entities",
	}

	commentsView := &data.EntityView{
		Meta: meta,

		Name:        "comment",
		PluralName:  "comments",
		Description: "The blog post comments entity",
		Schema: data.EntitySchema{
			Reserved: data.EntitySchemaReserved{
				Id: data.EntitySchemaReservedId{
					Type: "uuid",
				},
			},
			Fields: []data.EntitySchemaItem{
				{Name: "message", Type: data.String, LongDescription: "The comment message", Description: "The comment message"},
				{Name: "pubDate", Type: data.Date, LongDescription: "The comment date", Description: "The comment date"},
				{Name: "line", Type: data.Integer, LongDescription: "The comment line", Description: "The comment line",
					Optional: &data.OptionalField{Default: -1}},

				// Relationship post->hasMany->comments = comment.postId
				{Name: "postId", Type: data.String, LongDescription: "The postId relationship", Description: "The postId relationship"},
			},
		},
		Adapters: data.EntityAdapters{
			Input: data.InputAdapters{
				Http: &data.HttpAdapter{
					ResourceName: "comments",
				},
			},
			Output: data.OutputAdapters{
				Memory:  &data.MemoryAdapter{},
				MongoDB: &data.MongoDBAdapter{CollectionName: "comments"}, // TODO set as default value!
				Redis:   &data.RedisAdapter{TTL: -1},                      // TODO Set as default value
				Sql:     &data.SqlAdapter{TableName: "comments"},          // TODO set as default value!
			},
		},
	}

	postView := &data.EntityView{
		Meta:        meta,
		Name:        "post",
		PluralName:  "posts",
		Description: "The blog post entity",
		Schema: data.EntitySchema{
			Reserved: data.EntitySchemaReserved{
				Id: data.EntitySchemaReservedId{
					Type: "uuid",
				},
			},
			Fields: []data.EntitySchemaItem{
				{Name: "title", Type: data.String, LongDescription: "The post title", Description: "The post title",
					//Validator: &data.ItemValidatorRegex{GoPattern: "^4[0-9]{12}(?:[0-9]{3})?$"},
					//Validator: &data.ItemValidatorList{Type: data.String, List: []interface{}{"uno", "dos", "tres"}},
					//Validator: &data.ItemValidatorRange{Type: data.String, Min: "uno", Max: "dos"},
					Validator: &data.ItemValidatorValue{Type: data.String, Value: "uno"},
				},
				{Name: "pubDate", Type: data.Date, LongDescription: "The publication date", Description: "The publication date"},
				{Name: "published", Type: data.Boolean, LongDescription: "Sets the post to be publish", Description: "Sets the post to be publish"},
				{Name: "author", Type: data.String, LongDescription: "The post author", Description: "The post author"},
				{Name: "superSecret", Type: data.String, Encrypted: true, LongDescription: "The post secret", Description: "The post secret"},
			},
		},
		Adapters: data.EntityAdapters{
			Input: data.InputAdapters{
				Http: &data.HttpAdapter{
					ResourceName: "works-post",
					//FieldMap:     map[string]string{"published": "published-f"},
				},
			},
			Output: data.OutputAdapters{
				Memory:  &data.MemoryAdapter{},
				MongoDB: &data.MongoDBAdapter{CollectionName: "posts"},
				Redis:   &data.RedisAdapter{TTL: 3000},
				Sql:     &data.SqlAdapter{TableName: "posts"},
			},
		},
	}

	// Relationships

	postView.Relationships = map[string]data.EntityRelationship{
		"comments": {
			Name:          "comments",
			Type:          data.Rel,
			Ref:           commentsView,
			Cardinality:   data.HasMany,
			NestedLoading: &data.EntityRelationshipNestedLoading{MaxItems: 100},
		},
	}

	commentsView.Relationships = map[string]data.EntityRelationship{
		"posts": {
			Name:        "posts",
			Type:        data.Rel,
			Ref:         postView,
			Cardinality: data.BelongsTo,
		},
	}

	// Domain
	domainView := &data.DomainView{
		Package: "bitbucket.org/darksubmarine/torpedo/build",
		Path:    "/domain",
	}
	domainView.Entities = map[string]*data.EntityView{"comment": commentsView, "post": postView}

	//buf, err := views.RenderTpl("test", postView, views.ViewEntityBase)
	buf, err := views.RenderTpl("test", postView, views.TplInputDTOBase)
	if err != nil {
		t.Error(err)
	}

	if data, err := views.FormatCode(buf); err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(data))
	}

}
