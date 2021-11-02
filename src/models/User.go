package models

import (
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func Prepare(user *User) error {
	err := validateFields(user)
	if err != nil {
		return err
	}
	formatFields(user)

	return nil
}

func formatFields(user *User) {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
}

func validateFields(user *User) error {
	if user.Name == "" {
		return errors.New("name could not be empty")
	}

	if user.Email == "" {
		return errors.New("email could not be empty")
	}

	if user.Password == "" {
		return errors.New("password could not be empty")
	}

	return nil
}
