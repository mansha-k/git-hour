package main
import (
    "fmt"
    "time"
)

func hello() {
    fmt.Println("Hello from Goroutine")
}

func main() {
    go hello()
    time.Sleep(1 * time.Second)
    fmt.Println("Main finished")
}
Output (order may vary):

css
Copy code
Hello from Goroutine
Main finished
Insights:

go funcName() runs a function concurrently.

Lightweight threads (not OS threads).

Execution order not guaranteed.

