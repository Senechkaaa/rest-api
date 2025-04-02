package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponce(c *gin.Context, statusCode int, message string) {
	logrus.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{Message: message})
	//	 прерывает дальнейшую работу и отправляет клиенту статус код
}

type statusResponce struct {
	Status string `json:"status"`
}
