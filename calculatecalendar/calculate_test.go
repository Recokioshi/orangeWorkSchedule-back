package calculatecalendar

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func checkAndThrowForExpectedResponse(t *testing.T, status int, expected int) {
	if status != expected {
		t.Errorf("handler returned wrong status code. Got %v, expected %v", status, expected)
	}
}

func TestIndexHandlerPass(t *testing.T) {
	requestBody := strings.NewReader(`
	{
		"Workers": 0,
		"Calendar": [
				{
						"Active": true,
						"Whitelist": [],
						"Blacklist": [],
						"Shifts": {
								"0": {
										"StartTime": 0,
										"EndTime": 0
								}
						}
				}
		]
	}
	`)

	req, err := http.NewRequest("POST", "/", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)
	checkAndThrowForExpectedResponse(t, rr.Code, http.StatusOK)
}

func TestIndexHandlerEmptyBody(t *testing.T) {
	req, err := http.NewRequest("POST", "/", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)
	checkAndThrowForExpectedResponse(t, rr.Code, http.StatusBadRequest)

}

func TestIndexHandlerInvalidPath(t *testing.T) {
	req, err := http.NewRequest("POST", "/invalid", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(rr, req)

	checkAndThrowForExpectedResponse(t, rr.Code, http.StatusNotFound)

}

func TestIndexHandlerFailOnInvalidBodyStructure(t *testing.T) {
	requestBody := strings.NewReader(`
	{
		"Workers": 0,
		"BadAdditionalField",
		"Calendar": [
				{
						"Active": true,
						"Whitelist": [],
						"Blacklist": [],
						"Shifts": {
								"0": {
										"StartTime": 0,
										"EndTime": 0
								}
						}
				}
		]
	}
	`)

	req, err := http.NewRequest("POST", "/", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)
	handler.ServeHTTP(rr, req)
	checkAndThrowForExpectedResponse(t, rr.Code, http.StatusBadRequest)
}
