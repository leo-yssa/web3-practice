package swagger

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
