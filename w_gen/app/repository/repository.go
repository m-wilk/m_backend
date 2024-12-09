package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

const DB_TIMEOUT = time.Second * 3

var db *mongo.Client

type Repository struct {
	UserRepository UserRepository
}

func collection(databaseName string, collectionName string) *mongo.Collection {
	return db.Database(databaseName).Collection(collectionName)
}

func New(client *mongo.Client, databaseName string) *Repository {
	db = client

	return &Repository{
		UserRepository: NewUserRepository(collection(databaseName, "user")),
	}
}

func CloseDB() error {
	return db.Disconnect(context.TODO())
}
