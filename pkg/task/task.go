package task

type Task struct {
	TaskName string
	Finish   bool
}

var Tasks []Task = []Task{{TaskName: "Task List: ", Finish: false}}

func NewTask(taskName string) Task {
	newTask := Task{TaskName: taskName, Finish: false}
	return newTask
}
