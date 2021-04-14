package handler

import (
	"net/http"

	"github.com/logikone/vdpcr/pkg/poller"
)

type MetricsHandler struct {
	Poller      poller.Poller
	PromHandler http.Handler
}

func (m MetricsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m.Poller.Poll()
	m.PromHandler.ServeHTTP(writer, request)
}
