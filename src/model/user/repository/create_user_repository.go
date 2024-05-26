package repository

import (
	"context"
	"os"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.uber.org/zap"
)

var (
	ENV_MONGO_COLLECTION_USER = "MONGO_COLLECTION_USER"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia createUser repository", zap.String("journey", "createUser"))

	collectionName := os.Getenv(ENV_MONGO_COLLECTION_USER)

	collection := ur.databaseConnection.Collection(collectionName)
	value, err := userDomain.GetJSONValue()
	if err != nil {
		logger.Error("Erro ao tentar chamar GetJSONValue", err, zap.String("journey", "createUser"))

		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error("Erro ao tentar inserir dados a colletion de user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))

	logger.Info("CreateUser repository executado com sucesso", zap.String("journey", "createUser"))

	return userDomain, nil
}
