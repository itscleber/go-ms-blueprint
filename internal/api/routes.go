package api

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine) {
	registerHealthRoutes(r)
}
