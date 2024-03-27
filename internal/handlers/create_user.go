package handlers

import (
	"net/http"
)

type createUserHandler struct{}

func NewCreateUserHandler() Handler {
	return createUserHandler{}
}

func (h createUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	// var u models.User

	// logic of handler
	// r.Body.
	w.Write([]byte("Hello"))

}
