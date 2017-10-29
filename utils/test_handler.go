package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// CheckGetHandler tests reponse body and status code of a gin `GET` handler
func CheckGetHandlerResp(handler gin.HandlerFunc, tstatusCode int, tbody string, t *testing.T) {
	CheckHandlerResp(handler, http.MethodGet, tstatusCode, tbody, t)
}

// CheckHandler tests reponse body and status code of a generic gin handler
func CheckHandlerResp(handler gin.HandlerFunc, method string, tstatusCode int, tbody string, t *testing.T) {
	// setup the http request
	req, err := http.NewRequest(method, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	server := gin.New()
	server.Handle(method, "/", handler)

	server.ServeHTTP(rr, req)

	if rr.Code != tstatusCode {
		t.Fatalf("code: %d != %d", rr.Code, tstatusCode)
	}

	if body := rr.Body.String(); body != tbody {
		t.Fatalf("body: %s != %s", body, tbody)
	}
}
