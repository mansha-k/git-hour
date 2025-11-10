package main
import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string
    Age  int
}

func main() {
    u := User{Name: "Alice", Age: 23}
    data, _ := json.Marshal(u)
    fmt.Println(string(data))
}

/*Output:

{"Name":"Alice","Age":23}


Insights:

Struct fields must be capitalized to be exported for JSON.

You can add struct tags:

type User struct {
    Name string `json:"username"`
    Age  int    `json:"age"`
}


→ Output: {"username":"Alice","age":23}
*/