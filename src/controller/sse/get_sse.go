package sseController

import (
	"io"

	"github.com/devSobrinho/go-crud/src/configuration/sse"
	"github.com/gin-gonic/gin"
)

func (s *sseControllerStruct) SSE(c *gin.Context) {
	v, ok := c.Get("clientChan")
	if !ok {
		return
	}
	clientChan, ok := v.(sse.ClientChan)
	if !ok {
		return
	}
	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-clientChan; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
