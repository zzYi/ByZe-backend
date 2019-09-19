package v1

import (
	"ByZe/pkg/logging"
	"ByZe/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadFile(c *gin.Context) {

	_,file,err := c.Request.FormFile("file")

	if err != nil {
		logging.Warn(err)
		c.JSON(http.StatusInternalServerError,err)
		return
	}

	if file == nil {
		c.JSON(http.StatusBadRequest,err)
		return
	}

	fileName := upload.GetFileName(file.Filename)
	fileFullpath := upload.GetFileFullPath()

	src := fileFullpath + fileName

	if !upload.CheckFileExt(fileName){
		c.JSON(http.StatusBadRequest,nil)
		return
	}

	err = upload.CheckFile(fileFullpath)
	if err != nil{
		logging.Warn(err)
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	err = c.SaveUploadedFile(file,src)
	if err != nil {
		logging.Warn(err)
		c.JSON(http.StatusInternalServerError,nil)
		return
	}

	c.JSON(http.StatusOK,gin.H{"code":http.StatusOK,"msg":fileName})
}