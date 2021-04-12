package poller

import (
	"context"
	"sync"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-resty/resty/v2"

	"github.com/logikone/vdpcr/pkg/metrics"
)

var urls = []string{
	"https://httpstat.us/200",
	"https://httpstat.us/503",
}

type Poller struct {
	Client   *resty.Client
	Interval *time.Duration
	Log      logr.Logger
}

func (p Poller) Start(ctx context.Context) {
	t := time.Tick(*p.Interval)

	for {
		select {
		case <-t:
			if err := p.poll(); err != nil {
				p.Log.Error(err, "error polling")
			}
		case <-ctx.Done():
			return
		}
	}
}

func (p Poller) poll() error {
	wg := sync.WaitGroup{}

	p.Log.Info("polling ...")

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			res, err := p.Client.R().Get(u)

			if err != nil {
				p.Log.Error(err, "error making request", "url", u)
			}

			responseMetrics := metrics.ResponseTime.With(map[string]string{
				"url": u,
			})
			statusMetric := metrics.Status.With(map[string]string{
				"url": u,
			})

			if res.StatusCode() >= 200 && res.StatusCode() <= 399 {
				statusMetric.Set(0)
			} else {
				statusMetric.Set(1)
			}

			responseMetrics.Observe(float64(res.Time().Milliseconds()))

			res.Time()
		}(url)

	}

	wg.Wait()

	return nil
}
