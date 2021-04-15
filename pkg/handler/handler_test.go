package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/apex/log"
	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/metrics", strings.NewReader(""))

	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "OK",
			fields: fields{
				Poller: poller.Poller{
					Client: resty.New(),
					Log:    log.Log,
					URLs:   []string{"https://httpstat.us/200", "https://httpstat.us/503"},
				},
				PromHandler: promhttp.Handler(),
			},
			args: args{
				writer:  recorder,
				request: request,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MetricsHandler{
				Poller:      tt.fields.Poller,
				PromHandler: tt.fields.PromHandler,
			}

			m.ServeHTTP(tt.args.writer, tt.args.request)

			if recorder.Result().StatusCode != http.StatusOK {
				t.Errorf("unexpected http status code %d != %d",
					recorder.Result().StatusCode, http.StatusOK)
			}

			bodyBytes, err := ioutil.ReadAll(recorder.Body)
			if err != nil {
				t.Fatalf("error reading repsonse body: %s", err)
			}

			for _, url := range tt.fields.Poller.URLs {
				upRe, _ := regexp.Compile(fmt.Sprintf(`external_url_up{url="%s"}`, url))
				resSumRe, _ := regexp.Compile(fmt.Sprintf(`external_url_response_ms_sum{url="%s"}`, url))
				resCountRe, _ := regexp.Compile(fmt.Sprintf(`external_url_response_ms_count{url="%s"}`, url))

				foundUp := upRe.Find(bodyBytes)

				if len(foundUp) == 0 {
					t.Errorf("unable to find up metric for url=%s", url)
				}

				foundResSum := resSumRe.Find(bodyBytes)

				if len(foundResSum) == 0 {
					t.Errorf("unable to find response time sum metric for url=%s", url)
				}

				foundResCount := resCountRe.Find(bodyBytes)

				if len(foundResCount) == 0 {
					t.Errorf("unable to find response time count metric for url=%s", url)
				}
			}
		})
	}
}
