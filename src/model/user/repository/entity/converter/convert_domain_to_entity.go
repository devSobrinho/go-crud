package converter

import (
	model "github.com/devSobrinho/go-crud/src/model/user"
	"github.com/devSobrinho/go-crud/src/model/user/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
