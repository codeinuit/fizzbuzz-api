package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/codeinuit/fizzbuzz-api/pkg/database"
	"github.com/codeinuit/fizzbuzz-api/pkg/database/mysql"
	logger "github.com/codeinuit/fizzbuzz-api/pkg/log"
	"github.com/codeinuit/fizzbuzz-api/pkg/log/logrus"

	"github.com/gin-gonic/gin"
)

// FizzBuzz represents the main structure
type FizzBuzz struct {
	engine *gin.Engine
	srv    *http.Server
	log    logger.Logger

	// allows gracefull shutdown
	quit chan os.Signal
}

// setupRouter init the main structure and the http router as well
func setupRouter() (fb *FizzBuzz, err error) {
	l := logrus.NewLogrusLogger()
	_, isDebug := os.LookupEnv("DEBUG")
	if !isDebug {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	fb = &FizzBuzz{
		engine: r,
		quit:   make(chan os.Signal),
		log:    l,
	}

	return fb, nil
}

// Run will run main program along the http server.
// It should be run as goroutine and stopped using Stop function
func (fb *FizzBuzz) Run(port string) {
	fb.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: fb.engine,
	}

	fb.log.Info("running server on port ", port)
	err := fb.srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		fb.log.Error(err.Error())
	}
}

// Stop is used to stop the main program and its http server properly
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

// initRoutes is used to init the base routes and attach them
// to the router
func initRoutes(fb *FizzBuzz, db database.Database) {
	h := handlers{
		log: fb.log,
		db:  db,
	}

	fb.engine.GET("/health", h.healthcheck)
	fb.engine.GET("/fizzbuzz", h.fizzbuzz)
	fb.engine.GET("/stats", h.stats)
}

func main() {
	fb, err := setupRouter()
	if err != nil {
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil {
		fb.log.Error(err.Error())
		os.Exit(1)
	}

	db, err := mysql.InitDatabase()
	if err != nil {
		fb.log.Error(err.Error())
		os.Exit(1)
	}

	initRoutes(fb, db)

	signal.Notify(fb.quit, syscall.SIGINT, syscall.SIGTERM)
	go fb.Run(port)
	<-fb.quit

	fb.Stop()
}
