package main

func primes() (chan int64) {
	c := make(chan int64)
	go func() {
		ps := []int64{2}
		c <- 2
		n := int64(3)
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

func main() {
	sum := int64(0)
	ps := primes()
	for {
		p := <-ps
		if p >= 2000000 {
			break
		}
		sum += p
	}
	println(sum)
}
