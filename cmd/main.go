package main

import (
	"librarys/pkc/handler"
	"log"
)

func main () {
	handlers := new(handler.Handler)

	srv := new(cards.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf(err)
	}
}