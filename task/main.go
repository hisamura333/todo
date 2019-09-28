package task

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Task struct {
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