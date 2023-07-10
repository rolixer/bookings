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

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	// Get handlers
	{"home", "/", "get", []postData{}, http.StatusOK},
	{"about", "/about", "get", []postData{}, http.StatusOK},
	{"generals-quarters", "/generals-quarters", "get", []postData{}, http.StatusOK},
	{"majors-suite", "/majors-suite", "get", []postData{}, http.StatusOK},
	{"search-availability", "/search-availability", "get", []postData{}, http.StatusOK},
	{"contact", "/contact", "get", []postData{}, http.StatusOK},
	{"make-reservations", "/make-reservation", "get", []postData{}, http.StatusOK},
	// POST handlers
	{"post-search-availability", "/search-availability", "post", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-03"},
	}, http.StatusOK},
	{"post-search-availability-json", "/search-availability-json", "post", []postData{
		{key: "start", value: "2023-01-01"},
		{key: "end", value: "2023-01-03"},
	}, http.StatusOK},
	{"post-make-reservation", "/make-reservation", "post", []postData{
		{key: "first_name", value: "John"},
		{key: "last_name", value: "Smith"},
		{key: "email", value: "me@here.com"},
		{key: "phone", value: "123 456 789"},
	}, http.StatusOK}}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, test := range theTests {
		switch test.method {
		case "get":
			resp, err := ts.Client().Get(ts.URL + test.url)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		case "post":
			values := url.Values{}
			for _, x := range test.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+test.url, values)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != test.expectedStatusCode {
				t.Errorf("for %s expected status %d but got %d", test.name, test.expectedStatusCode, resp.StatusCode)
			}
		default:
			//do nothig
		}
	}
}
