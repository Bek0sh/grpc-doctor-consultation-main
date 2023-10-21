package main

import (
	"log"

	"github.com/Bek0sh/online-market/main-page/internal/handler"
	"github.com/Bek0sh/online-market/main-page/internal/repository"
	"github.com/Bek0sh/online-market/main-page/internal/routes"
	"github.com/Bek0sh/online-market/main-page/internal/service"
	"github.com/Bek0sh/online-market/main-page/pkg/db"
	"github.com/gin-gonic/gin"
)

var hand *handler.Handler

func init() {
	db.Connection()
	database := db.GetDb()
	repo := repository.NewRepository(database)
	serv := service.NewService(repo)
	hand = handler.NewHandler(serv)
}

func main() {
	r := gin.Default()

	routes.Consultation(r, *hand)

	log.Fatal(r.Run(":8000"))
}
