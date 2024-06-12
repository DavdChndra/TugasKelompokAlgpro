package main

import ("fmt")

func main() {
	var N int
	var suhu float64
	
	print("Berapakali perulangan suhu : ")
	fmt.Scan(&N)

	for i:=1; i<=N; i++ {
		fmt.Println(i)
		print("Masukkan Celcius : ")
		fmt.Scan(&suhu)

		f := Fahrenheit(suhu)
		fmt.Println(suhu, "Celcius = ", f, "Fahrenheit")
	}
}

func Fahrenheit(suhu float64) float64{
	return (9 * suhu / 5) + 32
}