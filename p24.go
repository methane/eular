package main

import (
    "fmt"
)

func fact(n int) (ret int64) {
    ret = 1
    for ; n > 0 ; n-- {
        ret *= int64(n)
    }
    return
}


func main() {
    var t int64 = 1000000-1
    digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

    take := func (k, n int) {
        p := k + n
        t := digits[p]
        for ; p > k ; p-- {
            digits[p] = digits[p-1]
        }
        digits[k] = t
    }

    k := 9
    for ; fact(10-k) <= t ; k-- {}

    for ; t > 0 ; k++ {
        f := fact(9-k)
        n := int(t / f)
        take(k, n)
        fmt.Println(t, k, f, n)
        fmt.Println(digits)
        t %= f
    }
}
