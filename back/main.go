package back

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func start() {
// 	config := GetConfig()
// 	fmt.Println("Конфигурация загружена:", config)

// 	taskFactory := TaskFactory{}
// 	taskManager := TaskManager{}
// 	observer := TaskObserver{}

// 	scanner := bufio.NewScanner(os.Stdin)
// 	for {
// 		fmt.Println("Введите команду (add <name>, done <id>, print, exit):")
// 		scanner.Scan()
// 		input := scanner.Text()
// 		args := strings.Fields(input)

// 		if len(args) == 0 {
// 			continue
// 		}

// 		switch args[0] {
// 		case "add":
// 			if len(args) < 2 {
// 				fmt.Println("Не указано имя задачи.")
// 				continue
// 			}
// 			taskName := strings.Join(args[1:], " ")
// 			task := taskFactory.CreateTask(taskName)
// 			taskManager.AddTask(task)
// 			observer.Notify(task)

// 		case "done":
// 			if len(args) < 2 {
// 				fmt.Println("Не указан ID задачи.")
// 				continue
// 			}
// 			taskID, err := strconv.Atoi(args[1])
// 			if err != nil {
// 				fmt.Println("ID задачи должен быть числом.")
// 				continue
// 			}
// 			taskManager.MarkDone(taskID)

// 		case "print":
// 			taskManager.PrintTasks()

// 		case "exit":
// 			fmt.Println("Выход из программы.")
// 			return

// 		default:
// 			fmt.Println("Неизвестная команда.")
// 		}
// 	}
// }
