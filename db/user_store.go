package db

import (
	"context"

	"github.com/Meenachinmay/hotel-reservation-golang/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
}

type MongoUserStore struct{
	client *mongo.Client
	dbName string
	collection *mongo.Collection
}

func NewMongoUserStore (client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		collection: client.Database(DBNAME).Collection(userColl),
		dbName: DBNAME,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	if err := s.collection.FindOne(ctx, bson.M{"_id": ToObjectID(id)}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}