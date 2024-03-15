package database

import (
	"context"

	"golang-mongodb-logger/configs"
	"golang-mongodb-logger/pkg/core"

	"github.com/go-logr/zerologr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection(ctx context.Context) (*mongo.Database, error) {
	logger := *core.Logger()

	sink := zerologr.
		New(&logger).
		GetSink()

	loggerOptions := options.
		Logger().
		SetSink(sink).
		SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)

	clientOptions := options.
		Client().
		ApplyURI(configs.Env("MONGOURI")).
		SetLoggerOptions(loggerOptions)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	db := client.Database(configs.Env("MONGOURI_DATABASE"))

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return db, nil
}
