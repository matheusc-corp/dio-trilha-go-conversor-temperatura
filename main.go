package main

import "fmt"

func main() {

	var tempKelvin float64 = 300
	tempCelsius := (tempKelvin - 273)

	fmt.Printf("%g°K em Celsius é %g°C", tempKelvin, tempCelsius)

}
