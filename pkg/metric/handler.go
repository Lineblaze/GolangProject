package metric

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	log.Println("heartbeat")
	w.WriteHeader(204)
}
