package core_metrics

import "github.com/prometheus/client_golang/prometheus"


type Metrics struct {
	HTTPReqCounter *prometheus.CounterVec
	ResolutionHisto *prometheus.HistogramVec
	CreatedLinksCounter prometheus.Counter
	RedirectCounter *prometheus.CounterVec
}

func MustRegisterMetrics() (Metrics) {
	HttpCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "url_shortener",
		Subsystem: "http",
		Name: "request_total",
		Help: "total count of http requests",
	}, []string{"method", "path", "status"})

	ResolutionTimeHisto := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "url_shortener",
		Subsystem: "http",
		Name: "request_duration_seconds",
		Help: "hystogram of resolution time for http requests.",
		Buckets: prometheus.DefBuckets,		
	}, []string{"method", "path"})

	CreatedUrlsCounter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "url_shortener",
		Name: "created_urls_total",
		Help: "Counter of created urls.",
	})

	RedirectsCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "url_shortener",
		Subsystem: "http",
		Name: "redirects_total",
		Help: "total count of redirects",
	}, []string{"source"})

	prometheus.MustRegister(HttpCounter, ResolutionTimeHisto, CreatedUrlsCounter, RedirectsCounter)

	return Metrics{
		HTTPReqCounter: HttpCounter,
		ResolutionHisto: ResolutionTimeHisto,
		CreatedLinksCounter: CreatedUrlsCounter,
		RedirectCounter: RedirectsCounter,
	}
}

