package handler

import (
	"io/ioutil"
	"net/http"

	"text/template"
	"encoding/json"
	"github.com/hisamura333/todo/task"

)

func Index(w http.ResponseWriter, r *http.Request) {
	var s task.Taskslice
	list, err := ioutil.ReadFile("data/todo_list.json")

	json.Unmarshal([]byte(list), &s)

	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, s.Lists)
	if err != nil {
		panic(err)
	}
}
