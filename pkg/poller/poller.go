package poller

import (
	"sync"

	"github.com/apex/log"
	"github.com/go-resty/resty/v2"

	"github.com/logikone/vdpcr/pkg/metrics"
)

type Poller struct {
	Client *resty.Client
	Log    log.Interface
	URLs   []string
}

func (p Poller) Poll() {
	wg := sync.WaitGroup{}

	p.Log.Debug("polling ...")

	for _, url := range p.URLs {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			res, err := p.Client.R().Get(u)

			if err != nil {
				p.Log.WithError(err).WithField("url", u).Error("error making request")
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
}
