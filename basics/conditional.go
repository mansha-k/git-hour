package main
import "fmt"
func main(){

	x:= -4
	ch:="tue"
	// no bracket in condition
	// code block in parenthesis 
	// else and } in same line - V IMP
	// switch variable can have any data type

	if x>0 {
	fmt.Println("hi mansha ")
	} else if x==0 {
		fmt.Println("hi ana")
	} else {
	fmt.Println("hi sid")
	}

	switch ch {
	case "mon": fmt.Print("office")
	case "tue" : fmt.Print("wfh")
	case "wed" : fmt.Print("leave")
	default: fmt.Print("sleep")
	}
}