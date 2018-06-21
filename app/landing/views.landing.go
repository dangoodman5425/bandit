package landing

import (
	"github.com/gin-gonic/gin"
	"fretwork/app"
)


func showLandingPage(c *gin.Context)	{
	app.Render(c, gin.H{
		"title": "Welcome",
	}, "landing.html")
}

