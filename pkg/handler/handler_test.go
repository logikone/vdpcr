package handler

import (
	"net/http"
	"testing"

	"github.com/logikone/vdpcr/pkg/poller"
)

func TestMetricsHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		Poller      poller.Poller
		PromHandler http.Handler
	}
	type args struct {
		writer  http.ResponseWriter
		request *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetricsHandler{
				Poller:      tt.fields.Poller,
				PromHandler: tt.fields.PromHandler,
			}

			m.ServeHTTP(tt.args.writer, tt.args.request)
		})
	}
}
