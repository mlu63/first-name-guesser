// ideally it would be best to avoid testing in main - however, I am acting as if
// I am a test engineer and I don't want to move around the source code even if it
// makes things a bit easier for me
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// verify that the status code and response are adequate
func TestNameHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(nameHandler)
	handler.ServeHTTP(rr, req)

	// expect status code OK (200)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code of '%v', received '%v'\n",
			http.StatusOK, status)
	}

	// expect an empty first name since no request was actually passed in
	expected := `{"first_name": ""}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body of '%v', received '%v'\n",
			expected, rr.Body.String())
	}

}
