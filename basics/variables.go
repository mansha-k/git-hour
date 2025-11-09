package main

import "fmt"

func main() {

	//var* name* type = value
	//name* := value
	 s := "hello" 

	fmt.Println("go" + "lang" + s)
	fmt.Println("1+1=", 1+1)
	fmt.Println("7.0/3.0 =", 7/3)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	var a = "mansha"
	fmt.Println(a)
	
	var b,c = "mansha", "kazmi"
	fmt.Println(b, c)

	//multi vals
	var (
		qq int =1
		ww = "mansha"
		ee = 4.5
	)
	fmt.Println(qq , ee, ww)
	
	var e int 
	fmt.Print(e, "\n")
 
	//type either be in none or just 1 - var or const place. 
	//type in both - different types- throws error
	const x = 35
	var y int= x
	fmt.Println("important case:", y)
	
}
