package render

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/mo0Oonnn/bookings/internal/config"
	"github.com/mo0Oonnn/bookings/internal/models"
)

var (
	sessionManager *scs.SessionManager
	testAppConfig  config.AppConfig
)

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	// testAppConfig.CacheSecure = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testAppConfig.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testAppConfig.ErrorLog = errorLog

	sessionManager = scs.New()
	sessionManager.Lifetime = time.Hour * 24
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = false

	testAppConfig.SessionManager = sessionManager

	appConfig = &testAppConfig

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {
}

func (tw *myWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
