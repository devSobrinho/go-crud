package repository

import (
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	model "github.com/devSobrinho/go-crud/src/model/user"
	"go.mongodb.org/mongo-driver/mongo"
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
	FindUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserById(
		id string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_err.RestErr)
	FindUserByEmailAndPassword(email, password string) (model.UserDomainInterface, *rest_err.RestErr)
	UpdateUser(id string, userDomain model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(id string) (string, *rest_err.RestErr)
}

func getCollection(ur *userRepository) *mongo.Collection {
	collectionName := os.Getenv(constants.ENV_MONGO_COLLECTION_USER)
	collection := ur.databaseConnection.Collection(collectionName)
	return collection
}

func errorTreatmentNoDocuments(err error, errMessageNoDocuments string, errMessageDefault string) *rest_err.RestErr {
	if err == mongo.ErrNoDocuments {
		return rest_err.NewNotFoundError(errMessageNoDocuments)
	}

	return rest_err.NewInternalServerError(errMessageDefault)
}
