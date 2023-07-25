package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
	"vaja1/API"
	"vaja1/DB/MariaDB"
	"vaja1/Logic"
)

func main() {
	// Create a new MariaDB object with info to connect to the database
	db := &MariaDB.MariaDB{
		User:     getEnvStr("DBUSER", "root"),
		Pass:     getEnvStr("DBPASS", "superSecretPass"),
		IP:       getEnvStr("DBIP", "127.0.0.1"),
		Port:     getEnvInt("DBPORT", 3306),
		Database: getEnvStr("DBNAME", "test"),
	}
	err := db.Init()
	if err != nil {
		fmt.Println(err.Error())
	}

	logic := Logic.NewController(db)

	// Create a router object
	var router Router
	router.engine = gin.Default()
	router.api = API.NewController(logic)

	// Register HTTP routes
	err = router.registerRoutes()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Create 2 channels and connect one of them to the system exit signal
	quit := make(chan os.Signal, 0)
	done := make(chan bool, 0)
	signal.Notify(quit, os.Interrupt)

	// Define the port, handler and timeouts for the HTTP server
	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// In a separate thread, wait for an interrupt signal and close the http server
	go func() {
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			fmt.Println(err.Error())
		}
		close(done)
	}()

	// Start a new HTTP server in a separate thread
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
		os.Exit(1)

	}()

	// Wait for all programs to finish executing, or an interrupt
	<-done
}

// getEnvStr returns the string value of the environment variable named by the `key`.
// If the variable is not present in the environment, the fallback value is returned.
func getEnvStr(key, fallback string) string {
	if value, found := os.LookupEnv(key); found {
		return value
	}
	return fallback
}

// getEnvInt returns the int value of the environment variable named by the `key`.
// If the variable is not present in the environment, the fallback value is returned.
func getEnvInt(key string, fallback int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}
	return fallback
}
