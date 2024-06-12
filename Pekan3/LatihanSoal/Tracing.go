package main

import ("fmt")

func main() {
    var a int 

    a = 10

    procB(&a, &a, &a)
    fmt.Println(a)
}

func procB(a *int, b, c *int) {
    *b = *b + *c
	*c = *a + *b + *c
}
