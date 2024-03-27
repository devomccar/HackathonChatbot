package handlers

import "net/http"

type customHandler struct{}

func NewCustomHandler() Handler {
	return customHandler{}
}

func (h customHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
