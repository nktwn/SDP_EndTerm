package back

import "fmt"

type Command interface {
	Execute()
}

type TaskManager struct {
	tasks []*Task
}

func (manager *TaskManager) AddTask(task *Task) {
	manager.tasks = append(manager.tasks, task)
	fmt.Printf("Задача '%s' добавлена.\n", task.Name)
}

func (manager *TaskManager) GetTasks() string {
	res := ""
	for _, task := range manager.tasks {
		doneLabel := "Не выполнено"
		if task.Done {
			doneLabel = "Выполнено"
		}
		res += fmt.Sprintf("%v: Задача: %s, Статус: %s\n", task.ID, task.Name, doneLabel)
	}
	return res
}

func (manager *TaskManager) MarkDone(taskID int) {
	for _, task := range manager.tasks {
		if task.ID == taskID {
			task.Done = true
			fmt.Printf("Задача '%s' отмечена как выполненная.\n", task.Name)
			return
		}
	}
	fmt.Println("Задача с указанным ID не найдена.")
}
