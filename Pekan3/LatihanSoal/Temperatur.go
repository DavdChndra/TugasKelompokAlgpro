package main

import (
	"fmt"
)

func temperatur(celcius float64, reamur, fahrenheit, kelvin *float64){
	*reamur = 4 * celcius / 5
  	*fahrenheit = (celcius * 9 / 5) + 32
  	*kelvin = celcius + 273.15
}

func main() {
	var celcius float64

	print("Masukkan Celcius: ")
	fmt.Scan(&celcius)

	var reamur, fahrenheit, kelvin float64

	temperatur(celcius, &reamur, &fahrenheit, &kelvin)

	fmt.Println("Suhu Reamur:", reamur, "R")
	fmt.Println("Suhu Fahrenheit:", fahrenheit, "F")
	fmt.Println("Suhu Kelvin:", kelvin, "K")
}