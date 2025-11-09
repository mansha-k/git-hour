package main

import "fmt"

//func fname params return type
func plus(a int, b int) int {
    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func divide(a, b int) (int, int) {
    return a / b, a % b
}

func main() {

    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)

	q, r  := divide(7,2)
	fmt.Println("7/2 =", q, "and" , r)

	_, c := divide(7,2)
	fmt.Println("7/2 =", c)

	//dynamic
	func sum(nums ...int) {
	}
}