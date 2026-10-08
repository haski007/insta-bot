package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var latencyBuckets = []float64{0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2, 3, 5, 8, 13, 20, 30, 50}

var (
	telegramUpdatesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "insta_bot",
		Name:      "telegram_updates_total",
		Help:      "Telegram updates received by the polling loop.",
	}, []string{"kind"})

	messageRoutesTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "insta_bot",
		Name:      "message_routes_total",
		Help:      "How incoming messages were routed (including ignored loaders).",
	}, []string{"route"})

	commandDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "command_duration_seconds",
		Help:      "Duration of a Telegram bot command handler.",
		Buckets:   latencyBuckets,
	}, []string{"command"})

	messageHandlerDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "message_handler_duration_seconds",
		Help:      "Duration of a non-command message handler.",
		Buckets:   latencyBuckets,
	}, []string{"kind"})

	llmRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "llm_request_duration_seconds",
		Help:      "Duration of a request to an LLM/classifier provider, including response parsing.",
		Buckets:   latencyBuckets,
	}, []string{"provider", "operation", "status"})

	externalRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "external_request_duration_seconds",
		Help:      "Duration of an outbound HTTP call to an integration (instloader, TikTok, ARC, ...).",
		Buckets:   latencyBuckets,
	}, []string{"service", "operation", "status"})

	mediaDownloadDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "media_download_duration_seconds",
		Help:      "Duration of downloading a media file to send to Telegram.",
		Buckets:   latencyBuckets,
	}, []string{"source", "status"})

	telegramSendDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "telegram_send_duration_seconds",
		Help:      "Duration of Telegram Bot API send calls.",
		Buckets:   latencyBuckets,
	}, []string{"kind", "status"})

	monitorRunsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: "insta_bot",
		Name:      "monitor_runs_total",
		Help:      "Background monitor ticks that actually did work (ARC, newsletter).",
	}, []string{"monitor", "status"})

	redisOperationDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "redis_operation_duration_seconds",
		Help:      "Duration of selected Redis operations.",
		Buckets:   []float64{0.001, 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2, 5},
	}, []string{"operation", "status"})

	redisReadonly = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "insta_bot",
		Name:      "redis_readonly",
		Help:      "1 if Redis is currently read-only (write probe failed), else 0.",
	})

	anglicismHandlerDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "insta_bot",
		Name:      "anglicism_handler_duration_seconds",
		Help:      "Duration of the whole anglicism handler, including reply delivery.",
		Buckets:   latencyBuckets,
	}, []string{"outcome"})
)

func status(err error) string {
	if err != nil {
		return "error"
	}
	return "ok"
}

func IncUpdate(kind string) {
	telegramUpdatesTotal.WithLabelValues(kind).Inc()
}

func IncRoute(route string) {
	messageRoutesTotal.WithLabelValues(route).Inc()
}

func ObserveCommand(command string, start time.Time) {
	commandDuration.WithLabelValues(command).Observe(time.Since(start).Seconds())
}

func ObserveMessageHandler(kind string, start time.Time) {
	messageHandlerDuration.WithLabelValues(kind).Observe(time.Since(start).Seconds())
}

func ObserveLLM(provider, operation string, start time.Time, err error) {
	llmRequestDuration.WithLabelValues(provider, operation, status(err)).Observe(time.Since(start).Seconds())
}

func ObserveExternal(service, operation string, start time.Time, err error) {
	externalRequestDuration.WithLabelValues(service, operation, status(err)).Observe(time.Since(start).Seconds())
}

func ObserveMediaDownload(source string, start time.Time, err error) {
	mediaDownloadDuration.WithLabelValues(source, status(err)).Observe(time.Since(start).Seconds())
}

func ObserveTelegramSend(kind string, start time.Time, err error) {
	telegramSendDuration.WithLabelValues(kind, status(err)).Observe(time.Since(start).Seconds())
}

func ObserveMonitor(monitor string, err error) {
	monitorRunsTotal.WithLabelValues(monitor, status(err)).Inc()
}

func ObserveRedis(operation string, start time.Time, err error) {
	redisOperationDuration.WithLabelValues(operation, status(err)).Observe(time.Since(start).Seconds())
}

func SetRedisReadonly(readonly bool) {
	if readonly {
		redisReadonly.Set(1)
		return
	}
	redisReadonly.Set(0)
}

func ObserveAnglicismHandler(outcome string, start time.Time) {
	anglicismHandlerDuration.WithLabelValues(outcome).Observe(time.Since(start).Seconds())
}
