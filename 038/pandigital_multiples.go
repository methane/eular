package main

import (
    "fmt"
)

func is_pandigital(n int) bool {
    digits := make([]bool, 10)
    k := 0
    for n > 0 {
        m := n % 10
        n /= 10
        if m == 0 || digits[m] {
            return false
        }
        digits[m] = true
        k++
    }
    return k == 9
}

func main() {
    max_pandigital := 0
    for n := 9876; n > 1; n-- {
        m := n
        for i := 2; i < 10; i++ {
            t := n*i
            var k int
            for k = 10; k <= t; k *= 10 {}
            m *= k
            m += t
            if m > 987654321 {
                break
            }
            if m < 123456789 {
                continue
            }
            if is_pandigital(m) {
                if m > max_pandigital {
                    max_pandigital = m
                }
                fmt.Println(n, m)
            }
        }
    }
    fmt.Println("max:", max_pandigital)
}
