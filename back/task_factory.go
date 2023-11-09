package back

import "time"

type TaskFactory struct {
	nextID      int
	taskManager *TaskManager
}

func (f *TaskFactory) CreateTask(name string) *TaskBasic {
	if name == "" || f.taskManager.IsTaskExists(name) {
		return nil
	}
	f.nextID++
	return &TaskBasic{ID: f.nextID, Name: name, Done: false}
}

func (f *TaskFactory) CreateTimedTask(name string, deadline time.Time) *TimedTask {
	if name == "" || f.taskManager.IsTaskExists(name) {
		return nil
	}
	task := f.CreateTask(name)
	if task == nil {
		return nil
	}
	return &TimedTask{
		TaskBasic: task,
		Deadline:  deadline,
	}
}
