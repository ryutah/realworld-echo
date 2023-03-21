package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/realworld-echo/api/rest/gen"
)

func Start(s *Server) {
	e := echo.New()

	gen.RegisterHandlers(e, s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := startServerwithGracefulShutdown(fmt.Sprintf(":%s", port), e, func() {
		log.Println("Server shutdown")
	}); err != nil {
		log.Fatal(err)
	}
}

func startServerwithGracefulShutdown(addr string, h http.Handler, cleanFunc func()) error {
	srv := &http.Server{
		Handler: h,
		Addr:    addr,
	}

	errChan := make(chan error)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	go func() {
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
		defer stop()
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if cleanFunc != nil {
			cleanFunc()
		}
		errChan <- srv.Shutdown(ctx)
	}()

	return <-errChan
}
