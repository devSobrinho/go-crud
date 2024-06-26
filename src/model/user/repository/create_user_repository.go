package repository

import (
	"context"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia createUser repository", zap.String("journey", "createUser"))

	collection := getCollection(u)
	defer disconnect(u)

	value := converter.ConvertDomainToEntity(userDomain)
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		logger.Error("Erro ao tentar inserir dados a colletion de user", err, zap.String("journey", "createUser"))
		return nil, rest_err.NewInternalServerError(err.Error())
	}
	value.ID = result.InsertedID.(primitive.ObjectID)
	response := converter.ConvertEntityToDomain(*value)

	logger.Info("CreateUser repository executado com sucesso", zap.String("journey", "createUser"))

	return response, nil
}
