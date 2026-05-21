package core_http_midlware

import (
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	core_metrics "github.com/tiotsyodev/url-shortener.git/internal/core/metrics"
)

type responseWriter struct {
    http.ResponseWriter
    statusCode int    
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code              
    rw.ResponseWriter.WriteHeader(code) 
}

type Middleware func(h http.Handler) http.Handler

func ChainMiddleware(handler http.Handler, mds ...Middleware) http.Handler {
	if len(mds) == 0 {
		return handler
	} 
	for i := len(mds) - 1; i >= 0; i-- {
		handler = mds[i](handler)
	}
	return handler
}

func MetricsMiddleware(metrics core_metrics.Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			promtimer := prometheus.NewTimer(metrics.ResolutionHisto.WithLabelValues(r.Method, r.Pattern))

			wrapped := responseWriter{w, http.StatusOK}
			next.ServeHTTP(&wrapped, r)

			metrics.HTTPReqCounter.WithLabelValues(r.Method, r.Pattern, strconv.Itoa(wrapped.statusCode)).Inc()
			promtimer.ObserveDuration()
        })

	}
}