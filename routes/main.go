package routes

import (
	"net/http"

	"github.com/henriqueassiss/todo-go/controllers"
)

func Routes() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/newTask", controllers.NewTask)
	http.HandleFunc("/insertTask", controllers.InsertTask)
	http.HandleFunc("/deleteTask", controllers.DeleteTask)
}
