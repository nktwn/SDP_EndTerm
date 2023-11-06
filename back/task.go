package back

type Task struct {
	ID   int
	Name string
	Done bool
}

type TaskFactory struct {
	nextID int
}

func (f *TaskFactory) CreateTask(name string) *Task {
	f.nextID++
	return &Task{ID: f.nextID, Name: name, Done: false}
}

