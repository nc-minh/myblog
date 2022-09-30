package timeline_services

import (
	"html/template"
	"log"
	"net/http"

	templates "myblog/templates"

	"github.com/gin-gonic/gin"
)

func TimelineService(ctx *gin.Context) {
	ts, err := template.ParseFiles(templates.VIEWS.Timeline)

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
		return
	}

	ctx.Header("Content-Type", "text/html; charset=utf-8")

	err = ts.Execute(ctx.Writer, gin.H{
		"title": "NCM - Timeline",
	})

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
	}
}
