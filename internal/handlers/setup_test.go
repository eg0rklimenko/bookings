package handlers

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/justinas/nosurf"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/models"
	"github.com/mo0Oonnn/bookings/internal/render"
)

var (
	appConfig       config.AppConfig
	sessionManager  *scs.SessionManager
	pathToTemplates = "./../../templates/"
)

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	appConfig.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	appConfig.ErrorLog = errorLog

	appConfig.CacheSecure = false

	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour * 24
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = appConfig.CacheSecure

	appConfig.SessionManager = sessionManager

	tc := CreateTestTemplateCache()

	appConfig.TemplateCache = tc
	appConfig.UseCache = true

	repo := NewRepo(&appConfig)
	SetRepo(repo)

	render.SetConfig(&appConfig)

	mux := chi.NewMux()

	mux.Use(middleware.Recoverer)
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/single-room", Repo.SingleRoom)
	mux.Get("/double-bed-room", Repo.DoubleBedRoom)
	mux.Get("/two-bed-room", Repo.TwoBedRoom)
	mux.Get("/family-room", Repo.FamilyRoom)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)

	fileserver := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   appConfig.CacheSecure,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return sessionManager.LoadAndSave(next)
}

func CreateTestTemplateCache() *jet.Set {
	views := jet.NewHTMLSet(pathToTemplates)
	return views
}
