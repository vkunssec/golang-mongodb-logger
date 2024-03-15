package main

import (
	"context"

	"github.com/vkunssec/golang-mongodb-logger/pkg/core"
	"github.com/vkunssec/golang-mongodb-logger/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	ctx := context.Background()

	logger := core.Logger()
	db, _ := database.Connection(ctx)

	coll := db.Client().Database("test").Collection("test")

	res, err := coll.InsertOne(ctx, bson.M{"Alice": "123"})
	if err != nil {
		logger.Fatal().Msgf("InsertOne failed: %v", err)
	}

	logger.Info().
		Msgf("Value saved - ObjectId %s", res.InsertedID.(primitive.ObjectID).Hex())
}
