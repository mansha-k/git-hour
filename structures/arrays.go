package main

import "fmt"
	
func main(){
	
	var arr [3]int
	fmt.Println("to start with :", arr)

	arr[1]=20
	fmt.Println("next step:", arr)
	fmt.Println("length: ", len(arr))

	arr2:=[5]int{1,2,3}
	fmt.Println("new array:", arr2)

	arr3:=[...]int{1,2,3}
	fmt.Println("new array:", arr3)

	arr4:=[...]int{ 3:400 , 7:400}
	fmt.Println("using index: ", arr4)

	var d2arr = [2][3]int{{1,2,3},{4,5,6}}
	fmt.Println(d2arr)

	var twod [3][4]int
	for i:= range 3{
		for j:=range 4{
			twod[i][j]=i*j+1
			fmt.Println(i,"*",j,"+ 1 = ",twod[i][j])
		}
		fmt.Println()
	}
	fmt.Println("final array:", twod)

	fmt.Println(arr2)
	for i,v :=range arr2{
		fmt.Println("at ",i," valued ",v)
	}

	//vallaues copies on assigneent
}