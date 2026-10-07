package listener

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var latencyBuckets = []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2, 3, 5, 8, 13, 20, 30, 50}

var (
	anglicismHandlerDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "anglicism_handler_duration_seconds",
		Help:      "Duration of the whole anglicism handler, including reply delivery.",
		Buckets:   latencyBuckets,
	}, []string{"outcome"})

	llmRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "llm_request_duration_seconds",
		Help:      "Duration of a request to an LLM/classifier provider, including response parsing.",
		Buckets:   latencyBuckets,
	}, []string{"provider", "operation", "status"})
)

const (
	anglicismOutcomeSkipped       = "skipped"
	anglicismOutcomeNotSubscribed = "not_subscribed"
	anglicismOutcomeIgnoredUser   = "ignored_user"
	anglicismOutcomeLLMError      = "llm_error"
	anglicismOutcomeNoAnglicism   = "no_anglicism"
	anglicismOutcomeReplied       = "replied"
)

func observeLLMRequest(provider, operation string, start time.Time, err error) {
	status := "ok"
	if err != nil {
		status = "error"
	}
	llmRequestDuration.WithLabelValues(provider, operation, status).Observe(time.Since(start).Seconds())
}
