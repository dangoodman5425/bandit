package main

import (
	"github.com/gin-gonic/gin"
	"fretwork/app/user"
	"fretwork/app/landing"
	"fretwork/managers"
	"html/template"
)

func initializeRoutes(r *gin.Engine)	{
	landing.InitializeLandingRoutes(r)
	user.InitializeUserRoutes(r)
}

func setFuncMaps(r *gin.Engine)	{
	funcMap := template.FuncMap{
		"GeneratePresignedUrl": managers.GeneratePresignedUrl,
	}
	r.SetFuncMap(funcMap)
}
