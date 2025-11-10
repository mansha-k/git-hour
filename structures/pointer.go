package main
import "fmt"

func main() {
    x := 10
    ptr := &x
    fmt.Println("Value:", *ptr)
}

// &x → address of x, *ptr → dereference.
// No pointer arithmetic in Go (unlike C++).
// Methods with pointer receivers allow mutating struct.