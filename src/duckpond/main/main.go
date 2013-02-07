package main 

import (
    "duckpond/pond"
    "fmt"
)

func main() {
    for i := 0; i < 2000000; i++ {
        c := make(chan string)
        if i % 3 == 0 {
            go pond.DuckAction(c)
            response := <-c
            fmt.Println(response)
        }
    }
}
