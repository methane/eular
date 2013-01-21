package main

import (
    "math/big"
    "fmt"
)

func fibs() (c chan *big.Int) {
    a := big.NewInt(1)
    b := big.NewInt(1)
    c = make(chan *big.Int)

    go func() {
        for {
            c <- a
            var x big.Int
            x.Add(a, b)
            a, b = b, &x
        }
    }()
    return
}

func main() {
    c := fibs()
    i := 1
    for a := range c {
        if len(a.String()) >= 1000 {
            fmt.Println(i, a)
            break
        }
        i++
    }
}
