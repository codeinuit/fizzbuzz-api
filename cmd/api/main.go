package main

import (
	logger "codeinuit/fizzbuzz-api/pkg/log"
	"codeinuit/fizzbuzz-api/pkg/log/logrus"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type FizzBuzz struct {
	engine *gin.Engine
	srv    *http.Server
	log    logger.Logger

	// allows gracefull shutdown
	quit chan os.Signal
}

func setupRouter() *gin.Engine {
	_, isDebug := os.LookupEnv("DEBUG")
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	return r
}

func (fb *FizzBuzz) Run(port string) {
	fb.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: fb.engine,
	}

	fb.log.Info("running server on port ", port)
	err := fb.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func (fb FizzBuzz) Stop() {
	fb.log.Warn("stop signal catched, closing server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := fb.srv.Shutdown(ctx)
	if err != nil {
		fb.log.Error("could not stop properly server : ", err.Error())
	}

	fb.log.Info("server closed, exiting")
}

func main() {
	l := logrus.NewLogrusLogger()

	fb := FizzBuzz{
		engine: setupRouter(),
		quit:   make(chan os.Signal),
		log:    l,
	}

	h := handlers{
		log: l,
	}

	fb.engine.GET("/health", h.healthcheck)

	signal.Notify(fb.quit, syscall.SIGINT, syscall.SIGTERM)
	go fb.Run(os.Getenv("PORT"))
	<-fb.quit

	fb.Stop()
}
