package utils

import (
	"strings"

	"github.com/devSobrinho/go-crud/src/configuration/logger"
	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

func removeBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

func ValidationTokenAndExtractClaims(secretKey string, tokenString string) (jwt.MapClaims, *rest_err.RestErr) {
	logger.Info("Inicia validação do token", zap.String("journey", "validationTokenAndExtractClaims"))

	secretKeyByte := []byte(secretKey)
	token, err := jwt.Parse(removeBearerPrefix(tokenString), func(token *jwt.Token) (interface{}, error) {
		return secretKeyByte, nil
	})

	if err != nil {
		logger.Error("Erro ao verificar token", err, zap.String("journey", "validationTokenAndExtractClaims"))
		errRest := rest_err.NewUnauthorizedError("Token inválido")
		return nil, errRest
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		logger.Error("Token inválido", err, zap.String("journey", "validationTokenAndExtractClaims"))
		errRest := rest_err.NewUnauthorizedError("Token inválido")
		return nil, errRest
	}

	logger.Info("Token validado com sucesso", zap.String("journey", "validationTokenAndExtractClaims"))
	return claims, nil
}

func TernaryCondition(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}
