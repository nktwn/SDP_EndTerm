package back

import "sync"

type Configuration struct {
	TaskFactory TaskFactory
	TaskManager TaskManager
	Observer    TaskObserver
}

var instance *Configuration
var once sync.Once

func GetConfig() *Configuration {
	taskManager := TaskManager{}
	taskFactory := TaskFactory{taskManager: &taskManager}
	observer := TaskObserver{}

	once.Do(func() {
		instance = &Configuration{
			TaskFactory: taskFactory,
			TaskManager: taskManager,
			Observer:    observer,
		}
	})
	return instance
}
