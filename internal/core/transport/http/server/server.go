package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	core_http_midlware "github.com/tiotsyodev/url-shortener.git/internal/core/transport/http/middleware.go"
)

type HTTPServer struct {
	Mux    *http.ServeMux
	Config Config
	Log    *slog.Logger
	Middlewares []core_http_midlware.Middleware
}

func NewHttpServer(
	mux    *http.ServeMux,
	cfg Config,
	log    *slog.Logger,
	middlewares []core_http_midlware.Middleware,
) *HTTPServer {
	return &HTTPServer {
		Mux:    mux,
		Config: cfg,
		Log:    log,
		Middlewares: middlewares,
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_midlware.ChainMiddleware(h.Mux, h.Middlewares...)
	server := http.Server{
		Addr: h.Config.Host + ":" + strconv.Itoa(h.Config.Port),
		Handler: mux,	
		IdleTimeout: h.Config.IdleTimeout,
		ReadTimeout: h.Config.Timeout,
	}

	exitCh := make(chan error, 1)

	go func() {
		defer close(exitCh)
		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			exitCh <- err
		}
	}()

	select {
	case err := <- exitCh:
		if err != nil {
			return fmt.Errorf("listen and server HTTP: %w", err)
		}
	case <- ctx.Done():
		shutDonwCtx, cancel := context.WithTimeout(context.Background(), h.Config.IdleTimeout)
		defer cancel()

		if err := server.Shutdown(shutDonwCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown HTTP server: %w", err)
		}
	}

	return nil
}