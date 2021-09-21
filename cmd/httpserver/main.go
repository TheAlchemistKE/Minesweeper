package main

import (
	gamesrv "github.com/TheAlchemistKE/Minesweeper/internal/core/services/gameserver"
	"github.com/TheAlchemistKE/Minesweeper/internal/handlers/gamehandler"
	"github.com/TheAlchemistKE/Minesweeper/internal/repositories/gamesrepo"
	"github.com/gin-gonic/gin"
)

func main() {
	gamesRepository := gamesrepo.NewMemKVS()
	gamesService := gamesrv.New(gamesRepository, uidgen.New())
	gamesHandler := gamehandler.NewHTTPHandler(gamesService)

	router := gin.New()
	router.GET("/games/:id", gamesHandler.Get)
	router.POST("/games", gamesHandler.Create)

	router.Run(":8080")
}
