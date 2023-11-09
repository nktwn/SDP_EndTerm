package cmd

import (
	back "SDP/back"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

var front *template.Template

var taskManager back.TaskManager
var taskFactory back.TaskFactory
var observer back.Observer

type pageData struct {
	Tasks     string
	ShowError int
	ErrorCall string
}

func web_page(result http.ResponseWriter, call *http.Request) {
	data := pageData{}

	if call.URL.Path != "/" {
		data.ShowError = 404
		data.ErrorCall = "Page Not Found"
		errorCall(result, call, &data)
		return
	}
	if call.Method == "GET" {
		front.ExecuteTemplate(result, "index.html", data)
	} else if call.Method == "POST" {
		switch call.FormValue("process") {
		case "add":
			nameOfTask := call.FormValue("input")

			task := taskFactory.CreateTask(nameOfTask)
			if task == nil {
				data.ShowError = 400
				data.ErrorCall = "Невозможно создать задачу: пустое имя или задача уже существует."
				errorCall(result, call, &data)
			} else {
				taskManager.AddTask(task)
				observer.Notify(task)

				data.Tasks = taskManager.GetTasks()
				front.ExecuteTemplate(result, "index.html", data)
			}
		case "add with time":
			nameOfTask := call.FormValue("input")

			task := taskFactory.CreateTimedTask(nameOfTask, time.Now().Add(48*time.Hour))
			taskManager.AddTask(task)
			observer.Notify(task.TaskBasic)

			data.Tasks = taskManager.GetTasks()
			front.ExecuteTemplate(result, "index.html", data)
		case "done":
			taskId, err := strconv.Atoi(call.FormValue("input"))

			if err != nil {
				data.ShowError = 400
				data.ErrorCall = "Bad request"
				errorCall(result, call, &data)
			}

			taskManager.MarkDone(taskId)
			data.Tasks = taskManager.GetTasks()
			front.ExecuteTemplate(result, "index.html", data)
		}

	} else {
		data.ShowError = 405
		data.ErrorCall = "Metod Not Allowed"
		errorCall(result, call, &data)
		return
	}
}

func errorCall(result http.ResponseWriter, call *http.Request, data *pageData) {
	result.WriteHeader(data.ShowError)
	front.ExecuteTemplate(result, "error.html", data)
}

func Start_page() {
	config := back.GetConfig()
	taskManager = config.TaskManager
	taskFactory = config.TaskFactory
	observer = &config.Observer

	front = template.Must(template.ParseGlob("html/*.html"))

	http.HandleFunc("/", web_page)
	fmt.Println("Server is listening to port #8080 ... ")

	http.ListenAndServe(":8080", nil)
}
