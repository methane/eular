package main

const N = 20

func main() {
	routes := make([]int64, N+1)
	for i := range routes {
		routes[i] = 1
	}
	for i := 0; i < N; i++ {
		for j := 1; j < N+1; j++ {
			routes[j] = routes[j-1] + routes[j]
		}
	}
	println(routes[N])
}
