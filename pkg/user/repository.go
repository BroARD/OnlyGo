package user

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user NewUser) error
	GetUsers(ctx context.Context) (*mongo.Cursor, error)

}

type userRepository struct {
	MongoDB *mongo.Client
}

func NewRepository(mongoDB *mongo.Client) UserRepository {
	return &userRepository{MongoDB: mongoDB}
}

func (r *userRepository) CreateUser(ctx context.Context, user NewUser) error {
	collection := r.MongoDB.Database("mongo").Collection("test1234")
	_, err := collection.InsertOne(ctx, user)
	return err
}

func (r *userRepository) GetUsers(ctx context.Context) (*mongo.Cursor, error) {
	collection := r.MongoDB.Database("mongo").Collection("test1234")
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	return cur, nil
}
