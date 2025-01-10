package db

import (
	"context"
	"log"
	"os"
	"tripatra-test-go/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var userCollection *mongo.Collection

func Connect() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("mongo_uri")))
	if err != nil {
		log.Fatal(err)
	}
	userCollection = client.Database(os.Getenv("db_name")).Collection("users")
}

func CreateUser(user *models.User) (*models.User, error) {
	result, err := userCollection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func UpdateUser(id string, name *string, email *string) (*models.User, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{}
	if name != nil {
		update["name"] = *name
	}
	if email != nil {
		update["email"] = *email
	}
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return GetUser(id)
}

func DeleteUser(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := userCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}

func GetUser(id string) (*models.User, error) {
	var user models.User
	objID, _ := primitive.ObjectIDFromHex(id)
	err := userCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
