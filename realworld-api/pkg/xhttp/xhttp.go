package xhttp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/multierr"
)

type CleanupFunc func() error

func RunWithGraceful(addr string, h http.Handler, cleanFunc func(context.Context) error) error {
	srv := &http.Server{
		Handler:           h,
		Addr:              addr,
		ReadHeaderTimeout: 60 * time.Second,
	}

	errChan := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- fmt.Errorf("failed to listen server: %w ", err)
		}
	}()

	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
		defer stop()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			errChan <- fmt.Errorf("failed to shutdown server: %w", err)
		}
		if cleanFunc != nil {
			if err := cleanFunc(ctx); err != nil {
				errChan <- fmt.Errorf("failed to cleanup func: %w", err)
			}
		}
		close(errChan)
	}()

	var err error
	for e := range errChan {
		err = multierr.Append(err, e)
	}
	return err
}
