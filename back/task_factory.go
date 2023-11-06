package back

import "time"

type TaskFactory struct {
	nextID int
}

func (f *TaskFactory) CreateTask(name string) *TaskBasic {
	f.nextID++
	return &TaskBasic{ID: f.nextID, Name: name, Done: false}
}

func (f *TaskFactory) CreateTimedTask(name string, deadline time.Time) *TimedTask {
	task := f.CreateTask(name)
	return &TimedTask{
		TaskBasic: task,
		Deadline:  deadline,
	}
}
