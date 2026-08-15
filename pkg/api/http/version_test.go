package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	srv := NewMockServer()
	handler := http.HandlerFunc(srv.versionHandler)

	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	body := rr.Body.String()

	for _, field := range []string{"unknown", "buildtime", "runtime"} {
		r := regexp.MustCompile(fmt.Sprintf("(?m:%s)", field))
		if !r.MatchString(body) {
			t.Fatalf("handler returned unexpected body:\ngot \n%v \nwant field %s",
				body, field)
		}
	}
}
