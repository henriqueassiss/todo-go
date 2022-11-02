package models

import (
	"github.com/henriqueassiss/todo-go/database"
)

type Task struct {
	Id                 int
	Title, Description string
	IsDone             bool
}

func Get() []Task {
	db := database.ConnectToDatabase()
	defer db.Close()

	taskRows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}

	var tasks []Task

	for taskRows.Next() {
		var id int
		var title, description string
		var isDone bool

		err = taskRows.Scan(&id, &title, &description, &isDone)
		if err != nil {
			panic(err.Error())
		}

		task := Task{id, title, description, isDone}
		tasks = append(tasks, task)
	}

	return tasks
}

func Insert(title, description string) {
	db := database.ConnectToDatabase()
	defer db.Close()

	insertData, err := db.Prepare("INSERT INTO tasks(title, description, isdone) values($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(title, description, false)
}

func Delete(id int) {
	db := database.ConnectToDatabase()
	defer db.Close()

	insertData, err := db.Prepare("DELETE FROM tasks WHERE id = $1")
	if err != nil {
		panic(err.Error())
	}

	insertData.Exec(id)
}
