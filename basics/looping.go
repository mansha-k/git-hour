package main
import "fmt"
func main(){
	
	// for loop
	for i:=1 ; i<=3 ; i++ {
		fmt.Println(" hi ",i)
	}

	// for as while
	
	fmt.Println("mak")

	nums := []int{20,30,40}
	for idx, val := range nums {
		fmt.Println(idx , val)
	}

	
	name := "mansha"
	for _, ch := range name {
		fmt.Println(string(ch))
	}
}