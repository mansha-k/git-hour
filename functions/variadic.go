package main

import "fmt"

//variadic can take 0 or more arguments as per requirement
//func name(name ...type)
//similar to taking input in a slice --> new slice each time
//retun format sum(nums...)

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	fmt.Println("total : ", total)
	return total
}

//hello -> prefix
//alice , bob , mansha --> temporary slice 
//one by one we pick value
func greet(prefix string, names ...string){
	for _, value := range names {
		fmt.Println(prefix , value)
	}
}
func main() {

	fmt.Println("first way: ")
	fmt.Println(sum(1, 2, 3)) // 6
	fmt.Println(sum())        // 0

	fmt.Println("second way: ")
	sum(1, 2, 3) //1+2+3
	sum(4, 5)

	nums := []int{6, 7, 8, 9}
	sum(nums...)

	greet("Hello", "Alice", "bob", "mansha")
}

/*first way:
total :  6
6
total :  0
0
second way:
total :  6
total :  9
total :  30
*/
