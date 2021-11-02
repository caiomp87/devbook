package repositories

import (
	"api/src/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (db userDatabaseHelper) List(ctx context.Context) ([]*models.User, error) {
	cursor, err := db.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	users := make([]*models.User, 0)

	for cursor.Next(ctx) {
		var user *models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (db userDatabaseHelper) GetByID(ctx context.Context, ID string) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return nil, err
	}

	var user *models.User
	result := db.collection.FindOne(ctx, bson.M{"_id": objectID})
	if err = result.Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (db userDatabaseHelper) UpdateByID(ctx context.Context, ID string, user *models.User) error {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{
		"name":      user.Name,
		"email":     user.Email,
		"password":  user.Password,
		"updatedAt": time.Now(),
	}}

	result, err := db.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("cannot update user")
	}

	return nil
}

func (db userDatabaseHelper) DeleteByID(ctx context.Context, ID string) error {
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	result, err := db.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("cannot delete user")
	}

	return nil
}
