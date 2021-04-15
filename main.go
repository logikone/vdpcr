package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/logfmt"
	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	flag "github.com/spf13/pflag"

	"github.com/logikone/vdpcr/pkg/handler"
	"github.com/logikone/vdpcr/pkg/poller"
)

var (
	addr           = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	defaultTargets = []string{"https://httpstat.us/200", "https://httpstat.us/503"}
	logLevel       = flag.String("log-level", "info", fmt.Sprintf("log level to set. (%s)", strings.Join(validLogLevels, ", ")))
	targets        = flag.StringArray("target", defaultTargets, "target(s) to scrape. can be specified multiple times")
	timeout        = flag.Duration("timeout", time.Second*10, "http client timeout")
	validLogLevels = []string{"debug", "error", "fatal", "info", "warn"}
)

func main() {
	flag.Parse()
	log.SetLevelFromString(*logLevel)

	httpClient := resty.New()
	httpClient.SetTimeout(*timeout)

	http.Handle("/metrics", &handler.MetricsHandler{
		Poller: poller.Poller{
			Client: httpClient,
			Log:    log.Log,
			URLs:   *targets,
		},
		PromHandler: promhttp.Handler(),
	})

	log.Infof("listening on %s", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.WithError(err).Error("error starting server")
	}
}

func init() {
	log.SetHandler(logfmt.New(os.Stdout))
}
