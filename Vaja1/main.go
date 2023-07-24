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
	db := &MariaDB.MariaDB{
		User:     getEnvStr("DBUSER", "root"),
		Pass:     getEnvStr("DBPASS", "superSecretPass"),
		IP:       getEnvStr("DBIP", "127.0.0.1"),
		Port:     getEnvInt("DBPORT", 3306),
		Database: getEnvStr("DBNAME", "test"),
	} // Ce podatki za povezavo z db niso v envVar, uporabimo default vrednosti
	err := db.Init()
	if err != nil {
		fmt.Println(err.Error())
	}

	logic := Logic.NewController(db)

	// Kreiramo naš router objekt
	var router Router
	router.engine = gin.Default()
	router.api = API.NewController(logic)

	// Registriramo HTTP REST API povezave
	err = router.registerRoutes()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Naredimo 2 kanala in enega od njih povežemo na sistemski exit signal
	quit := make(chan os.Signal, 0)
	done := make(chan bool, 0)
	signal.Notify(quit, os.Interrupt)

	// Definiramo port, handler in timeoute za HTTP server
	srv := &http.Server{
		Addr:         ":80",
		Handler:      router.engine,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// V ločenem threadu čakamo na exit signal in nato izklopimo http server
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

	// V ločenem threadu zaženemo HTTP server
	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err.Error())
		}
		os.Exit(1)

	}()

	// Čakamo na konec izvajanja. Vsi deli programa so sedaj zagnani v ločenih threadih.
	// Dalje od tod pridemo samo če program izklopimo ali se zruži
	<-done

}

func getEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if value == 0 || err != nil {
		return fallback
	}
	return value
}
