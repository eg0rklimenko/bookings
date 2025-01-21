package main

import (
	"net/http"
	"testing"
)

func testMiddlewares(h http.Handler, t *testing.T) {
	switch T := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Errorf("type is not http.Handler, but is %T", T)
	}

}

func TestNoSurf(t *testing.T) {
	var handler Handler

	h := NoSurf(&handler)

	testMiddlewares(h, t)
}

func TestSessionLoad(t *testing.T) {
	var handler Handler

	h := SessionLoad(&handler)

	testMiddlewares(h, t)
}
