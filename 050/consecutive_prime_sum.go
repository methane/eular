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
                break
            }
        }
    }
    return ps
}

func main() {
    MILLION := 1000000
    primes := PrimesSmallerLeN(MILLION)

    psmap := make(map[int]bool)
    for _, p := range primes {
        psmap[p] = true
    }

    max_sum := 0
    max_con := 0
    N := len(primes)
    for i := 0; i < N; i++ {
        sum := 0
        con := 0
        for j := i; j < N; j++ {
            con++
            sum += primes[j]
            if sum > MILLION {
                break
            }
            if con > max_con && psmap[sum] {
                max_con = con
                max_sum = sum
                fmt.Println(i, j, sum)
            }
        }
    }
    fmt.Println(max_con, max_sum)
}
