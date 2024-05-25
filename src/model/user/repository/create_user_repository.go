package repository

import (
	"context"
	"os"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
)

var (
	ENV_MONGO_COLLECTION_USER = "MONGO_COLLECTION_USER"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collectionName := os.Getenv(ENV_MONGO_COLLECTION_USER)

	collection := ur.databaseConnection.Collection(collectionName)
	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}
