package types

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

const (
	minFirstNameLen = 2
	minLastNameLen  = 3
	maxFirstNameLen = 12
	maxLastNameLen  = 12
	minPasswordLen  = 7
	maxPasswordLen  = 18
)

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() error {

	if len(params.FirstName) < minFirstNameLen {
		return fmt.Errorf("firstName length should at least %d chars", minFirstNameLen)
	}

	if len(params.LastName) < minLastNameLen {
		return fmt.Errorf("lastName length should at least %d chars", minLastNameLen)
	}

	if len(params.Password) < minPasswordLen {
		return fmt.Errorf("password length should at least %d chars", minPasswordLen)
	}

	//if !isEmailValid(params.Email) {
	//	return fmt.Errorf("Email is invalid")
	//}

	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+|*]+@[a-z0-9. \-]+\-[a-z]{2,4}$*`)
	return emailRegex.MatchString(e)
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
