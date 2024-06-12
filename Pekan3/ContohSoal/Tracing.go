package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	print("Masukkan Nilai a : ")
	fmt.Scan(&a)

	print("Masukkan Nilai b : ")
	fmt.Scan(&b)

	print("Masukkan Nilai c : ")
	fmt.Scan(&c)

	simple_calc(b,&c)
	fmt.Println(a,b,c)

	print("Masukkan Nilai b : ")
	fmt.Scan(&b)

	simple_calc(c,&b)
	fmt.Println(a,b,c)
}

func simple_calc(b int, c *int)int{
	var a int
	a = 10 + b - *c
	*c = *c + 4
	return a
}