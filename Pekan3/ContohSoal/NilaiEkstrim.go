package main

import (
	"fmt"
)

func mencariNilaiEkstrim(n1, n2 int, min, max *int){
	if n1 > n2 {
		*min = n2
		*max = n1
	}else {
		*min = n1
		*max = n2
	}
}

func main() {
	var a, b, c, d int
	var min1, min2, max1, max2, min, max, temp int

	print("Input nilai a : ")
	fmt.Scan(&a)

	print("Input nilai b : ")
	fmt.Scan(&b)

	print("Input nilai c : ")
	fmt.Scan(&c)

	print("Input nilai d : ")
	fmt.Scan(&d)

	mencariNilaiEkstrim(a, b, &min1,&max1)
	mencariNilaiEkstrim(c, d, &min2,&max2)
	mencariNilaiEkstrim(max1, max2, &temp, &max)
	mencariNilaiEkstrim(min1, min2, &min, &temp)

	fmt.Println(max, min)

}