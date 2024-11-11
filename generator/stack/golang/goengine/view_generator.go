package goengine

import (
	"fmt"
	"github.com/darksubmarine/torpedo/generator/stack/golang/views/data"
	v1 "github.com/darksubmarine/torpedo/parserx/v1"
	"github.com/darksubmarine/torpedo/parserx/vx"
	"os"
	"path"
)

type IEntityViewGenerator interface {
	GenerateEntityView() (*data.EntityView, error)
	HasRelationships() bool
	HydrateRelationships(views map[string]IEntityViewGenerator)
	EntityView() *data.EntityView
	AddRelationshipField(fieldId data.EntitySchemaItem)
}

func NewEntityViewGenerator(version vx.V, data interface{}, opts ViewOptions) (IEntityViewGenerator, error) {
	switch version {
	case vx.V1:
		if entity, ok := data.(v1.RootEntity); ok {
			return NewEntityViewGeneratorV1(entity.Entity, opts)
		}
		return nil, ErrInvalidEntityProvidedVersion
	default:
		return nil, ErrEntityProvidedVersionNotSupported
	}
}

type ViewOptions struct {
	AppPackage string
	EntityPath string
	ProjectDir string
}

type EntityViewGeneratorV1 struct {
	options ViewOptions
	entity  v1.EntitySpec
	view    *data.EntityView
}

func NewEntityViewGeneratorV1(entity v1.EntitySpec, opts ViewOptions) (*EntityViewGeneratorV1, error) {
	vg := &EntityViewGeneratorV1{entity: entity, options: opts}
	_, err := vg.GenerateEntityView()
	return vg, err
}

func (v *EntityViewGeneratorV1) EntityView() *data.EntityView {
	return v.view
}

func (v *EntityViewGeneratorV1) HasRelationships() bool {
	return len(v.entity.Relationships) > 0
}

func (v *EntityViewGeneratorV1) AddRelationshipField(fieldId data.EntitySchemaItem) {
	v.view.Schema.Fields = append(v.view.Schema.Fields, fieldId)
}

func (v *EntityViewGeneratorV1) HydrateRelationships(viewGenerators map[string]IEntityViewGenerator) {
	for _, rel := range v.entity.Relationships {
		if rel.Type == v1.Rel && rel.Cardinality == v1.HasMany {
			_, fileName := path.Split(rel.Ref)                     // reference file
			refEntityView := viewGenerators[fileName].EntityView() // reference view

			if _, exists := v.view.Relationships[refEntityView.PluralName]; exists {
				continue
			}

			// Adds rel id. Relationship currentView->hasMany->relView = relView.currentViewId
			viewGenerators[fileName].AddRelationshipField(
				data.EntitySchemaItem{
					Name:            fmt.Sprintf("%sId", v.view.Name),
					Type:            data.String,
					LongDescription: fmt.Sprintf("The %sId relationship", v.view.Name),
					Description:     fmt.Sprintf("The %sId relationship", v.view.Name)})

			var nestedLoading *data.EntityRelationshipNestedLoading = nil
			if rel.Load != nil && rel.Load.Type == v1.Nested {
				maxItems := 100
				if iMaxItems, ok := rel.Load.Metadata[v1.MaxItems]; ok {
					if mi, ok := iMaxItems.(int); ok {
						maxItems = mi
					}
				}
				nestedLoading = &data.EntityRelationshipNestedLoading{MaxItems: maxItems}
			}

			v.view.Relationships[refEntityView.PluralName] = data.EntityRelationship{
				Name:          refEntityView.PluralName,
				Type:          data.Rel,
				Cardinality:   data.HasMany,
				Ref:           refEntityView,
				NestedLoading: nestedLoading,
			}

			refEntityView.Relationships[v.view.PluralName] = data.EntityRelationship{
				Name:        v.view.PluralName,
				Type:        data.Rel,
				Ref:         v.view,
				Cardinality: data.BelongsTo,
			}
		} else if rel.Type == v1.Rel && rel.Cardinality == v1.HasOne {
			_, fileName := path.Split(rel.Ref)                     // reference file
			refEntityView := viewGenerators[fileName].EntityView() // reference view

			if _, exists := v.view.Relationships[refEntityView.PluralName]; exists {
				continue
			}

			// Adds rel id. Relationship currentView->hasOne->refView
			v.AddRelationshipField(
				data.EntitySchemaItem{
					Name:            fmt.Sprintf("%sId", refEntityView.Name),
					Type:            data.String,
					LongDescription: fmt.Sprintf("The %sId relationship", refEntityView.Name),
					Description:     fmt.Sprintf("The %sId relationship", refEntityView.Name),
				})

			var nestedLoading *data.EntityRelationshipNestedLoading = nil
			if rel.Load != nil && rel.Load.Type == v1.Nested {
				nestedLoading = &data.EntityRelationshipNestedLoading{MaxItems: 1} //HasOne
			}

			v.view.Relationships[refEntityView.PluralName] = data.EntityRelationship{
				Name:          refEntityView.PluralName,
				Type:          data.Rel,
				Cardinality:   data.HasOne,
				Ref:           refEntityView,
				NestedLoading: nestedLoading,
			}

		}
	}
}

