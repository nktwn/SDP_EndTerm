package back

import (
	"fmt"
)

type TaskBasic struct {
	ID   int
	Name string
	Done bool
}

func (tb *TaskBasic) String() string {
	doneLabel := "Не выполнено"
	if tb.Done {
		doneLabel = "Выполнено"
	}
	return fmt.Sprintf("%v: Задача: %s, Статус: %s", tb.ID, tb.Name, doneLabel)
}
