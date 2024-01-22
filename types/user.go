package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName      string             `bson:"firstName" json:"firstName"`
	LastName       string             `bson:"lastName" json:"lastName"`
	Email          string             `bson:"email" json:"email"`
	HashedPassword string             `bson:"hashedPassword" json: "-", `
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 12)

	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:      params.FirstName,
		LastName:       params.LastName,
		Email:          params.Email,
		HashedPassword: string(hashedPassword),
	}, nil
}
