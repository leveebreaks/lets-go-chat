package repository

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/leveebreaks/hasher"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type mongoDbAuthRepo struct {
	db *mongo.Database
}

// NewMongoDbAuthRepo ...
func NewMongoDbAuthRepo(db *mongo.Database) Auth {
	return &mongoDbAuthRepo{db}
}

func (repo *mongoDbAuthRepo) CreateUser(userName, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	coll := repo.db.Collection("users")
	res := coll.FindOne(ctx, bson.D{{"userName", userName}})
	if ctx.Err() != nil {
		panic(ctx.Err())
	}
	if res.Err() == mongo.ErrNoDocuments {
		return "", errors.New("user with such name already exists")
	}

	hashedPass, err := hasher.HashPassword(password)
	if err != nil {
		return "", err
	}

	uid := uuid.NewString()
	_, err = coll.InsertOne(ctx, bson.D{{"userName", userName}, {"password", hashedPass}, {"uid", uid}})
	if err != nil {
		return "", err
	}

	return uid, nil
}

func (repo *mongoDbAuthRepo) CheckUser(userName, password string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	coll := repo.db.Collection("users")
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
