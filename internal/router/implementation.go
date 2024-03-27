package router

import (
	"github.com/devomccar/HackathonChatbot/internal/handlers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type route struct {
	router *httprouter.Router
}

// New creates a new logger which uses zerolog in the backend
func New() Router {
	return route{
		router: httprouter.New(),
	}
}

func (r route) HandlerFunc(method, path string, handler handlers.Handler) {
	r.router.HandlerFunc(method, path, handler.Handle)
}

func (r route) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
