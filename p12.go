package main

import (
	"fmt"
)

func primes() (chan int) {
	c := make(chan int)
	go func() {
		ps := []int{2}
		c <- 2
		n := int(3)
		for {
			for _, val := range ps {
				if val * val > n {
					ps = append(ps, n)
					c <- n
					break
				}
				if n % val == 0 {
					break
				}
			}
			n += 2
		}
	}()
	return c
}

func factorize(n int) (ps, counts []int) {
	for p := range primes() {
		count := 0
		for n % p == 0 {
			count++
			n /= p
		}
		if count > 0 {
			ps = append(ps, p)
			counts = append(counts, count)
		}
		if n == 1 {
			break
		}
	}
	return
}

func countDivisors(n int) (ret int) {
	ret = 1
	ps, counts := factorize(n)
	for _, c := range counts {
		ret *= (c + 1)
		_ = ps
		//fmt.Printf("%v: %v -> %v %v\n", ps, counts, c, ret)
	}
	return
}

func main() {
	tri := 1
	sum := 1
	for i := 1;; i++ {
		//fmt.Printf("%v: %v\n", i, countDivisors(i))
		c := countDivisors(sum)
		if c > 500 {
			fmt.Printf("%v: %v\n", sum, c)
			break
		}
		tri ++
		sum += tri
	}
}
