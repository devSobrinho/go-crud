package repository

import (
	"context"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) UpdateUser(
	id string,
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Inicia updateUser repository", zap.String("journey", "updateUser"))

	collection := getCollection(u)
	defer disconnect(u)

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}
	age := userDomain.GetAge()
	name := userDomain.GetName()
	update := bson.D{{"$set", bson.D{{"name", name}, {"age", age}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		logger.Error("Erro ao tentar atualizar dados a colletion de user", err, zap.String("journey", "updateUser"))
		return rest_err.NewInternalServerError(err.Error())
	}

	logger.Info("UpdateUser repository executado com sucesso", zap.String("journey", "updateUser"))

	return nil
}
