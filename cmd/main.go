package main

import (
	"fmt"
	internalDb "github.com/devomccar/HackathonChatbot/internal/db"
	internalHandler "github.com/devomccar/HackathonChatbot/internal/handlers"
	internalRouter "github.com/devomccar/HackathonChatbot/internal/router"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

func main() {
	// get db
	db := internalDb.NewRedis()

	// get handlers
	handlers := getHandlers(db)

	// get routes by passing in handler
	router := internalRouter.New()
	getRoutes(router, handlers)

	// start the server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func getHandlers(db *redis.Client) handlers {
	return handlers{
		customHandler: internalHandler.NewCustomHandler(db),
	}
}

func getRoutes(r internalRouter.Router, h handlers) {
	//linking the end points to the handler
	r.HandlerFunc(http.MethodGet, "/custom", h.customHandler)

}
