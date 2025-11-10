package main
import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Message from goroutine"
    }()

    msg := <-ch
    fmt.Println(msg)
}
Output:

csharp
Copy code
Message from goroutine
Insights:

Channels = communication between goroutines.

ch <- val sends, <-ch receives.

Default channels are blocking until paired.