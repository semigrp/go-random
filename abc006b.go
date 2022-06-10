package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var dp [n + 2]int
	dp[0] = 0
	dp[1] = 0
	dp[2] = 1

	for i := 3; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	fmt.Println(dp[n-1])
}
