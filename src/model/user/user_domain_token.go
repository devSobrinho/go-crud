package model

import (
	"os"
	"time"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	rest_err "github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

var (
	ENV_JWT_SECRET = "JWT_SECRET"
)

func (u *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv(ENV_JWT_SECRET)

	claims := jwt.MapClaims{
		"id":    u.GetID(),
		"email": u.GetEmail(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Error("Erro ao gerar token", err, zap.String("journey", "generateToken"))
		return "", rest_err.NewInternalServerError("Erro ao gerar token")
	}

	return tokenString, nil
}
