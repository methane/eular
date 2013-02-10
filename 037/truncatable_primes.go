package main

import (
    "fmt"
)

func primes() chan int {
    c := make(chan int, 2)
    ps := []int {2}
    c <- 2
    go func() {
        for {
            for i := 3;; i+=2 {
                for _, p := range ps {
                    if i % p == 0 {
                        break
                    }
                    if p*p > i {
                        ps = append(ps, i)
                        c <- i
                        break
                    }
                }
            }
        }
    }()
    return c
}

func main() {
    ltr := make(map[int]bool)
    rtl := make(map[int]bool)

    ps := make(map[int]bool)
    pc := primes()
    n := 10
    for p := range pc {
        ps[p] = true
        if p >= n {
            break
        }
        ltr[p] = true
        rtl[p] = true
    }

    sum := 0

    for len(ltr)>0 && len(rtl)>0 {
        fmt.Println("n=", n)
        n *= 10
        for p := range pc {
            ps[p] = true
            if p >= n {
                break
            }
        }

        ltr2 := make(map[int]bool)
        rtl2 := make(map[int]bool)
        m := n/10

        for p := range ltr {
            for i := 1; i < 10; i+=1 {
                if i > 2 && i % 2 == 0 {
                    continue
                }
                x := i * m + p
                if ps[x] {
                    ltr2[x] = true
                }
            }
        }
        ltr = ltr2

        for p := range rtl {
            for i := 1; i < 10; i+=1 {
                if i > 2 && i % 2 == 0 {
                    continue
                }
                x := p*10 + i
                if ps[x] {
                    rtl2[x] = true
                    if ltr[x] {
                        fmt.Println("truncatable:", x)
                        sum += x
                        fmt.Println("sum=", sum)
                    }
                }
            }
        }
        rtl = rtl2
    }
}
