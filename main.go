package main

import (
	"github.com/gin-gonic/gin"
)

func main()	{
	router := gin.Default()
	setFuncMaps(router)
	router.LoadHTMLGlob("html/*/*.html")
	router.Static("/static", "./static")
	initializeRoutes(router)
	router.Run(":6464")
}