package main

import (
	"fmt"
	"maps"
)
func main(){

	//1 allocate memory
	m1:= make(map[string]int)

	//copy
	m2 := m1

	//direct values
	m := map[string]int{"cats" : 2, "dogs" :4}
	fmt.Println(m)

	//oprations
	m["ants"] = 5
	m["cats"] = m["cats"]+1

	fmt.Println(m)
	fmt.Println(m["snake"]) // index not present --> gives default of value of that type intead of error

	val := m["dogs"]
	fmt.Println("dogs=",val)
	fmt.Println("number of entries in the map =", len(m))

	delete(m,"cats")
	fmt.Println(m)

	if maps.Equal(m,m2){
		fmt.Println("equal hai")
	} else {
		fmt.Println("nope")
	}

	//important use of flags
	m2["t1"] = 4
	m2["t3"] = 6
	m2["t2"] = 2

    score, flag := m2["t3"]
	if flag {
		fmt.Print("t4 exists in race with score ", score)
	} else {
		fmt.Print("not found")
	}
}