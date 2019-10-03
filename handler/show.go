package handler

import (
	"encoding/json"
	"github.com/hisamura333/todo/task"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
)

func Show(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	slice := strings.Split(path, "/")
	taskID, _ := strconv.Atoi(slice[2])

	var ts task.Taskslice
	ts.TodoListUnmershal()


	var selectedTask task.Task
	for _, task := range ts.Lists {
		if taskID == task.Id {
			selectedTask.Details = task.Details
			selectedTask.Id = task.Id
			selectedTask.Priority = task.Priority
			selectedTask.Name = task.Name
		}
	}
	tmpl, err := template.ParseFiles("view/show.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, selectedTask)
	if err != nil {
		panic(err)
	}

}

func CreateDetails(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	slice := strings.Split(path, "/")
	taskID, _ := strconv.Atoi(slice[2])

	detailsValue := r.FormValue("details")

	var ts task.Taskslice
	ts.TodoListUnmershal()

	var Newts task.Taskslice
	var selectedID int
	for i, task := range ts.Lists {
		if taskID == task.Id {
			task.Details = detailsValue
			Newts.Lists = append(Newts.Lists, task)
			selectedID = i
		} else {
			Newts.Lists = append(Newts.Lists, task)
		}
	}
	NewtsJson, err := json.MarshalIndent(Newts, "", "    ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("data/todo_list.json", NewtsJson, os.ModePerm)

	tmpl, err := template.ParseFiles("view/show.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, Newts.Lists[selectedID])
	if err != nil {
		panic(err)
	}
}
