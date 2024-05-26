package repository

import (
	"context"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userRepository) DeleteUser(id string) (string, *rest_err.RestErr) {
	logger.Info("Inicia deleteUser repository", zap.String("journey", "deleteUser"))

	collection := getCollection(u)
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		errorMessage := "Usuário não encontrado"
		logger.Error(errorMessage, err, zap.String("journey", "deleteUser"))
		return "", errorTreatmentNoDocuments(err, errorMessage, errorMessage)
	}

	logger.Info("DeleteUser repository executado com sucesso", zap.String("journey", "deleteUser"))
	return "Usuário deletado com sucesso", nil
}
