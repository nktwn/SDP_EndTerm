package main

import (
	"fmt"
	"sync"
)

type Observer interface {
	Update(temperature float64)
}

type WeatherStation struct {
	observers   []Observer
	temperature float64
}

func (ws *WeatherStation) SetTemperature(temp float64) {
	ws.temperature = temp
	ws.NotifyObservers()
}

func (ws *WeatherStation) NotifyObservers() {
	for _, o := range ws.observers {
		o.Update(ws.temperature)
	}
}

type Command interface {
	Execute()
}

type UpdateTemperatureCommand struct {
	weatherStation *WeatherStation
	temperature    float64
}

func (utc *UpdateTemperatureCommand) Execute() {
	utc.weatherStation.SetTemperature(utc.temperature)
}

type NotificationStrategy interface {
	Notify(temperature float64)
}

type EmailNotification struct{}

func (en *EmailNotification) Notify(temperature float64) {
	fmt.Printf("New notification: Temperature is %.2f°C\n", temperature)
}

var weatherStationInstance *WeatherStation
var once sync.Once

func GetWeatherStationInstance() *WeatherStation {
	once.Do(func() {
		weatherStationInstance = &WeatherStation{}
	})
	return weatherStationInstance
}

func main() {
	weatherStation := GetWeatherStationInstance()

	command := &UpdateTemperatureCommand{
		weatherStation: weatherStation,
		temperature:    6.0,
	}
	command.Execute()

	strategy := &EmailNotification{}
	strategy.Notify(6.0)
}
