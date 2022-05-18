package helpers

import (
	"golang-test2/schemas"
	"log"

	"github.com/gin-gonic/gin"
)

func APIResponse(ctx *gin.Context, Message string, StatusCode int, Method string, Data interface{}) {

	jsonResponse := schemas.SchemaResponses{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
	}
	log.Printf("%s %d", Message, StatusCode)
	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}

func APIResponseList(
	ctx *gin.Context,
	Message string,
	StatusCode int,
	Method string,
	Data interface{},
	Count int,
	PerPage int,
	PageNo int,
) {
	if PageNo == 0 {
		PageNo = 1
	}
	if PerPage == 0 {
		PerPage = 5
	}
	lastPage := float64(Count) / float64(PerPage)
	if PerPage != 0 {
		if Count%PerPage != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = float64(Count) / float64(5)
	}

	jsonResponse := schemas.SchemaResponsesList{
		StatusCode: StatusCode,
		Method:     Method,
		Message:    Message,
		Data:       Data,
		Count:      Count,
		PerPage:    PerPage,
		PageNo:     PageNo,
		From:       ((PageNo * PerPage) - PerPage) + 1,
		To:         PageNo * PerPage,
		LasPage:    int(lastPage),
	}
	if Data == nil {
		jsonResponse.Data = []interface{}{}
	} else {
		jsonResponse.Data = Data
	}
	log.Printf("%s %d", Message, StatusCode)
	if StatusCode >= 400 {
		ctx.AbortWithStatusJSON(StatusCode, jsonResponse)
	} else {
		ctx.JSON(StatusCode, jsonResponse)
	}
}
