package main

import "fmt"

// Func for finding max value
func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to compute the Longest Common Subsequence (LCS)
func longestCommonSubsequence(word1, word2 string) string {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// fill the dp table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = maximum(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	// backtracking to find LCS string
	lcs := ""
	i, j := m, n
	for i > 0 && j > 0 {
		if word1[i-1] == word2[j-1] {
			lcs = string(word1[i-1]) + lcs
			// move diagonal (to the left and then to the top)
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	return lcs
}

func main() {
	word1 := "blue"
	word2 := "clues"

	resutl := longestCommonSubsequence(word1, word2)
	fmt.Println(resutl)
}
