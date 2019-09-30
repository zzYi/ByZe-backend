package v1

import (
	"ByZe/pkg/app"
	"ByZe/pkg/e"
	"ByZe/pkg/gredis"
	"ByZe/pkg/logging"
	"ByZe/service"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFileData(c *gin.Context) {
	appG := app.Gin{C: c}
	fileid := c.Param("fileid")
	if fileid == "" {
		appG.Response(http.StatusBadRequest, e.InvalidParams, nil)
	}

	if err := service.ProcessFileDataService(fileid); err != nil {
		logging.Error(err)
		return
	}

	rawdata, err := gredis.Get(fileid)
	if err != nil {
		logging.Error(err)
		return
	}

	var data map[string]interface{}
	if err = json.Unmarshal(rawdata, &data); err != nil {
		logging.Error(err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, data)

}

func GetDistributed(c *gin.Context) {

}
