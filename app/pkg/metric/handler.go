package metric

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	URL = "/api/heartbeat"
)

type Handler struct {
}

//Register

func (h *Handler) Register(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, URL, h.Heartbeat)
}

// Heartbeat
// @Summary Heartbeat metric
// @Tags Metrics
// @Success 204
// @Router /api/heartbeat [get]
func (h *Handler) Heartbeat(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(204)
}
