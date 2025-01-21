package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"single-room", "/single-room", "GET", []postData{}, http.StatusOK},
	{"double-bed-room", "/double-bed-room", "GET", []postData{}, http.StatusOK},
	{"two-bed-room", "/two-bed-room", "GET", []postData{}, http.StatusOK},
	{"family-room", "/family-room", "GET", []postData{}, http.StatusOK},
	{"make-reservation", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "GET", []postData{}, http.StatusOK},

	{"post-search-availability", "/search-availability", "POST", []postData{
		{"start", "15-06-2024"},
		{"end", "19-06-2024"},
		{"people", "1"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "POST", []postData{
		{"first_name", "name"},
		{"last_name", "name"},
		{"date_from", "15-06-2024"},
		{"date_to", "19-06-2024"},
		{"phone", "+7 (999)-999-99-99"},
		{"email", "email@example.com"},
		{"room", "1"},
	}, http.StatusOK}}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range tests {
		if test.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status code %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		} else {
			values := url.Values{}
			for _, param := range test.params {
				values.Add(param.key, param.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status code %d, but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}
