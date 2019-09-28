package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"
	"text/template"
	"github.com/hisamura333/todo/task"

)

func Create(w http.ResponseWriter, r *http.Request)  {
	taskValue := r.FormValue("task")
	priority := r.FormValue("priority")
	priorityInt, _ := strconv.Atoi(priority)

	var s task.Taskslice
	list, err := ioutil.ReadFile("data/todo_list.json")

	if err != nil {
		panic(err)
	}

	json.Unmarshal([]byte(list), &s)

	lastTask := s.Lists[len(s.Lists) - 1]


	sNew := task.Task{Id: lastTask.Id+1,Name: taskValue, Priority: priorityInt }

    s.AddTask(sNew, "data/todo_list.json")

	tmpl, err := template.ParseFiles("view/index.html")

	err = tmpl.Execute(w, s.Lists)
	if err != nil {
		panic(err)
	}
}