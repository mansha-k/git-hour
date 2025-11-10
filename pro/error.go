package main
import (
    "errors"
    "fmt"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    res, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", res)
    }
}

/*Go doesn’t have exceptions. Errors are just values (error type).

Idiomatic pattern: res, err := function().

if err != nil {} check always.
*/