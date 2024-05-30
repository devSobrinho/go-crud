package sse

import (
	"fmt"
	"log"

	request "github.com/devSobrinho/go-crud/src/controller/model/request/user"
	"github.com/gin-gonic/gin"
)

var (
	Stream *Event
)

type ClientChan chan string

type UserClient struct {
	ID   string
	Chan ClientChan
}

type Event struct {
	Message       chan string
	NewClients    chan UserClient
	ClosedClients chan UserClient
	TotalClients  map[string]ClientChan
}

func NewServer() (event *Event) {
	event = &Event{
		Message:       make(chan string),
		NewClients:    make(chan UserClient),
		ClosedClients: make(chan UserClient),
		TotalClients:  make(map[string]ClientChan),
	}

	go event.listen()

	return
}

func (stream *Event) listen() {
	for {
		select {
		case userClient := <-stream.NewClients:
			stream.TotalClients[userClient.ID] = userClient.Chan
			log.Printf("id:  %d", userClient.ID)
			log.Printf("Client added. %d registered clients", len(stream.TotalClients))

		case userClient := <-stream.ClosedClients:
			delete(stream.TotalClients, userClient.ID)
			close(userClient.Chan)
			log.Printf("Removed client. %d registered clients", len(stream.TotalClients))

		case eventMsg := <-stream.Message:
			fmt.Println("EventMsg: ", eventMsg)
			for _, clientChan := range stream.TotalClients {
				fmt.Println("ENVIA ClientChan: ", clientChan)
				clientChan <- eventMsg
			}
		}
	}
}

func (stream *Event) ServeHTTP() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Get("user")
		if !ok {
			c.AbortWithStatus(401)
			return
		}
		userID := user.(request.UserRequest).ID

		clientChan := make(ClientChan)

		stream.NewClients <- UserClient{ID: userID, Chan: clientChan}

		defer func() {
			fmt.Println("ClosedClients: ", userID)
			stream.ClosedClients <- UserClient{ID: userID, Chan: clientChan}
		}()

		c.Set("clientChan", clientChan)

		c.Next()
	}
}
