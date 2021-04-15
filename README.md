<div align="center">
    <h1>vdpcr</h1>
    <i>A sample Prometheus Exporter</i>
</div>

## Installation

Example kubernetes manifests are in the [deploy/manifests](deploy/manifests) directory and can be deployed as is. The
manifests will create a `vdpcr` namespace as well as a deployment and service. The deployment will add
the `prometheus.io/scrape=true` annotation which may work for you depending on your prometheus setup; if not an example
prometheus scrape configuration can be found [below](#prometheus).

## Usage

Multiple targets can be polled by specifying the `--target` flag multiple times in the deployment manifest. By default
if no target(s) are specified then the following two urls will be polled: https://httpstat.us/200
and https://httpstat.us/503

**Note:** Configured URLs are only polled when prometheus scrapes the metrics, allowing you to decide how often you
would like sites to be polled from the prometheus side. Depending on the number of targets being polled, you may need to
increase the scrape timeout in the prometheus scrape configuration using the
[scrape_timeout](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#scrape_config) option.

### Example of polling custom urls

Edit the [deployment.yaml](deploy/manifests/deployment.yaml) to include the following (as an example):

```yaml
containers:
  - name: vdpcr
    image: logikone/vdpcr:latest
    args:
      - --target https://google.com
      - --target https://reddit.com
```

## Prometheus

Here is an example scrape configuration for the exporter:

```yaml
  - job_name: vdpcr
    metrics_path: /metrics
    relabel_configs:
      - source_labels: [ __meta_kubernetes_pod_label_app_kubernetes_io_name ]
      - regex: vdpcr
      - action: keep
```

### Exposed Metrics

The following metrics are exposed given the default configuration:

```text
# HELP external_url_response_ms 
# TYPE external_url_response_ms histogram
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.005"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.01"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.025"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.05"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.1"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.25"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="0.5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="1"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="2.5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="10"} 0
external_url_response_ms_bucket{url="https://httpstat.us/200",le="+Inf"} 256
external_url_response_ms_sum{url="https://httpstat.us/200"} 67003
external_url_response_ms_count{url="https://httpstat.us/200"} 256
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.005"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.01"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.025"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.05"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.1"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.25"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="0.5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="1"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="2.5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="5"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="10"} 0
external_url_response_ms_bucket{url="https://httpstat.us/503",le="+Inf"} 256
external_url_response_ms_sum{url="https://httpstat.us/503"} 67620
external_url_response_ms_count{url="https://httpstat.us/503"} 256

# HELP external_url_up 
# TYPE external_url_up gauge
external_url_up{url="https://httpstat.us/200"} 0
external_url_up{url="https://httpstat.us/503"} 1
```

## Grafana

There is a sample Grafana dashboard included in [deploy/grafana/dashboard.json](deploy/grafana/dashboard.json)

![Dashboard](dashboard.png)

### CLI Reference

```shell
Usage of /vdpcr:
      --listen-address string   The address to listen on for HTTP requests. (default ":8080")
      --log-level string        log level to set. (debug, error, fatal, info, warn) (default "info")
      --target stringArray      target(s) to scrape. can be specified multiple times (default [https://httpstat.us/200,https://httpstat.us/503])
      --timeout duration        http client timeout (default 10s)
```

