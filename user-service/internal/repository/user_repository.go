package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user-service/internal/models"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	// Check if username already exists
	filter := bson.M{"username": user.Username}
	if err := r.collection.FindOne(ctx, filter).Err(); err == nil {
		return errors.New("username already exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	user.GenerateID()
	_, err := r.collection.InsertOne(ctx, user)
	return err
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}