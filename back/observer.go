package back

import "fmt"

type Observer interface {
	Notify(task *Task)
}

type TaskObserver struct {
}

func (o *TaskObserver) Notify(task *Task) {
	fmt.Printf("Задача '%s' обновлена.\n", task.Name)
}
