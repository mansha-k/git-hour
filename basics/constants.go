package main
import (
   "fmt"
   "math"
)

func main(){
	const abc int =4 
	const xyz = 5
	const ghi= 6
	const hi float32 = 6
	fmt.Println(abc, " " , xyz + ghi+ hi )
	// + gets added 
	// float float32 6 printed as 6 and not 6.0 
	const add  = 4+ 6.6 // printed as 10 and not 10.0
	const add2 = hi + 9
	fmt.Println(add , " ", add2 , " hehe")

	const b = "mansha" + "kaz"
	fmt.Println(b)

	//multiple initliations in one
	const ( 
		name = "mansha"
		surname = "fatima"
		college = "zhcet"
	)
	fmt.Println(name, college)

	const (
		x= iota + 2000
		y
		z
		z1
	)
	fmt.Println(z, z1 , "\v", y)

	//using enums and iota
	const (
		_ = iota 
		kb = 1<<(10 * iota)
		mb
		gb
	)
	fmt.Println(gb , mb, kb)
 
	//multiple types allowed
	const (
		qq=1
		ww= "mansha"
		ee
		rr
	) // if not initialised takes prev value
	fmt.Println(qq , ee , rr)

	const n = 500000000000
	const d = 3e20/n
	fmt.Println(d)
    fmt.Println(int64(d))
    fmt.Println(math.Sin(n))
}
