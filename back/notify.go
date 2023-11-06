package back

import "fmt"

type Observer interface {
	Notify(task *TaskBasic)
}

type TaskObserver struct {
}

func (o *TaskObserver) Notify(task *TaskBasic) {
	fmt.Printf("Задача '%s' обновлена.\n", task.Name)
}
