package main

import (
    "fmt"
)

func gcd(n, m int) int {
    if n < m {
        return gcd(m, n)
    }
    if n % m == 0 {
        return m
    }
    return gcd(m, n%m)
}

func equals(n1, d1, n2, d2 int) bool {
    if n2 == 0 || d2 == 0 {
        return false
    }
    g1 := gcd(n1, d1)
    g2 := gcd(n2, d2)
    return n1 / g1 == n2 / g2 && d1 / g1 == d2 / g2
}

func curious_fraction(n, d int) bool {
    if n % 10 == d / 10 && equals(n, d, n / 10, d % 10) {
        return true
    }
    if n / 10 == d % 10 && equals(n, d, n % 10, d / 10) {
        return true
    }
    return false
}

func main() {
    sn := 1
    sd := 1
    for numerator := 10; numerator < 99; numerator++ {
        for denominator := numerator+1; denominator < 100; denominator++ {
            if numerator % 10 == 0 && denominator % 10 == 0 {
                continue  // skip trivial case
            }
            if curious_fraction(numerator, denominator) {
                fmt.Println("curious: ", numerator, "/", denominator)
                sn *= numerator
                sd *= denominator
            }
        }
    }
    fmt.Println(sn, "/", sd)
    gcd_ := gcd(sn, sd)
    fmt.Println(sn / gcd_, "/", sd / gcd_)
}
