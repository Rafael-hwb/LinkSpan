package bootstrap

import (
	"linkspan/project/internal/router"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	router.Register(r)
}
