package templates

type HtmlName struct {
	Home      string
	Dashboard string
	Timeline string
}

var TemplateEndpoints = "./templates"

var VIEWS = &HtmlName{
	Home:      TemplateEndpoints + "/home/index.html",
	Dashboard: TemplateEndpoints + "/dashboard/dashboard.html",
	Timeline: TemplateEndpoints + "/timeline/timeline.html",
}
