package back

import (
	"fmt"
)

type Tasker interface {
	String() string
}

type TaskManager struct {
	tasks []Tasker
}

func (manager *TaskManager) AddTask(task Tasker) {
	manager.tasks = append(manager.tasks, task)
	fmt.Println(task.String())
}

func (manager *TaskManager) GetTasks() string {
	res := ""
	for _, task := range manager.tasks {
		res += task.String() + "\n"
	}
	return res
}

func (manager *TaskManager) MarkDone(taskID int) {
	for _, task := range manager.tasks {
		if basicTask, ok := task.(*TaskBasic); ok {
			if basicTask.ID == taskID {
				basicTask.Done = true
				return
			}
		}
	}
}
