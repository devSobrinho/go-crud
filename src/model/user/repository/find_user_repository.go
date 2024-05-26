package repository

import (
	"context"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUserByEmail repository", zap.String("journey", "findUserByEmail"))
	collection := getCollection(ur)
	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Usuário com email %s não encontrado"
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Erro ao tentar buscar usuário por email"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmail"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	logger.Info("CreateUser repository executado com sucesso", zap.String("journey", "createUser"))
	response := converter.ConvertEntityToDomain(*userEntity)
	return response, nil
}

func (ur *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))
	collection := getCollection(ur)
	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}}
	userEntity := &entity.UserEntity{}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := "Dados inválidos"
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmailAndPassword"))

			return nil, rest_err.NewNotFoundError(errorMessage)
		}
		errorMessage := "Dados inválidos"
		logger.Error(errorMessage,
			err,
			zap.String("journey", "findUserByEmailAndPassword"))

		return nil, rest_err.NewInternalServerError(errorMessage)
	}

	response := converter.ConvertEntityToDomain(*userEntity)

	return response, nil
}
