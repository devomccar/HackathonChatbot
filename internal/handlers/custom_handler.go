package handlers

import (
	"github.com/go-redis/redis"
	"net/http"
	"time"
)

type customHandler struct {
	db *redis.Client
}

func NewCustomHandler(db *redis.Client) Handler {
	return customHandler{db: db}
}

func (h customHandler) Handle(w http.ResponseWriter, r *http.Request) {

	// product question: how long do we want to persist the data for?
	err := h.db.Set("firstName", "devon", time.Hour).Err()
	// logic of handler
	w.Write([]byte("Hello"))

}
