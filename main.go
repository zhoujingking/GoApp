package main

import "fmt"
import "goapp/packages/custommodule"

func main() {
    fmt.Println("Hello, World!")
		fmt.Println(custommodule.Add(1, 2))
}
