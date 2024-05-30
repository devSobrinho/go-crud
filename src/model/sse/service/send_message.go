package serviceSSE

import (
	"fmt"

	"github.com/devSobrinho/go-crud/src/configuration/rest_err"
	"github.com/devSobrinho/go-crud/src/configuration/sse"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (*sseService) SendMessage(message string, id string) (string, *rest_err.RestErr) {

	if id != "" {
		if _, err := primitive.ObjectIDFromHex(id); err != nil {
			return "", rest_err.NewBadRequestError("ID inválido")
		}

		userChan := sse.Stream.TotalClients[id]
		if userChan == nil {
			return "", rest_err.NewNotFoundError("Usuário não encontrado")
		}
		userChan <- message
	} else {
		fmt.Println("TotalClients: ", sse.Stream.TotalClients)
		for _, clientChan := range sse.Stream.TotalClients {
			fmt.Println("ClientChan: ", clientChan)
			clientChan <- message
		}
	}

	return "Mensagem enviada com sucesso", nil
}
