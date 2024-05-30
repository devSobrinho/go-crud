package repository

import (
	"context"
	"fmt"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) FindUser(
	userDomain model.UserDomainInterfacePagination,
) ([]model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUser repository", zap.String("journey", "findUser"))

	collection := getCollection(u)
	userEntity := &entity.UserEntity{}
	getUser := userDomain.GetUser()

	filter := bson.D{}
	if getUser.GetID() != "" {
		objectId, _ := primitive.ObjectIDFromHex(getUser.GetID())
		filter = append(filter, bson.E{Key: "_id", Value: objectId})
	}

	if getUser.GetEmail() != "" {
		filter = append(filter, bson.E{Key: "email", Value: getUser.GetEmail()})
	}

	fmt.Println(">>>", userDomain.GetPagination().Order)

	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logger.Error("Erro findUser ao buscar usuario", err, zap.String("journey", "findUser"))

		errorMessage := "Dados inválidos"
		errRest := errorTreatmentNoDocuments(err, errorMessage, errorMessage)
		return nil, errRest
	}

	_ = cur.All(context.Background(), userEntity)

	logger.Info("FindUser repository executado com sucesso", zap.String("journey", "findUser"))
	var usersEntity []entity.UserEntity
	if err = cur.All(context.Background(), &usersEntity); err != nil {
		panic(err)
	}

	var results []model.UserDomainInterface
	for _, user := range usersEntity {
		response := converter.ConvertEntityToDomain(user)
		results = append(results, response)
	}

	return results, nil
}

func (u *userRepository) FindUserById(
	id string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUserById repository", zap.String("journey", "findUserById"))

	collection := getCollection(u)
	userEntity := &entity.UserEntity{}

	logger.Info("Inicia findUserById repository", zap.String("journey", "findUserById"))

	objectId, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: objectId}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)
	if err != nil {
		logger.Error("Erro findUserById ao buscar usuario", err, zap.String("journey", "findUserByEmailAndPassword"))

		errorMessage := "Dados inválidos"
		errRest := errorTreatmentNoDocuments(err, errorMessage, errorMessage)
		return nil, errRest
	}
	logger.Info(userEntity.Name, zap.String("journey", "findUserById"))
	response := converter.ConvertEntityToDomain(*userEntity)
	return response, nil
}

func (u *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUserByEmail repository", zap.String("journey", "findUserByEmail"))

	collection := getCollection(u)
	userEntity := &entity.UserEntity{}
	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		logger.Error("Erro findUserByEmail ao buscar usuario", err, zap.String("journey", "findUserByEmail"))

		errRest := errorTreatmentNoDocuments(err, "Usuário não encontrado", "Erro ao tentar buscar usuário")
		return nil, errRest
	}

	logger.Info("FindUserByEmail repository executado com sucesso", zap.String("journey", "findUserByEmail"))
	response := converter.ConvertEntityToDomain(*userEntity)
	return response, nil
}

func (u *userRepository) FindUserByEmailAndPassword(
	email string,
	password string,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Inicia findUserByEmailAndPassword repository", zap.String("journey", "findUserByEmailAndPassword"))

	collection := getCollection(u)
	defer disconnect(u)
	filter := bson.D{{Key: "email", Value: email}, {Key: "password", Value: password}}
	userEntity := &entity.UserEntity{}

	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		logger.Error("Erro findUserByEmailAndPassword ao buscar usuario", err, zap.String("journey", "findUserByEmailAndPassword"))

		errorMessage := "Dados inválidos"
		errRest := errorTreatmentNoDocuments(err, errorMessage, errorMessage)
		return nil, errRest
	}

	response := converter.ConvertEntityToDomain(*userEntity)

	return response, nil
}
