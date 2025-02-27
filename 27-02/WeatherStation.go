package main

import "fmt"

// Observer Interface - every display needs to implement this
type Observer interface {
	Update(temp float64, humidity float64, pressure float64)
}

// Subject Interface - Weather station must implement this
type Subject interface {
	RegisterObserver()
	RemoveObserver()
	NotifyObserver()
}

// Concrete Subject: Weather Station
type WeatherStation struct {
	observer []Observer
	temp     float64
	humidity float64
	pressure float64
}

// RegisterObserver adds observer to list
func (ws *WeatherStation) RegisterObserver(o Observer) {
	ws.observer = append(ws.observer, o)
}

// RemoveObserver from the list
func (ws *WeatherStation) RemoveObserver(o Observer) {
	for i, observer := range ws.observer {
		if observer == o {
			ws.observer = append(ws.observer[:1], ws.observer[i+1:]...)
			break
		}
	}
}

// NotifyObserver updates all observers when state changes
func (ws *WeatherStation) NotifyObserver() {
	for _, observer := range ws.observer {
		observer.Update(ws.temp, ws.humidity, ws.pressure)
	}
}

// SetMeasurements updates the weather data and notifies observers
func (ws *WeatherStation) SetMeasurements(temp, humidity, pressure float64) {
	ws.temp = temp
	ws.humidity = humidity
	ws.pressure = pressure
	ws.NotifyObserver()
}

// Concrete observer - Displays
type CurrentConditionsDisplay struct {
	temp     float64
	humidity float64
	pressure float64
}

// Update receives new weather data
func (ccd *CurrentConditionsDisplay) Update(temp, humidity, pressure float64) {
	ccd.temp = temp
	ccd.humidity = humidity
	ccd.pressure = pressure
	ccd.Display()
}

// Display prints the current conditions
func (ccd *CurrentConditionsDisplay) Display() {
	fmt.Printf("Current conditions: Temperature %.2f°C, Humidity %.2f%%, Pressure %.2f hPa\n", ccd.temp, ccd.humidity, ccd.pressure)
}

// Concrete Observer: StatisticsDisplay
type StatisticsDisplay struct {
	tempSum     float64
	numReadings int
}

// Update receives new weather data
func (sd *StatisticsDisplay) Update(temp, humidity, pressure float64) {
	sd.tempSum += temp
	sd.numReadings++
	sd.Display()
}

// Display prints the average temperature
func (sd *StatisticsDisplay) Display() {
	avgTemp := sd.tempSum / float64(sd.numReadings)
	fmt.Printf("Statistics: Average Temperature %.2f°C over %d readings\n", avgTemp, sd.numReadings)
}

func main() {
	// Create the weather station (Subject)
	weatherStation := &WeatherStation{}

	// Create displays (Observers)
	ccDisplay := &CurrentConditionsDisplay{}
	statsDisplay := &StatisticsDisplay{}

	// Register observers
	weatherStation.RegisterObserver(ccDisplay)
	weatherStation.RegisterObserver(statsDisplay)

	// Simulate new weather data
	weatherStation.SetMeasurements(25.3, 65, 1013)
	weatherStation.SetMeasurements(26.1, 70, 1012)
	weatherStation.SetMeasurements(27.8, 60, 1011)

	// Remove an observer (e.g., stats display unsubscribes)
	weatherStation.RemoveObserver(statsDisplay)

	// Another update (only CurrentConditionsDisplay gets updated now)
	weatherStation.SetMeasurements(28.5, 55, 1010)
}
