package home_services

import (
	"html/template"
	"log"
	"net/http"

	templates "myblog/templates"

	"github.com/gin-gonic/gin"
)

func HomeService(ctx *gin.Context) {
	ts, err := template.ParseFiles(templates.VIEWS.Home)

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(ctx.Writer, gin.H{
		"title": "NCM - Blog",
	})

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
	}
}
