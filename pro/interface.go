package main
import "fmt"

type Shape interface {
    Area() float64
}

type Square struct{ Side float64 }

func (s Square) Area() float64 { return s.Side * s.Side }

func main() {
    var s Shape = Square{Side: 4}
    fmt.Println("Square Area:", s.Area())
}
Output:

mathematica
Copy code
Square Area: 16
Insights:

Interfaces = contract of methods.

Any type implementing all methods → implicitly satisfies interface (no implements keyword like Java).

interface{} = empty interface (like Object in Java).