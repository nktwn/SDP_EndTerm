package cmd

import (
	back "SDP/back"
	"fmt"
	"html/template"
	"net/http"
)

var front *template.Template

var config *back.Configuration

type pageData struct {
	Tasks     string
	ShowError int
	ErrorCall string
}

func web_page(result http.ResponseWriter, call *http.Request) {
	data := pageData{}
	taskManager := config.TaskManager
	taskFactory := config.TaskFactory
	observer := config.Observer

	if call.URL.Path != "/" {
		data.ShowError = 404
		data.ErrorCall = "Page Not Found"
		errorCall(result, call, &data)
		return
	}
	if call.Method == "GET" {
		front.ExecuteTemplate(result, "index.html", data)
	} else if call.Method == "POST" {
		if call.FormValue("process") == "add" {
			nameOfTask := call.FormValue("input")

			task := taskFactory.CreateTask(nameOfTask)
			taskManager.AddTask(task)
			observer.Notify(task)

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
	config = back.GetConfig()

	front = template.Must(template.ParseGlob("html/*.html"))

	http.HandleFunc("/", web_page)
	fmt.Println("Server is listening to port #8080 ... ")

	http.ListenAndServe(":8080", nil)
}
