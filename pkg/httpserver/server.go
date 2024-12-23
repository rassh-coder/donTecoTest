package httpserver

import (
	"context"
	"donTecoTest/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Server(r *gin.Engine, cfg *config.Config) error {
	srv := &http.Server{
		Addr:    cfg.Host.Port,
		Handler: r,
	}
	go func() {
		fmt.Println(fmt.Sprintf("%s:%s", cfg.Host.Host, cfg.Host.Port))
		// service connections
		if err := r.Run(fmt.Sprintf("%s:%s", cfg.Host.Host, cfg.Host.Port)); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

	return nil
}
