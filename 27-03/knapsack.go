package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Knapsack(weight []int, names []string, values []int, capacity int) int {
	n := len(weight)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for w := 0; w <= capacity; w++ {
			if weight[i-1] <= w {
				// Option 1: Don't take the item
				// Option 2: Take the item and add its value
				dp[i][w] = max(dp[i-1][w], values[i-1]+dp[i-1][w-weight[i-1]])
			} else {
				// If item is too heavy, we can't take it
				dp[i][w] = dp[i-1][w]
			}
		}
	}
	// Print DP table for debugging
	fmt.Println("DP Table:")
	for _, row := range dp {
		fmt.Println(row)
	}

	// Max value possible with given capacity
	return dp[n][capacity]

}

func main() {
	weights := []int{2, 3, 4, 5}                               // Weights of items
	names := []string{"Book", "Laptop", "Headphones", "Phone"} // Item names
	values := []int{3, 4, 5, 6}                                // Values of items
	capacity := 5                                              // Maximum weight the knapsack can carry

	maxValue := Knapsack(weights, names, values, capacity)
	fmt.Println("Maximum value that can be obtained:", maxValue)
}
