package main

import (
    "fmt"
)

func coin_sums(price int, coins []int) (number int) {
    if price < 0 {
        return 0
    }
    if price == 0 || len(coins) == 0 {
        return 1
    }
    coin := coins[0]
    for price >= 0 {
        number += coin_sums(price, coins[1:])
        price -= coin
    }
    return
}

func main() {
    coins := [...]int{200, 100, 50, 20, 10, 5, 2}
    fmt.Println(coin_sums(200, coins[:]));
}
