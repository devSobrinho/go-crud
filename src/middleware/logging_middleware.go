package middleware

import (
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	userRequest "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	"github.com/devSobrinho/go-crud/src/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Logging(c *gin.Context) {
	tokenHeader := c.GetHeader("Authorization")
	cookie, errCookie := c.Cookie("token")
	if errCookie != nil {
		logger.Info("Cookie não encontrado", zap.String("journey", "logging"))
	}

	if tokenHeader == "" && cookie == "" {
		logger.Info("Token não encontrado", zap.String("journey", "logging"))
		errRest := rest_err.NewUnauthorizedError("Token não encontrado")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	token := utils.TernaryCondition(tokenHeader == "", cookie, tokenHeader).(string)

	secretKey := os.Getenv(constants.ENV_JWT_SECRET)

	claims, err := utils.ValidationTokenAndExtractClaims(secretKey, token)
	if err != nil {
		c.JSON(err.Code, err)
		c.Abort()
		return
	}

	user := userRequest.UserRequest{
		ID:    claims["id"].(string),
		Email: claims["email"].(string),
	}

	c.Set("user", user)
	c.Next()
}
