package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/leveebreaks/hasher"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type mongoDbAuthRepo struct {
	uri string
}

// NewMongoDbAuthRepo mongodb://localhost:27017
func NewMongoDbAuthRepo(uri string) *mongoDbAuthRepo {
	return &mongoDbAuthRepo{uri: uri}
}

func (repo *mongoDbAuthRepo) CreateUser(userName, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(repo.uri))
	if err != nil {
		return "", err
	}

	coll := client.Database("auth").Collection("users")
	res := coll.FindOne(ctx, bson.D{{"userName", userName}})
	if res.Err() == nil {
		return "", errors.New("user with such name already exists")
	}

	hashedPass, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}

	uid := uuid.NewString()
	_, err = coll.InsertOne(ctx, bson.D{{"userName", userName}, {"password", hashedPass}, {"uid", uid}})
	if err == nil {
		return "", err
	}

	return uid, nil
}

func (repo *mongoDbAuthRepo) CheckUser(userName, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(repo.uri))
	if err != nil {
		return false
	}
	coll := client.Database("auth").Collection("users")
	hashedPass, err := hasher.HashPassword(password)
	if err != nil {
		return false
	}
	res := coll.FindOne(ctx, bson.D{{"userName", userName}, {"password", hashedPass}})
	if res.Err() == mongo.ErrNoDocuments {
		return false
	}

	return true
}
