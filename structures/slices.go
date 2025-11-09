package main

import (
	"fmt"
	"slices"
	//"slices"
)
func main(){
	
	//slices more flexi than array
	//dyanmic in nature
	//uninitialise slice means --> len=0
	//a struct with 1.ptr 2. lenght 3. capacity

	//ways to initialise
	//1. direct values
	sli:=  []int{5,6,7,8}

	//2. from arrays or slices
	array := [...]int{1,2,3,4}
	sli_from_arr := array
	sli_from_sli := sli

	fmt.Println("arr:", array ," len= ", len(array))
	fmt.Println("sli original:", sli ," len= ", len(sli), " cap=", cap(sli))
	fmt.Println("ali_from_arr: " ,sli_from_arr)
	fmt.Println("ali_from_sli: " ,sli_from_sli)
	
	//3. create a slice from an array
	arr := [...]int{1,2,3,4,5}
	slice := arr[1:4]
	fmt.Println("created from an array :", slice)

	//4. using make()
	slice2 := make([]int, 3, 5)
	fmt.Printf("slice2 from make() : type = %T, legnth = %d, capacity = %d, slice =%d", slice2,len(slice2),cap(slice2), slice2 )
	
	slice3 := make ([]int , 4)
	fmt.Printf("\nslice3 from make() : type = %T, legnth = %d, capacity = %d, slice =%d", slice3,len(slice3),cap(slice3), slice3 )
	
	slice4 := make ([]int, 0, 4) 
	fmt.Printf("\nslice3 from make() : type = %T, legnth = %d, capacity = %d, slice =%d", slice4,len(slice4),cap(slice4), slice4 )
	
	// sli2 := append(sli2,10) not allowed 
	sli2 := append(sli_from_sli,9)
	fmt.Println("\nnow: sli_from_sli: " ,sli2)

	s := make( []string , 3)
	fmt.Println("empty slice:",s,"len:",len(s),"cap:",cap(s))
	s[0] = "a"
	s[2] = "c"
	fmt.Println("empty slice:",s,"len:",len(s),"cap:",cap(s))
	
	s = append(s, "d" , "e")
	c := make([]string, len(s))
	copy(c,s) //copies only jitna capacity hai - no error
	fmt.Println("copy", c)

	l := s[:4]
	fmt.Println("s[:4] = ", l)

	l = s[2:]
	fmt.Println("s[2:] = ", l)

	///multidimentional: 
	matrix := [][]int{
		{1,2},{3,4},{5,6},
	}
	fmt.Println(matrix)

	
	if slices.Equal(s,c){
		fmt.Print("same hai")
	} else {
		fmt.Print("not same")
	}
}