package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/henriqueassiss/todo-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	tasks := models.Get()
	templates.ExecuteTemplate(w, "Index", tasks)
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewTask", nil)
}

func InsertTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		title := r.FormValue("title")
		description := r.FormValue("description")

		models.Insert(title, description)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	http.Redirect(w, r, "/", http.StatusBadRequest)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err.Error())
	}

	models.Delete(id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
