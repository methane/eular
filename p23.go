package main

import (
    "fmt"
)

func is_abundant(n int) bool {
    sum_of_divisors := 0
    for i := 1; i <= n/2; i++ {
        if n % i == 0 {
            sum_of_divisors += i
        }
    }
    return sum_of_divisors > n
}

func main() {
    abundants := make([]int, 0)
    sum_of_abundants := make([]bool, 28124)

    for i := 1; i < 28123; i++ {
        if is_abundant(i) {
            abundants = append(abundants, i)
            for _, j := range abundants {
                k := i + j
                if k <= 28123 {
                    sum_of_abundants[k] = true
                }
            }
        }
    }

    sum := 0
    for i, a := range sum_of_abundants {
        if !a {
            fmt.Println(i, "is not sum of abundant")
            sum += i
        }
    }
    fmt.Println(sum)
}
