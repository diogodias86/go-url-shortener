package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func defaultContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func getClient() *mongo.Client {
	ctx, cancel := defaultContext()
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@localhost:27017"))
	if err != nil {
		panic(err)
	}

	return client
}

func getDatabase() *mongo.Database {
	return getClient().Database("go_url_shortener")
}

func TestConnection() {
	client := getClient()

	ctx, cancel := defaultContext()
	defer cancel()

	fmt.Println("Realizando conexão com o MongoDB....")

	client.Ping(ctx, readpref.Primary())

	fmt.Println("Conexão realizada com sucesso.")
}

func Insert(newURL string, shortURL string) {
	ctx, cancel := defaultContext()
	defer cancel()

	_, err := getDatabase().Collection("urls").InsertOne(ctx,
		bson.M{"originalurl": newURL, "shorturl": shortURL},
	)

	if err != nil {
		panic(err)
	}
}

func GetURL(shortURL string) string {
	ctx, cancel := defaultContext()
	defer cancel()

	var result struct {
		OriginalURL string
	}

	err := getDatabase().Collection("urls").FindOne(ctx,
		bson.M{"shorturl": shortURL},
	).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return ""
		}

		panic(err)
	}

	return result.OriginalURL
}
