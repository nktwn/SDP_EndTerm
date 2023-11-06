package back

import "sync"

type Configuration struct {
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
