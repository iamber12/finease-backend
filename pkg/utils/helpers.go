package utils

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResponseRenderer(message string, payload ...gin.H) gin.H {
	response := gin.H{
		"meta": gin.H{
			"message": message,
		},
	}
	if len(payload) > 0 {
		response["payload"] = payload[0]
	}
	return response
}

func Hash(payload string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload), 10)
	return string(bytes), err
}
