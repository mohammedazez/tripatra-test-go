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
var productCollection *mongo.Collection

func Connect() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("mongo_uri")))
	if err != nil {
		log.Fatal(err)
	}
	userCollection = client.Database(os.Getenv("db_name")).Collection("users")
	productCollection = client.Database(os.Getenv("db_name")).Collection("products")
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

func GetUsers() ([]*models.User, error) {
	cursor, err := userCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var users []*models.User
	for cursor.Next(context.Background()) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func CreateProduct(product *models.Product) (*models.Product, error) {
	result, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		return nil, err
	}
	product.ID = result.InsertedID.(primitive.ObjectID)
	return product, nil
}

func UpdateProduct(id string, name *string, price *float64, stock *int) (*models.Product, error) {
	objID, _ := primitive.ObjectIDFromHex(id)
	update := bson.M{}
	if name != nil {
		update["name"] = *name
	}
	if price != nil {
		update["price"] = *price
	}
	if stock != nil {
		update["stock"] = *stock
	}
	_, err := productCollection.UpdateOne(context.Background(), bson.M{"_id": objID}, bson.M{"$set": update})
	if err != nil {
		return nil, err
	}
	return GetProduct(id)
}

func DeleteProduct(id string) error {
	objID, _ := primitive.ObjectIDFromHex(id)
	_, err := productCollection.DeleteOne(context.Background(), bson.M{"_id": objID})
	return err
}

func GetProducts() ([]*models.Product, error) {
	cursor, err := productCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	var products []*models.Product
	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func GetProduct(id string) (*models.Product, error) {
	var product models.Product
	objID, _ := primitive.ObjectIDFromHex(id)
	err := productCollection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&product)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
