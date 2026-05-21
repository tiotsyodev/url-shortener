package main

import (
	"context"
	"log/slog"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	core_cache "github.com/tiotsyodev/url-shortener.git/internal/core/cache"
	core_logger "github.com/tiotsyodev/url-shortener.git/internal/core/logger"
	core_metrics "github.com/tiotsyodev/url-shortener.git/internal/core/metrics"
	core_repo "github.com/tiotsyodev/url-shortener.git/internal/core/repo"
	core_http_midlware "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/middleware.go"
	core_http_server "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/server"
	redirect_service "github.com/tiotsyodev/url-shortener.git/internal/features/redirect/service"
	redirect_transport "github.com/tiotsyodev/url-shortener.git/internal/features/redirect/trasnport"
	stats_repo "github.com/tiotsyodev/url-shortener.git/internal/features/stat/repo"
	stats_service "github.com/tiotsyodev/url-shortener.git/internal/features/stat/service"
	stat_transport "github.com/tiotsyodev/url-shortener.git/internal/features/stat/transport"
	url_repo "github.com/tiotsyodev/url-shortener.git/internal/features/url/repo"
	url_service "github.com/tiotsyodev/url-shortener.git/internal/features/url/service"
	url_transport "github.com/tiotsyodev/url-shortener.git/internal/features/url/transport"
)

func main() {
	_ = godotenv.Load()

	httpServerConfig := core_http_server.ConfigMustLoad()
	logger := core_logger.MustSetupLogger()

	dbpool, err := core_repo.NewConnectionPool(context.Background(), core_repo.NewConfigMustLoad())
	if err != nil {
		logger.Error("initializing conn pool", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		os.Exit(1)
	}
	redisClient, err := core_cache.NewRedisClient(core_cache.NewConfigMust())
	if err != nil {
		logger.Error("initializing redis db", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		os.Exit(1)
	}

	metrics := core_metrics.MustRegisterMetrics()

	urlRepo := url_repo.NewUrlRepository(dbpool)
	statsRepo := stats_repo.NewStatRepo(dbpool)
	
	urlSvc := url_service.NewUserService(urlRepo, metrics)
	statSvc := stats_service.NewStatService(statsRepo)
	redirectSvc := redirect_service.NewRedirectService(urlRepo, statsRepo, redisClient, metrics)

	urlHandler := url_transport.NewUrlHandler(logger, urlSvc)
	redirectHandler := redirect_transport.NewRedirectHandler(logger, redirectSvc, statSvc)
	statHandler := stat_transport.NewStatHandler(logger, statSvc)
	
	srv := core_http_server.NewHttpServer(http.NewServeMux(), httpServerConfig, logger, []core_http_midlware.Middleware{core_http_midlware.MetricsMiddleware(metrics)})
	router := core_http_server.NewRouter(srv.Mux)

	srv.Mux.Handle("GET /metrics", promhttp.Handler())
	router.RegisterRoutes(statHandler.GetRoutes()...)
	router.RegisterRoutes(urlHandler.GetRoutes()...)
	router.RegisterRoutes(redirectHandler.GetRoutes()...)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
    defer cancel()

    if err := srv.Run(ctx); err != nil {
		logger.Error("server error", slog.Attr{Key: "error", Value: slog.StringValue(err.Error())})
		os.Exit(1)
	}
	dbpool.Close()
	redisClient.RDB.Close()
}