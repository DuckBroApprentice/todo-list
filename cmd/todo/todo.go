package main

import (
	"flag"
	"tasklist/pkg/task"
)

func main() {

	addTask := flag.String("a", "", "Add a new task")
	remTask := flag.Int("r", 0, "Remove a task")
	doneTask := flag.Int("d", 0, "Change task stat")
	flag.Parse()
	newTask := *addTask
	remove := *remTask
	finishTask := *doneTask

	switch {
	case newTask != "":
		task.Insert(newTask)
	case remove != 0:
		task.RemTask(remove)
	case finishTask != 0:
		task.DoneTask(finishTask)
	}

}
