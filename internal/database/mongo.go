package database

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongostore struct {
	db *mongo.Database
}

// If you want, you can pass these global variables as arguments to the NewMongo function.
// For simplicity, I have chosen to use global variables in this example, but it's easy to change.
var (
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	host     = os.Getenv("DB_HOST")
	name     = os.Getenv("DB_NAME")
)

func NewMongo() (Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(fmt.Sprintf(host, user, password)))
	if err != nil {
		return nil, err
	}

	return &mongostore{
		db: client.Database(name),
	}, nil
}

func (s *mongostore) Health(context context.Context) error {
	return s.db.Client().Ping(context, nil)
}

// func (s *mongostore) GetUser(userid string) (models.User, error) {
// 	var user models.User

// 	err := s.db.Collection("users").FindOne(context.Background(), bson.M{"userid": userid}).Decode(&user)
// 	if err != nil {
// 		return models.User{}, err
// 	}

// 	if user.UserID == "" {
// 		return models.User{}, fmt.Errorf("Brugeren blev ikke fundet.")
// 	}

// 	return user, nil
// }
