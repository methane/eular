package main

import (
    "fmt"
)

func PrimesSmallerLeN(n int) []int {
    ps := []int {2}
    for i := 3; i <= n; i+=2 {
        for _, p := range ps {
            if i % p == 0 {
                break
            }
            if p*p > i {
                ps = append(ps, i)
                if is_pandigital(i) {
                    fmt.Println(i)
                }
                break
            }
        }
    }
    return ps
}

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
    for i := 1; i<=k; i++ {
        if !digits[i] {
            return false
        }
    }
    return true
}

func main() {
    PrimesSmallerLeN(7654321)
}
