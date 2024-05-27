package middleware

import (
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	userRequest "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	"github.com/devSobrinho/go-crud/src/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func Logging(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	secretKey := []byte(os.Getenv(constants.ENV_JWT_SECRET))
	token, err := jwt.Parse(utils.RemoveBearerPrefix(authHeader), func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		logger.Error("Erro ao verificar token", err, zap.String("journey", "veryfyToken"))
		errRest := rest_err.NewUnauthorizedError("Token inválido")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("Token inválido", err, zap.String("journey", "veryfyToken"))
		errRest := rest_err.NewUnauthorizedError("Token inválido")
		c.JSON(errRest.Code, errRest)
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
