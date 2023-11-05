package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// ==================================
// singleton

type Configuration struct {
	// Добавьте любые глобальные настройки сюда.
}

var instance *Configuration
var once sync.Once

func GetConfig() *Configuration {
	once.Do(func() {
		instance = &Configuration{
			// инициализация конфигурации.
		}
	})
	return instance
}

// ==================================
// Factory
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

// ==================================
// Command

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

func (manager *TaskManager) PrintTasks() {
	for _, task := range manager.tasks {
		doneLabel := "Не выполнено"
		if task.Done {
			doneLabel = "Выполнено"
		}
		fmt.Printf("Задача: %s, Статус: %s\n", task.Name, doneLabel)
	}
}

// ==================================
// Observer

type Observer interface {
	Notify(task *Task)
}

type TaskObserver struct {
	// Добавьте свои поля, если необходимо.
}

func (o *TaskObserver) Notify(task *Task) {
	fmt.Printf("Задача '%s' обновлена.\n", task.Name)
}

// ==================================
// MAIN

func main() {
	config := GetConfig()
	fmt.Println("Конфигурация загружена:", config)

	taskFactory := TaskFactory{}
	taskManager := TaskManager{}
	observer := TaskObserver{}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите команду (add <name>, done <id>, print, exit):")
		scanner.Scan()
		input := scanner.Text()
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "add":
			if len(args) < 2 {
				fmt.Println("Не указано имя задачи.")
				continue
			}
			taskName := strings.Join(args[1:], " ")
			task := taskFactory.CreateTask(taskName)
			taskManager.AddTask(task)
			observer.Notify(task)

		case "done":
			fmt.Println("Функция 'done' еще не реализована.")

		case "print":
			taskManager.PrintTasks()

		case "exit":
			fmt.Println("Выход из программы.")
			return

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
