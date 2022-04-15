package views

import (
	"net/http"
	"path"
	"text/template"
)

type ViewsResource struct {
	pathToTemplates string
}

func (views *ViewsResource) renderHTML(w http.ResponseWriter, data any, mainTemplateName string, templateNames ...string) error {
	tmpl := template.New(mainTemplateName)

	for i, templateName := range templateNames {
		templateNames[i] = path.Join(views.pathToTemplates, templateName)
	}
	tmpl.ParseFiles(templateNames...)

	return tmpl.Execute(w, data)
}

func (views *ViewsResource) IndexPage(w http.ResponseWriter, r *http.Request) {
	templateFileNames := []string{"index.html"}
	views.renderHTML(w, nil, "index.html", templateFileNames...)
}

func (views *ViewsResource) LoginPage(w http.ResponseWriter, r *http.Request) {
	templateFileNames := []string{"login.html"}
	views.renderHTML(w, nil, "login.html", templateFileNames...)
}
