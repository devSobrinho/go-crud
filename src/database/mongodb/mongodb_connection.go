package mongodb

import (
	"context"
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongoURI := os.Getenv(constants.ENV_MONGO_URL)
	mongoDB := os.Getenv(constants.ENV_MONGO_DB)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		logger.Error("Erro ao conectar ao MongoDB", err, zap.String("journey", "NewMongoDBConnection"))
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Erro ao fazer ping no banco de dados", err, zap.String("journey", "NewMongoDBConnection"))
		return nil, err
	}

	logger.Info("Conectado ao MongoDB", zap.String("journey", "NewMongoDBConnection"))

	return client.Database(mongoDB), nil
}
