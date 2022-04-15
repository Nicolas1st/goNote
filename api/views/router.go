package views

import "net/http"

func NewViewsRouter(pathToTemplates string) http.Handler {
	m := http.NewServeMux()

	viewsResource := ViewsResource{
		pathToTemplates: pathToTemplates,
	}

	m.HandleFunc("/", viewsResource.IndexPage)
	m.HandleFunc("/login", viewsResource.LoginPage)

	return m
}
