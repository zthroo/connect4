package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type gamesList struct {
	Games []string `json:"games"`
}

//each spot on the board will be either empty, or filled with a player string to indicate they've placed a token there.
//The format will be columns, then rows. For instance index 0,0 would be the bottom leftmost spot
type game struct {
	GameID     string     `json:"gameID"`
	Players    []string   `json:"players"`
	Columns    int        `json:"columns"`
	Rows       int        `json:"rows"`
	BoardState [][]string `json:"boardState"`
	GameState  string     `json:"gameState"`
}

//have a finished game struct to facilitate the requirement of no winner key if game in progress.
type finishedGame struct {
	GameID     string     `json:"gameID"`
	Players    []string   `json:"players"`
	Columns    int        `json:"columns"`
	Rows       int        `json:"rows"`
	BoardState [][]string `json:"boardState"`
	GameState  string     `json:"gameState"`
	Winner     string     `json:"winner"`
}

type gameID struct {
	GameID int64 `json:"gameId"`
}

func openGamesDB() (*sql.DB, error) {
	gamesDB, err := sql.Open("sqlite3", "./local.db")
	if err != nil {
		return nil, err
	}

	return gamesDB, err
}

func getGames(gamesDB *sql.DB) (gamesList, error) {
	var games gamesList

	const query = `SELECT game_id FROM game_table WHERE game_state = 'IN_PROGRESS'`
	rows, err := gamesDB.Query(query)
	if err != nil {
		return games, err
	}
	defer rows.Close()
	var game string
	for rows.Next() {
		err := rows.Scan(&game)
		if err != nil {
			return games, err
		}
		log.Println(game)
		games.Games = append(games.Games, game)
	}
	err = rows.Err()
	if err != nil {
		return games, err
	}
	return games, err
}

func createGame(players []string, columns, rows int, gamesDB *sql.DB) (gameID, error) {
	var gameID gameID
	var boardState [4][4]string
	playersString := strings.Join(players, ",")
	boardJSON, err := json.Marshal(boardState)
	const insertStmt = `INSERT INTO game_table (
		players,
		columns,
		rows,
		board_state,
		game_state
		) VALUES (?,?,?,?,'IN_PROGRESS')`
	_, err = gamesDB.Exec(insertStmt, playersString, columns, rows, boardJSON)
	if err != nil {
		return gameID, err
	}

	const getIDStmt = `select MAX(game_id) FROM game_table`
	var id int64
	err = gamesDB.QueryRow(getIDStmt).Scan(&id)
	if err != nil {
		return gameID, err
	}
	gameID.GameID = id
	return gameID, err
}
