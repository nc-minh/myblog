package dashboard_services

import (
	"html/template"
	"log"
	"myblog/templates"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DashboardService(ctx *gin.Context) {
	ts, err := template.ParseFiles(templates.VIEWS.Dashboard)

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(ctx.Writer, nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(ctx.Writer, "Internal Server Error", 500)
	}
}
