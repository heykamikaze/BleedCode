package main

import "fmt"

func fibonnaci() func() int {
    prev := 0
    last := 1
    return func() int {
        defer func() {
            prev, last = last, prev + last
        }()
        return prev
    }
}

func main() {
    f := fibonnaci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
