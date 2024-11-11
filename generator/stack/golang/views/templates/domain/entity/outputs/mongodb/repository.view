// Package mongodb is an output adapter to store entities in MongoDB
package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// mongoDBRepository logical repository to save entity in MongoDB
// Custom repository logic should be placed here.
type mongoDBRepository struct {
	*mongoDBRepositoryBase // DO NOT REMOVE IT
}

// NewMongoDBRepository repository constructor
func NewMongoDBRepository(collection *mongo.Collection, cryptoKey []byte) *mongoDBRepository {
	return &mongoDBRepository{mongoDBRepositoryBase: newMongoDBRepositoryBase(collection, cryptoKey)}
}

// NewMongoDBRepositoryWithHooks repository constructor
func NewMongoDBRepositoryWithHooks(collection *mongo.Collection, cryptoKey []byte, hooks *HookBuilder) *mongoDBRepository {
	return &mongoDBRepository{mongoDBRepositoryBase: newMongoDBRepositoryBaseWithHooks(collection, cryptoKey, hooks)}
}
