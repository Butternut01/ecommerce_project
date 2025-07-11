package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}

func (u *User) GenerateID() {
	u.ID = primitive.NewObjectID()
}

func (u *User) GetID() string {
	return u.ID.Hex()
}