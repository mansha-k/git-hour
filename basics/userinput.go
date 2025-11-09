package main
import "fmt"
func main(){
	var name string
	fmt.Println("enter name")
	fmt.Scanln(&name)

	var contact int32
	fmt.Println("enter phone numner")
	fmt.Scanln(&contact)

	fmt.Println("information is" + name+ " and "+ fmt.Sprint(contact))
	fmt.Printf("information is %s and %d\n",name,contact)

}