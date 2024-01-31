package db

import (
	"context"
	"fmt"

	"github.com/Meenachinmay/hotel-reservation-golang/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, bson.M, bson.M) error
}

type MongoUserStore struct {
	client     *mongo.Client
	dbName     string
	collection *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(userColl),
		dbName:     DBNAME,
	}
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err := s.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter, values bson.M) error {

	update := bson.D{
		{
			"$set", values,
		},
	}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*types.User
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}

	return users, nil
}

func (s *MongoUserStore) CreateUser(ctx context.Context, user *types.User) (*types.User, error) {
	result, err := s.collection.InsertOne(ctx, user)

	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	res, err := s.collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
