package main

import (
    "fmt"
)

func is_pandigital(a, b int) bool {
    var used [10]bool

    check := func(n int) bool {
        for n > 0 {
            m := n % 10
            if m == 0 || used[m] {
                return false
            }
            used[m] = true
            n /= 10
        }
        return true
    }

    return check(a) && check(b) && check(a*b)
}

func num_digits(n int) (m int) {
    for n > 0 {
        m++
        n /= 10
    }
    return
}

func main() {
    products := make(map[int]bool)
    // There are 9 digits.
    // n digits number * m digits number is at least (n + m - 1) digits number
    // and at most (n + m) digits numbers.
    // We consider about n <= m < n*m case:
    // - 1 digits * 4 digits = 4 digits
    // - 2 digits * 3 digits = 4 digits
    for a := 2; a < 10; a++ {
        for b := 1234; b <= 9876; b++ {
            if is_pandigital(a, b) {
                fmt.Println(a, b, a*b)
                products[a*b] = true
            }
        }
    }
    for a := 12; a < 99; a++ {
        if a % 10 == 0 {
            continue
        }
        for b := 123; b <= 987; b++ {
            if is_pandigital(a, b) {
                fmt.Println(a, b, a*b)
                products[a*b] = true
            }
        }
    }
    sum := 0
    for v, _ := range products {
        fmt.Println(v)
        sum += v
    }
    fmt.Println(sum)
}
