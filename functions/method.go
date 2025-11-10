package main
import "fmt"

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    c := Circle{Radius: 5}
    fmt.Println("Area:", c.Area())
}
Output:

makefile
Copy code
Area: 78.5
Insights:

Methods are functions with a receiver: (c Circle).

Receiver can be value or pointer → func (c *Circle) Scale(f float64).

No this keyword. Receiver name is up to you.