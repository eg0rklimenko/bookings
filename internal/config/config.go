package config

import (
	"log"

	"github.com/CloudyKit/jet"
	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache       bool
	TemplateCache  *jet.Set
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	CacheSecure    bool
	SessionManager *scs.SessionManager
}