func (v *EntityViewGeneratorV1) GenerateEntityView() (*data.EntityView, error) {

	meta := data.EntityViewMeta{
		Package:    v.options.AppPackage,
		EntityPath: v.options.EntityPath,
	}

	// Add docs
	var documentation string
	if v.entity.Doc != "" {
		documentationPath := path.Join(v.options.ProjectDir, TorpedoDir, TorpedoEntitiesDir, TorpedoDocsDir, v.entity.Doc)
		if _, err := os.Stat(documentationPath); err == nil {
			if docs, err := os.ReadFile(documentationPath); err == nil {
				documentation = string(docs)
			} else {
				return nil, err
			}
		} else {
			documentation = v.entity.Doc
		}
	}

	v.view = &data.EntityView{
		Meta: meta,

		Name:        v.entity.Name,
		PluralName:  v.entity.Plural,
		Description: v.entity.Description,
		Docs:        documentation,

		Schema: data.EntitySchema{
			Reserved: data.EntitySchemaReserved{
				Id: data.EntitySchemaReservedId{
					Type: v.entity.Schema.Reserved.Id.Type,
				},
			},
			Fields: []data.EntitySchemaItem{},
		},
		Relationships: map[string]data.EntityRelationship{},
		Adapters:      data.EntityAdapters{},
	}

	// Add docs
	//if v.entity.Doc != "" {
	//	if docs, err := os.ReadFile(v.entity.Doc); err == nil {
	//		v.view.Docs = string(docs)
	//	} else {
	//		return nil, err
	//	}
	//}

	// Add fields
	for _, item := range v.entity.Schema.Fields {
		fieldType := data.NewDataTypeEnumFromString(item.Type)
		toAdd := data.EntitySchemaItem{
			Name:            item.Name,
			Type:            fieldType,
			Description:     item.Description,
			LongDescription: item.Doc,
			Encrypted:       item.Encrypted,
			ReadOnly:        item.ReadOnly,
		}

		if item.Optional != nil {
			toAdd.Optional = &data.OptionalField{Default: item.Optional.Default}
		}

		if item.Validate != nil {
			toAdd.Validator = v.generateFieldValidatorView(fieldType, item.Validate)
		}

		v.view.Schema.Fields = append(v.view.Schema.Fields, toAdd)
	}

	// Add adapters
	v.view.Adapters.Input = v.generateInputsView()
	v.view.Adapters.Output = v.generateOutputsView()

	return v.view, nil
}

func (v *EntityViewGeneratorV1) generateOutputsView() data.OutputAdapters {
	outputAdapters := data.OutputAdapters{}

	for _, item := range v.entity.Adapters.Output {
		if item.Type == v1.Memory {
			outputAdapters.Memory = &data.MemoryAdapter{}
		}

		if item.Type == v1.Redis {
			outputAdapters.Redis = &data.RedisAdapter{}
			if iTtl, ok := item.Metadata[v1.TTL]; ok {
				if ttl, ok := iTtl.(int); ok {
					outputAdapters.Redis.TTL = ttl
				}
			}
		}

		if item.Type == v1.MongoDB {
			outputAdapters.MongoDB = &data.MongoDBAdapter{CollectionName: v.entity.Plural}
			if iCollection, ok := item.Metadata[v1.Collection]; ok {
				if collection, ok := iCollection.(string); ok {
					outputAdapters.MongoDB.CollectionName = collection
				}
			}
		}

		if item.Type == v1.SQL {
			outputAdapters.Sql = &data.SqlAdapter{TableName: v.entity.Plural}
			if iTable, ok := item.Metadata[v1.Table]; ok {
				if table, ok := iTable.(string); ok {
					outputAdapters.Sql.TableName = table
				}
			}
		}

		if item.Type == v1.RedisMongoDB {
			outputAdapters.RedisMongoDB = true
		}

		if item.Type == v1.RedisSQL {
			outputAdapters.RedisSql = true
		}
	}

	return outputAdapters
}

func (v *EntityViewGeneratorV1) generateInputsView() data.InputAdapters {
	inputAdapters := data.InputAdapters{}

	for _, item := range v.entity.Adapters.Input {
		if item.Type == v1.HTTP {

			// get resource name
			resourceName := v.entity.Plural
			if rname, ok := item.Metadata[v1.ResourceName]; ok {
				if name, ok := rname.(string); ok && name != "" {
					resourceName = name
				}
			}

			// get fields map
			fieldMap := map[string]string{}
			if rmap, ok := item.Metadata[v1.Map]; ok {
				if mmap, ok := rmap.(map[string]interface{}); ok {
					for fieldName, fieldRenamed := range mmap {
						if val, ok := fieldRenamed.(string); ok {
							fieldMap[fieldName] = val
						}
					}
				}
			}

			inputAdapters.Http = &data.HttpAdapter{
				ResourceName: resourceName,
				FieldMap:     fieldMap,
			}
		}
	}

	return inputAdapters
}

func (v *EntityViewGeneratorV1) generateFieldValidatorView(fieldType data.DataTypeEnum, validate *v1.ValidateItem) data.IItemValidator {
	if validate == nil {
		return nil
	}

	if validate.Value != nil {
		return &data.ItemValidatorValue{Type: fieldType, Value: validate.Value}
	}

	if validate.List != nil {
		return &data.ItemValidatorList{Type: fieldType, List: validate.List.Values}
	}

	if validate.Range != nil {
		return &data.ItemValidatorRange{Type: fieldType, Min: validate.Range.Min, Max: validate.Range.Max}
	}

	if validate.Regex != nil {
		return &data.ItemValidatorRegex{Pattern: validate.Regex.Default, GoPattern: validate.Regex.GoPattern}
	}

	return nil
}
