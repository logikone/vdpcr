package main

import (
	"context"
	"flag"
	"net/http"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/iand/logfmtr"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/logikone/vdpcr/pkg/poller"
)

var (
	addr         = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
	pollInterval = flag.Duration("poll-interval", time.Second*30, "poll interval")
)

func main() {
	flag.Parse()

	log := logfmtr.New().WithName("vDPCR")

	http.Handle("/metrics", promhttp.Handler())

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go (&poller.Poller{
		Client:   resty.New(),
		Interval: pollInterval,
		Log:      log.WithName("Poller"),
	}).Start(ctx)

	log.Info("listening", "address", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Error(err, "error starting server")
	}
}
