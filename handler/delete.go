package handler

import (
	"github.com/hisamura333/todo/task"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"encoding/json"

	"text/template"
)

func Delete(w http.ResponseWriter, r *http.Request)  {
	path := r.URL.Path
	slice := strings.Split(path, "/")
	taskID, _ := strconv.Atoi(slice[2])


	var ts task.Taskslice
    ts.TodoListUnmershal()

	var Deletedts task.Taskslice
	for _, value  := range ts.Lists {
		if value.Id != taskID {
			Deletedts.Lists = append(Deletedts.Lists, value)
		}
	}
	NewtsJson, err := json.MarshalIndent(Deletedts, "", "    ")
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("data/todo_list.json", NewtsJson, os.ModePerm)


	var Newts task.Taskslice
	Newts.TodoListUnmershal()

	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(w, Newts.Lists)
}