package user

import (
	"github.com/gin-gonic/gin"
	"fretwork/app"
	"strconv"
	"log"
	"fretwork/managers"
)

func showUserPage(c *gin.Context)	{
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		log.Panic(err)
	}
	app.Render(c, gin.H{
		"title": "User Page",
		"user_id": userID,
	}, "user.html")
}

func getTestText(c *gin.Context)	{
	objectPath := "images/profanity.jpg"
	managers.DownloadObject(objectPath, c.Writer)
}
