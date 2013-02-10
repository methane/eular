package main

import (
    "fmt"
)

var facts [10]int

func is_curious(n int) bool {
    m := n
    s := 0
    for m > 0 {
        s += facts[m % 10]
        m /= 10
    }
    return s == n
}

func main() {
    var f int = 1
    var i int
    facts[0] = 1
    for i = 1; i<10; i++ {
        f *= i
        fmt.Println(i, "! = ", f)
        facts[i] = f
    }

    s := 0
    for n := 3; n < 10000000; n++ {
        if is_curious(n) {
            fmt.Println(n)
            s += n
        }
    }
    fmt.Println(s)
}
