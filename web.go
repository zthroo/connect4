package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getGamesWeb(c *gin.Context) {
	gamesDB, err := openGamesDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	games, err := getGames(gamesDB)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, games)
}

type CreateGameInput struct {
	Players []string `json:"players" binding:"required"`
	Columns int      `json:"columns" binding:"required"`
	Rows    int      `json:"rows" binding:"required"`
}

func createGameWeb(c *gin.Context) {
	var input CreateGameInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	gamesDb, err := openGamesDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gameID, err := createGame(input.Players, input.Columns, input.Rows, gamesDb)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gameID)
}

func getGameStateWeb(c *gin.Context) {
	id := c.Param("gameId")
	intId, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	gamesDb, err := openGamesDB()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	log.Println("id:", intId)
	gameState, err := getGamesState(intId, gamesDb)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, gameState)
}
