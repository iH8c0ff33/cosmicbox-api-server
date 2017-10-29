package controllers

import (
	"net/http"
	"testing"

	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/models"
	"git.deutron.ml/iH8c0ff33/cosmicbox-api-server/utils"
)

func Test_getAllEvents(t *testing.T) {
	models.Initialize()
	utils.CheckGetHandlerResp(getAllEvents, http.StatusOK, "[]", t)
}
