package views

import (
	"html/template"
	"net/http"

	"github.com/sodle/gojira/internal/db"
)

func ListProjects(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	projects, err := db.ListProjects()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	template.Execute(w, projects)
}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")

	template, err := template.ParseFiles("web/templates/project.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	project, err := db.GetProjectByKey(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	template.Execute(w, project)
}
