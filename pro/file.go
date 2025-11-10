package main
import (
    "fmt"
    "os"
)

func main() {
    os.WriteFile("test.txt", []byte("Hello File!"), 0644)
    data, _ := os.ReadFile("test.txt")
    fmt.Println(string(data))
}
/* Output:

Hello File!
Insights:

Files use byte slices ([]byte).
Permissions: 0644 means owner can read/write, others read-only.
Use defer file.Close() when opening with os.Open().
*/