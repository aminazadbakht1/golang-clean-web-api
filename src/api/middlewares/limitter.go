package middlewares

import (
	"net/http"

	"github.com/aminazadbakht1/golang-clean-web-api/api/helper"
	"github.com/didip/tollbooth/v7"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc{
	limitter := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context){
		err := tollbooth.LimitByRequest(limitter, ctx.Writer, ctx.Request)
		if err != nil{
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false,-100, err))
			return
		}else{
			ctx.Next()
		}
	}
}