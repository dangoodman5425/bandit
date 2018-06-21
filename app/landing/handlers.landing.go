package landing

import "github.com/gin-gonic/gin"

func InitializeLandingRoutes(r *gin.Engine)	{
	r.GET("/", showLandingPage)
}
