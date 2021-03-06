package controllers

import (
	"github.com/laof/filesserver/models"
	"github.com/laof/filesserver/utils"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type filesInfo struct {
	*models.Files
	Success bool `json:"success"`
}

func GetList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	files := models.GetDirInfo()
	w.Write(utils.JsonData(filesInfo{files, true}))
}
