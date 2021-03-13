package canteenController

import (
	"fmt"
	"funnel/app/errors"
	"funnel/app/service/canteenService"
	"funnel/app/utils"
	"github.com/gin-gonic/gin"
)

func Flow(context *gin.Context) {
	data, err := canteenService.FetchFlow()
	if err == nil {
		utils.ContextDataResponseJson(
			context,
			utils.SuccessResponseJson(data))
	} else {
		utils.ContextDataResponseJson(
			context,
			utils.FailResponseJson(errors.RequestFailed, fmt.Sprintf("%v: %v", data, err)))
	}
}
