package middleware

import (
	"os"

	constants "github.com/devSobrinho/go-crud/src/configuration/contants"
	userRequest "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	"github.com/devSobrinho/go-crud/src/utils"
	"github.com/gin-gonic/gin"
)

func Logging(c *gin.Context) {
	token := c.GetHeader("Authorization")

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
