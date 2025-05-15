package task

import (
	"fmt"
)

const dataFile = "tasks.json"

func AllTask() {
	fmt.Println(Tasks[0].TaskName)
	TaskLength := len(Tasks)
	for i := 1; i < TaskLength; i++ {
		stat := isFinish(i)
		fmt.Printf("%d  %s 任務狀態: %s\n", i, Tasks[i].TaskName, stat)
	}

}

func isFinish(num int) string {
	if Tasks[num].Finish {
		return "已完成"
	} else {
		return "未完成"
	}
}

func CheckTask(taskNum int) {
	if taskNum == 0 {
		fmt.Print("不存在0號任務")
		return
	}
	stat := isFinish(taskNum)
	fmt.Printf("%d  %s 任務狀態: %s", taskNum, Tasks[taskNum].TaskName, stat)
}

func Insert(task string) {
	newTask := NewTask(task)
	// temp := make([]Task, 1)
	// temp[0] = newTask
	Tasks = append(Tasks, newTask)
	AllTask()
}

func RemTask(taskNum int) []Task {
	TaskLength := len(Tasks)
	if taskNum == 0 {
		fmt.Print("任務編號從1開始")
		if TaskLength > 1 {
			return Tasks
		}
		return nil
	}
	if taskNum > TaskLength {
		fmt.Printf("超出目前任務表長度:%d\n", TaskLength)
		AllTask()
		return Tasks
	}
	for i := taskNum; i < TaskLength-1; i++ {
		Tasks[i] = Tasks[i+1]
	}
	AllTask()
	return Tasks[:TaskLength]
}

func UpdateTaskList(taskNum int, task string) []Task {
	TaskLength := len(Tasks)
	if taskNum == 0 {
		fmt.Print("任務編號從1開始")
		if TaskLength > 1 {
			return Tasks
		}
		return nil
	}
	if taskNum > TaskLength {
		fmt.Printf("超出目前任務表長度:%d", TaskLength)
		AllTask()
		return Tasks
	}
	Tasks[taskNum].TaskName = task
	AllTask()
	return Tasks
}

func DoneTask(taskNum int) bool {
	Tasks[taskNum].Finish = true
	return Tasks[taskNum].Finish
}
