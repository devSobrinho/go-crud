package controller

import (
	"net/http"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/devSobrinho/go-crud/src/configuration/validation"
	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	response "github.com/devSobrinho/go-crud/src/controller/model/response/user"
	model "github.com/devSobrinho/go-crud/src/model/user"
	view "github.com/devSobrinho/go-crud/src/view/user"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (u *userController) FindUser(c *gin.Context) {
	logger.Info("Inicia FindUser controller", zap.String("journey", "findUser"))
	var userRequest request.UserListRequest
	if err := c.ShouldBindQuery(&userRequest); err != nil {
		logger.Error("Erro ao tentar validar as informações da query", err, zap.String("journey", "findUser"))

		errRest := validation.ValidateError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	if userRequest.ID != "" {
		if _, err := primitive.ObjectIDFromHex(userRequest.ID); err != nil {
			logger.Error("Erro ao tentar validar o ID do usuário", err, zap.String("journey", "findUser"))
			errRest := rest_err.NewBadRequestError("ID do usuário inválido")
			c.JSON(errRest.Code, errRest)
			return
		}
	}

	domain := model.NewUserDomainPagination(
		userRequest.Email,
		userRequest.ID,
		userRequest.PaginationQuery.Page,
		userRequest.PaginationQuery.Size,
		userRequest.PaginationQuery.Order,
	)

	domainResult, err := u.service.FindUser(domain)
	if err != nil {
		logger.Error("Erro ao tentar chamar FindUser service", err, zap.String("journey", "findUser"))

		c.JSON(err.Code, err)
		return
	}

	var response []response.UserCreateResponse
	for _, userDomain := range domainResult {
		item := view.ConvertUserDomainToResponse(userDomain)
		response = append(response, item)
	}

	c.JSON(http.StatusOK, response)
}

func (u *userController) FindUserById(c *gin.Context) {
	logger.Info("Inicia FindUserById controller", zap.String("journey", "findUserById"))
	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Erro ao tentar validar o ID do usuário", err, zap.String("journey", "findUserById"))
		errRest := rest_err.NewBadRequestError("ID do usuário inválido")
		c.JSON(errRest.Code, errRest)
		return
	}

	domainResult, err := u.service.FindUserById(userId)
	if err != nil {
		logger.Error("Erro ao tentar chamar FindUserById service", err, zap.String("journey", "findUserById"))

		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertUserDomainToResponse(domainResult)
	c.JSON(http.StatusOK, response)
}

func (u *userController) FindUserByEmail(c *gin.Context) {
	logger.Info("Inicia FindUserByEmail controller", zap.String("journey", "findUserByEmail"))
	userEmail := c.Param("userEmail")

	domainResult, err := u.service.FindUserByEmail(userEmail)
	if err != nil {
		logger.Error("Erro ao tentar chamar FindUserById service", err, zap.String("journey", "findUserByEmail"))

		c.JSON(err.Code, err)
		return
	}

	response := view.ConvertUserDomainToResponse(domainResult)
	c.JSON(http.StatusOK, response)
}
