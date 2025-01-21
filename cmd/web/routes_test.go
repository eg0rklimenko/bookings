package main

import (
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoutes(t *testing.T) {
	mux := routes()

	switch T := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("type is not *chi.Mux, type is %T", T)
	}
}
