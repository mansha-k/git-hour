package main
import "fmt"

func change(s []int) {
    s[0] = 99 // changes underlying array
}

func main() {
   s := "hello"
for i, ch := range s {
    fmt.Println(i, ch)
} // [99 77 3]
}