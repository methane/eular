package main

import (
    "fmt"
)

func primes(n int) ([]int) {
    ps := []int{2, 3}
    for i := 5; i < n; i += 2 {
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

func contains_bisect(n int, ps []int) bool {
    L := len(ps)
    //fmt.Println(n, L)
    if L == 0 {
        return false
    }
    if L == 1 {
        return ps[0] == n
    }
    m := L / 2
    //fmt.Println("m = ", m, "ps[m]=", ps[m])
    if ps[m] == n {
        return true
    }
    if n < ps[m] {
        return contains_bisect(n, ps[:m])
    }
    return contains_bisect(n, ps[m+1:])
}

func circular(n int, ps []int) bool {
    k := 0
    for i := n; i > 0; i /= 10 {
        k++
    }
    j := 1
    for i := 1; i < k; i++ {
        j *= 10
    }
    m := n
    for i := 0; i < k; i++ {
        //fmt.Println(k, m)
        if !contains_bisect(m, ps) {
            //fmt.Println("not in primes")
            return false
        }
        m = m / 10 + (m % 10) * j
    }
    return true
}

func main() {
    ps := primes(1000000)
    cnt := 0
    for _, i := range ps {
        /*
        if i > 100 {
            break
        }
        */
        if circular(i, ps) {
            cnt++
            fmt.Println(i)
        }
    }
    fmt.Println(cnt)
}
