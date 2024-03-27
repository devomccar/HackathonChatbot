package main

import (
	"fmt"
	internalHandler "github.com/devomccar/HackathonChatbot/internal/handlers"
	internalRouter "github.com/devomccar/HackathonChatbot/internal/router"
	"log"
	"net/http"
)

func main() {
	// get handlers
	handlers := getHandlers()

	// get routes by passing in handler
	router := internalRouter.New()
	getRoutes(router, handlers)

	// start the server
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

func getHandlers() handlers {
	return handlers{
		customHandler: internalHandler.NewCustomHandler(),
	}
}

func getRoutes(r internalRouter.Router, h handlers) {
	//linking the end points to the handler
	r.HandlerFunc(http.MethodGet, "/custom", h.customHandler)

}
