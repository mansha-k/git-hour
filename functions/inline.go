package main
import "fmt"

//without a name
//assigned to variables 
//passed as arguments - normal
//returned from functions
//func(params)returntype{body}
//immediately invoke

/*normal function : func name(params) rts {body}
 call: add(2,3)
inline : variable name =func(params) rts{body}
 call add(2,3) or direct (2,3)
*/
func main(){
	
	mak := func(a,b int) int { return a+b } (5,3)
	fmt.Println(mak)

	add := func(x,y int) int { return x+y} 
	fmt.Println(add(2,3))

// 	anonymous function
}