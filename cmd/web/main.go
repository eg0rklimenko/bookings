package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/handlers"
	"github.com/mo0Oonnn/bookings/internal/helpers"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"
)

const portNumber = ":8080"

var (
	appConfig      config.AppConfig
	sessionManager *scs.SessionManager
	infoLog        *log.Logger
	errorLog       *log.Logger
)

// main is the main application module
func main() {
	err := runApp()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func runApp() error {
	gob.Register(models.Reservation{})

	appConfig.CacheSecure = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog

	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour * 24
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = appConfig.CacheSecure

	appConfig.SessionManager = sessionManager

	tc := render.CreateTemplateCache()

	appConfig.TemplateCache = tc
	appConfig.UseCache = true

	repo := handlers.NewRepo(&appConfig)
	handlers.SetRepo(repo)

	render.SetConfig(&appConfig)
	helpers.SetConfig(&appConfig)
	return nil
}
