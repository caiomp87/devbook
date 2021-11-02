package repositories

import (
	"api/src/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type IUser interface {
	Create(context.Context, *models.User) error
	List(context.Context) ([]*models.User, error)
	GetByID(context.Context, string) (*models.User, error)
	UpdateByID(context.Context, string, *models.User) error
	DeleteByID(context.Context, string) error
}

type userDatabaseHelper struct {
	collection *mongo.Collection
}

var UserCollection IUser

func NewUserCollection() IUser {
	return &userDatabaseHelper{
		collection: Database.Collection("users"),
	}
}

func (db userDatabaseHelper) Create(ctx context.Context, user *models.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := db.collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (db userDatabaseHelper) List(ctx context.Context) (users []*models.User, err error) {
	return nil, nil
}

func (db userDatabaseHelper) GetByID(ctx context.Context, ID string) (*models.User, error) {
	return nil, nil
}

func (db userDatabaseHelper) UpdateByID(ctx context.Context, ID string, user *models.User) error {
	return nil
}

func (db userDatabaseHelper) DeleteByID(ctx context.Context, ID string) error {
	return nil
}
