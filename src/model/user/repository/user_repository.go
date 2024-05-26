package repository

import (
	"os"

	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ENV_MONGO_COLLECTION_USER = "MONGO_COLLECTION_USER"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepositoryInterface {
	return &userRepository{
		databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepositoryInterface interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
}

func getCollection(ur *userRepository) *mongo.Collection {
	collectionName := os.Getenv(ENV_MONGO_COLLECTION_USER)
	collection := ur.databaseConnection.Collection(collectionName)
	return collection
}
