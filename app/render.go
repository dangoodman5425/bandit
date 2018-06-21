package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func Render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}

func Redirect(c *gin.Context, data gin.H, redirectURL string, templateName string)	{
	c.Redirect(301, redirectURL)
	c.HTML(http.StatusOK, templateName, data)
}
