package routers

import (
	"myblog/routers/health"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	dashboard_services "myblog/services/dashboard"
	home_services "myblog/services/home"
	middleware "myblog/middlewares"
	timeline_services "myblog/services/timeline"
)

const (
	PATH_PREFIX string = "myblog"
)

func InitializeRouters() error {
	r := gin.Default()

	//static file
	r.Static("/assets", "./assets")
	r.Use(cors.AllowAll())
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/", home_services.HomeService)

	r.GET("/health", health.CheckHealth)

	r.GET("/dashboard", dashboard_services.DashboardService)

	r.GET("/timeline", timeline_services.TimelineService)

	r.Use(middleware.ErrorHandler)

	return r.Run(":8080")
}
