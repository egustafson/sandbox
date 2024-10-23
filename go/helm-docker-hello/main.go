package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Masterminds/log-go"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port int `json:"port" yaml:"port"`
}

func apiRun(ctx context.Context, config *Config) (err error) {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Port),
		Handler: router,
	}
	log.Infof("listening on port %d", config.Port)
	idleConnsClosed := make(chan struct{})
	go func() {
		<-ctx.Done() // block until context canceled
		const timeoutTime = 5 * time.Second
		timeoutCtx, cancel := context.WithTimeout(context.Background(), timeoutTime)
		defer cancel()
		if err := s.Shutdown(timeoutCtx); err != nil {
			log.Warnf("http server Shutdown: %v", err)
			s.Close()
		}
		close(idleConnsClosed)
	}()

	if err = s.ListenAndServe(); err != http.ErrServerClosed {
		log.Errorf("http server ListenAndServer: %v", err)
	} else {
		err = nil // don't report the server closing, it's normal
	}

	<-idleConnsClosed
	log.Info("http server shutdown complete")
	return
}

func main() {
	config := &Config{
		Port: 8080,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		log.Infof("received signal: %s", sig.String())
		cancel()
	}()

	apiRun(ctx, config)
	<-ctx.Done()
	log.Info("done.")
}
