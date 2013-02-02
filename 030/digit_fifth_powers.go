package main

import (
    "fmt"
)

func fifth_power(x int) int {
    x %= 10
    return x*x*x*x*x
}

func is_target(n int) bool {
    s := fifth_power(n/100000) +
         fifth_power(n/10000) +
         fifth_power(n/1000) +
         fifth_power(n/100) +
         fifth_power(n/10) +
         fifth_power(n)
    return s == n
}

func main() {
    sum := 0
    // 9**5 * 6 == 354294
    for i := 10; i <= 354294; i++ {
        if is_target(i) {
            fmt.Println(i)
            sum += i
        }
    }
    fmt.Println(sum)
}
