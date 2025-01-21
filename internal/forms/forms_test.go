package forms

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	if !form.IsValid() {
		t.Error("form shows as invalid when it should be valid")
	}

	form.Errors.AddErr("field", "err")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}
}

func TestForm_Required(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.Required("a", "b", "c")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	postedData := url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")
	postedData.Add("c", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if !form.IsValid() {
		t.Error("form shows as invalid when it should be valid")
	}

	postedData = url.Values{}
	postedData.Add("a", "a")
	postedData.Add("b", "a")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)
	form.Required("a", "b", "c")

	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}
}

func TestForm_CheckLength(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	form := New(r.PostForm)

	form.CheckLength("a", 3)
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	postedData := url.Values{}
	postedData.Add("a", "aaaa")

	r, _ = http.NewRequest("POST", "/whatever", nil)

	r.PostForm = postedData
	form = New(r.PostForm)

	form.CheckLength("a", 3)
	if !form.IsValid() {
		t.Error("form shows as invalid when it should be valid")
	}

	postedData = url.Values{}
	postedData.Add("a", "aa")

	r, _ = http.NewRequest("POST", "/whatever", nil)
	r.PostForm = postedData
	form = New(r.PostForm)

	form.CheckLength("a", 3)
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}
}

func TestForm_IsEmail(t *testing.T) {
	r := httptest.NewRequest("POST", "/whatever", nil)

	postedData := url.Values{}
	postedData.Add("email", "a@a.com")

	r.PostForm = postedData
	form := New(r.PostForm)

	form.IsEmail("email")
	if !form.IsValid() {
		t.Error("form shows as invalid when it should be valid")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "a")

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "a@")

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "a@a")

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "@")

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}

	r, _ = http.NewRequest("POST", "/whatever", nil)

	postedData = url.Values{}
	postedData.Add("email", "a@a.")

	r.PostForm = postedData
	form = New(r.PostForm)

	form.IsEmail("email")
	if form.IsValid() {
		t.Error("form shows as valid when it should be invalid")
	}
}

func TestError_AddErr(t *testing.T) {
	errs := errors{}
	errs.AddErr("field", "err")

	if errs.Err("field") != "err" {
		t.Error("expected 'err'")
	}
}

func TestError_Err(t *testing.T) {
	errs := errors{}
	errs.AddErr("field", "err")

	if errs.Err("invalid-field") != "" {
		t.Error("expected empty string")
	}

	if errs.Err("field") != "err" {
		t.Error("expected 'err'")
	}
}
