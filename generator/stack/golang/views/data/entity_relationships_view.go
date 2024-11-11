package data

/*
   - name: comments
     type: $rel # $rel or $urn
     ref: ".torpedo/entities/comment.yaml"
     cardinality: hasMany # hasOne, hasMany
     eagerLoading:
       maxItems: 100
*/

type EntityRelationship struct {
	Name          string
	Type          RelationshipTypeEnum
	Ref           *EntityView
	Cardinality   CardinalityTypeEnum
	NestedLoading *EntityRelationshipNestedLoading
}

func (e *EntityRelationship) HasNestedLoading() bool {
	return e.NestedLoading != nil
}

type EntityRelationshipNestedLoading struct {
	MaxItems int
}
