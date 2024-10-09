package routers

import (
	"GinniBackend/handlers"

	"github.com/gin-gonic/gin"
)

func TranslateRouters(r *gin.Engine) {
	r.POST("/translate", handlers.TranslateHandler)
}
