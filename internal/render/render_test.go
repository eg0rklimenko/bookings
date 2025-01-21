package render

import (
	"net/http"
	"testing"

	"github.com/CloudyKit/jet"

	"github.com/mo0Oonnn/bookings/internal/config"
)

func TestAddDefaultData(t *testing.T) {
	td := make(jet.VarMap)

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	sessionManager.Put(r.Context(), "error", "some error")
	result := AddDefaultData(td, r)

	if result == nil {
		t.Error("failed to add default data to template data")
	}
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	tc := CreateTemplateCache()
	if tc == nil {
		t.Error("tc is nil")
	}

	appConfig.UseCache = false
	appConfig.TemplateCache = tc
	appConfig.SessionManager = sessionManager

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	if err := RenderTemplate(&ww, r, "home.jet", jet.VarMap{}); err != nil {
		t.Error("error writing template to browser:", err)
	}
	if err := RenderTemplate(&ww, r, "non-existent.jet", jet.VarMap{}); err == nil {
		t.Error("rendered template that does not exist")
	}
}

func TestSetConfig(t *testing.T) {
	SetConfig(&config.AppConfig{})
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	if tc := CreateTemplateCache(); tc == nil {
		t.Error("tc is nil")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/url", nil)
	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, err = sessionManager.Load(ctx, r.Header.Get("X-Session"))
	if err != nil {
		return nil, err
	}
	r = r.WithContext(ctx)

	return r, nil
}
