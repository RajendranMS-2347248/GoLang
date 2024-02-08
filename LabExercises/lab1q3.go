package main

import (
    "fmt"
    "time"
)

// Define a struct to represent weather measurements
type Weather struct {
    Location   string
    Time       time.Time
    Temperature float64 // in Celsius
    Humidity   float64 // in percentage
    WindSpeed  float64 // in meters per second
}

func main() {
    // Prompt the user to enter weather measurements
    var weather Weather
    fmt.Println("Enter weather measurements:")

    // Take input for location
    fmt.Print("Location: ")
    fmt.Scanln(&weather.Location)

    // Set current time
    weather.Time = time.Now()

    // Take input for temperature
    fmt.Print("Temperature (in Celsius): ")
    fmt.Scanln(&weather.Temperature)

    // Take input for humidity
    fmt.Print("Humidity (in percentage): ")
    fmt.Scanln(&weather.Humidity)

    // Take input for wind speed
    fmt.Print("Wind Speed (in meters per second): ")
    fmt.Scanln(&weather.WindSpeed)

    // Print out the weather information
    fmt.Println("\nWeather Information:")
    fmt.Println("Location:", weather.Location)
    fmt.Println("Time:", weather.Time.Format("2006-01-02 15:04:05"))
    fmt.Println("Temperature:", weather.Temperature, "°C")
    fmt.Println("Humidity:", weather.Humidity, "%")
    fmt.Println("Wind Speed:", weather.WindSpeed, "m/s")
}
