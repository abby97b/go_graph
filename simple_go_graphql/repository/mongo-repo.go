package repository

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abbyb97/simple_go_graphql/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	Save(user *model.User)
	FindAll() []*model.User
	UpdateUser(id string, updatedUser *model.User)
}

type database struct {
	client *mongo.Client
}

const (
	DATABASE   = "simple_go_graphql"
	COLLECTION = "users"
)

func New() UserRepository {

	MONGODB := "mongodb+srv://abby:abby%40123@cluster0.chmma.mongodb.net/simple_go_graphql?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(MONGODB)
	clientOptions = clientOptions.SetMaxPoolSize(50)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	dbClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected to mongodb")

	return &database{
		client: dbClient,
	}
}

func (db *database) Save(user *model.User) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	_, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
}

func (db *database) UpdateUser(id string, updatedUser *model.User) {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)

	exp := 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), exp)
	defer cancel()

	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": bson.M{"FirstName": updatedUser.FirstName},
	}

	upsert := true
	after := options.After
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	// 8) Find one result and update it
	result := collection.FindOneAndUpdate(ctx, filter, update, &opt)
	if result.Err() != nil {
		log.Fatal(result)
	}

	// user := collection.FindOneAndUpdate(ctx, filter, update)
	// _, err := collection.ReplaceOne(ctx, user, updatedUser)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func (db *database) FindAll() []*model.User {
	collection := db.client.Database(DATABASE).Collection(COLLECTION)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())
	var result []*model.User
	for cursor.Next(context.TODO()) {
		var v *model.User
		err := cursor.Decode(&v)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, v)
	}
	return result
}
