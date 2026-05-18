package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type HTTPServer struct {
	Mux    *http.ServeMux
	Config Config
	Log    *slog.Logger
}

func NewHttpServer(
	mux    *http.ServeMux,
	cfg Config,
	log    *slog.Logger,
) *HTTPServer {
	return &HTTPServer {
		Mux:    mux,
		Config: cfg,
		Log:    log,
		
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	server := http.Server{
		Addr: h.Config.Host + ":" + strconv.Itoa(h.Config.Port),
		Handler: h.Mux,	
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