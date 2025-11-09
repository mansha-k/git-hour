package main
import "fmt"
type Person struct{
	Name string
	age int
}
func main(){
	p:=Person{Name:"Mansha",age:23}
	fmt.Println(p.Name,p.age)
}