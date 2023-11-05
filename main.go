package main

import (
	"fmt"
	"sync"
)

// --- Singleton Pattern ---
type Config struct {
	AppName string
}

var instance *Config
var once sync.Once

func GetConfigInstance() *Config {
	once.Do(func() {
		instance = &Config{AppName: "Task Management System"}
	})
	return instance
}

// --- Factory Pattern ---
type Task interface {
	Describe() string
}

type PersonalTask struct {
	Name string
}

func (p *PersonalTask) Describe() string {
	return "Personal: " + p.Name
}

type WorkTask struct {
	Name string
}

func (w *WorkTask) Describe() string {
	return "Work: " + w.Name
}

func NewTask(taskType, name string) Task {
	switch taskType {
	case "personal":
		return &PersonalTask{Name: name}
	case "work":
		return &WorkTask{Name: name}
	default:
		return nil
	}
}

// --- Command Pattern ---
type Command interface {
	Execute()
}

type AddTaskCommand struct {
	TaskType, TaskName string
	TaskList           *[]Task
}

func (a *AddTaskCommand) Execute() {
	task := NewTask(a.TaskType, a.TaskName)
	*a.TaskList = append(*a.TaskList, task)
	fmt.Println("Added: " + task.Describe())
}

// --- Observer Pattern ---
type Observer interface {
	Update(task Task)
}

type TaskObserver struct {
	Name string
}

func (to *TaskObserver) Update(task Task) {
	fmt.Printf("Observer %s: New task added - %s\n", to.Name, task.Describe())
}

// --- Subject ---
type TaskManager struct {
	observers []Observer
	tasks     []Task
}

func (tm *TaskManager) AddObserver(observer Observer) {
	tm.observers = append(tm.observers, observer)
}

func (tm *TaskManager) NotifyObservers(task Task) {
	for _, observer := range tm.observers {
		observer.Update(task)
	}
}

func (tm *TaskManager) AddTask(taskType, taskName string) {
	cmd := &AddTaskCommand{TaskType: taskType, TaskName: taskName, TaskList: &tm.tasks}
	cmd.Execute()
	tm.NotifyObservers(NewTask(taskType, taskName))
}

func main() {
	// Singleton
	config := GetConfigInstance()
	fmt.Println("App: " + config.AppName)

	// Observer
	taskManager := &TaskManager{}
	observer1 := &TaskObserver{Name: "Observer1"}
	observer2 := &TaskObserver{Name: "Observer2"}

	taskManager.AddObserver(observer1)
	taskManager.AddObserver(observer2)

	// Command and Factory
	taskManager.AddTask("personal", "Buy groceries")
	taskManager.AddTask("work", "Finish project")
}
