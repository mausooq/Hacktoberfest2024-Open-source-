package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
    return (c * 9/5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
    return (f - 32) * 5/9
}

func main() {
    c := 25.0
    f := 77.0

    fmt.Printf("%.2f°C is %.2f°F\n", c, celsiusToFahrenheit(c))
    fmt.Printf("%.2f°F is %.2f°C\n", f, fahrenheitToCelsius(f))
}
