package render

import (
	"bytes"
	"errors"
	"log"
	"net/http"

	"github.com/CloudyKit/jet"
	"github.com/justinas/nosurf"

	"github.com/mo0Oonnn/bookings/internal/config"
)

var (
	appConfig       *config.AppConfig
	pathToTemplates = "./templates"
)

func SetConfig(a *config.AppConfig) {
	appConfig = a
}

func AddDefaultData(td jet.VarMap, r *http.Request) jet.VarMap {
	td.Set("success", appConfig.SessionManager.PopString(r.Context(), "success"))
	td.Set("error", appConfig.SessionManager.PopString(r.Context(), "error"))
	td.Set("warning", appConfig.SessionManager.PopString(r.Context(), "warning"))

	td.Set("csrf_token", nosurf.Token(r))
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td jet.VarMap) error {
	var tc *jet.Set

	if appConfig.UseCache {
		tc = appConfig.TemplateCache
	} else {
		tc = CreateTemplateCache()
	}

	t, err := tc.GetTemplate(tmpl)
	if err != nil {
		log.Println("cannot get template from template cache: ", err)
		return errors.New("cannot get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	if err = t.Execute(buf, td, nil); err != nil {
		log.Println(err)
		return err
	}

	if _, err = buf.WriteTo(w); err != nil {
		log.Println("error writing template to browser: ", err)
		return err
	}
	return nil
}

func CreateTemplateCache() *jet.Set {
	views := jet.NewHTMLSet(pathToTemplates)
	return views
}
