package v1

import (
	"ByZe/pkg/app"
	"ByZe/pkg/e"
	"ByZe/pkg/logging"
	"ByZe/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {
	appG := app.Gin{C: c}
	_, file, err := c.Request.FormFile("file")

	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if file == nil {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
		return
	}

	fileName := upload.GetFileName(file.Filename)
	fileFullpath := upload.GetFileFullPath()

	src := fileFullpath + fileName

	if !upload.CheckFileExt(fileName) {
		appG.Response(http.StatusBadRequest, e.ErrorUploadCheckFileFormat, nil)
		return
	}

	err = upload.CheckFile(fileFullpath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadCheckFileFail, nil)
		return
	}

	err = c.SaveUploadedFile(file, src)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ErrorUploadSaveFileFail, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, fileName)
}
