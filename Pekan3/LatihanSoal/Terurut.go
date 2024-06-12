package main

import "fmt"

func main() {
  var x, y int
  for {
    print("Masukkan Angka Pertama : ")
	fmt.Scan(&x)

	print("Masukkan Angka Kedua : ")
	
	fmt.Scan(&y)
    if x == y {
      break
    }
    mengurutkan(&x, &y)
    fmt.Println(x, y)
  }
}

func mengurutkan(x, y *int) {
  if *x > *y {
    temp := *x
    *x = *y
    *y = temp
  }
}
