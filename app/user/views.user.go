package user

import "github.com/gin-gonic/gin"

func InitializeUserRoutes(r *gin.Engine)	{
	r.GET("/u/:userID", showUserPage)
	r.POST("/u/:userID", getTestText)
}