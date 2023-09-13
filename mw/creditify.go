package mw

import "github.com/gin-gonic/gin"

func Credits(ctx *gin.Context) {
	defer ctx.Next()
	{
		ctx.Header("Author", "z3ntl3 (efdal)")
		ctx.Header("Server", "CF Reverse Proxy -> Caddy -> Go Gin")
	}
}
