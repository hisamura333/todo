package handler

import (
	"net/http"

	"github.com/hisamura333/todo/task"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var ts task.Taskslice
	ts.TodoListUnmershal()

	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, ts.Lists)
	if err != nil {
		panic(err)
	}
}
