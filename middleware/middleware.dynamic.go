package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	Dynamic struct {
		Host string
		TLS  string
	}
)

func NewDynamic() *Dynamic {
	return &Dynamic{}
}
func (this_ *Dynamic) Use() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		this_.Host = ctx.Request.Host
		fmt.Println(ctx.Request.Host)
		fmt.Println(ctx.Request.Header.Get("authority"))
		if ctx.Request.TLS != nil {
			this_.TLS = "https"
		} else {
			this_.TLS = "http"
		}
	})
}
