package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("drop_token", getGamesWeb)
	router.POST("drop_token", createGameWeb)

	router.Run("localhost:8080")
}
