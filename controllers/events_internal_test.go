package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/models"
)

func Test_getAllEvents(t *testing.T) {
	req, err := http.NewRequest("GET", "/events", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := gin.New()
	handler.GET("/events", getAllEvents)

	handler.ServeHTTP(rr, req)

	models.Initialize()

	if rr.Code != http.StatusOK {
		t.Fatalf("code: %d != %d", rr.Code, http.StatusOK)
	}

	if rr.Body.String() != "[]" {
		t.Fatalf("body: %s != %s", rr.Body.String(), "[]")
	}
}
