package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Priority   int   `json:"priority"`
}

type Taskslice struct {
	Lists []Task
}

func (tl *Taskslice)AddTask(task Task, fileName string)  {
	tl.Lists = append(tl.Lists, task)

	sJson, err := json.MarshalIndent(tl, "", "    ")
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile(fileName, sJson, os.ModePerm)
}

func (tl *Taskslice)TodoListUnmershal() error {
	list, err := ioutil.ReadFile("data/todo_list.json")
	if err != nil {
		return err
	}

	json.Unmarshal([]byte(list), &tl)
	return nil
}