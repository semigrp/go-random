package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(dp(n-1) % 10007)
}

func dp(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return dp(n-1) + dp(n-2) + dp(n-3)
}
