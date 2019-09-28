package handler

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"text/template"
	"encoding/json"

)

type List struct {
	Name string `json:"name"`
	Priority   int   `json:"priority"`
}

type Listslice struct {
	Lists []List
}

func Index(w http.ResponseWriter, r *http.Request) {
	var s Listslice
	list, err := ioutil.ReadFile("data/todo_list.json")

	json.Unmarshal([]byte(list), &s)


	sNew := List{Name: "いいい", Priority: 4 }
	s.Lists = append(s.Lists, sNew)


	sJson, err := json.MarshalIndent(s, "", "    ")
	ioutil.WriteFile("data/todo_list.json", sJson, os.ModePerm)


	tmpl, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, s.Lists)
	if err != nil {
		panic(err)
	}
}
