package main

import "fmt"

//给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列。
//
// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。 
//例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
// 
//
// 若这两个字符串没有公共子序列，则返回 0。 
//
// 
//
// 示例 1: 
//
// 输入：text1 = "abcde", text2 = "ace" 
//输出：3  
//解释：最长公共子序列是 "ace"，它的长度为 3。
// 
//
// 示例 2: 
//
// 输入：text1 = "abc", text2 = "abc"
//输出：3
//解释：最长公共子序列是 "abc"，它的长度为 3。
// 
//
// 示例 3: 
//
// 输入：text1 = "abc", text2 = "def"
//输出：0
//解释：两个字符串没有公共子序列，返回 0。
// 
//
// 
//
// 提示: 
//
// 
// 1 <= text1.length <= 1000 
// 1 <= text2.length <= 1000 
// 输入的字符串只含有小写英文字符。 
// 
// Related Topics 动态规划

//leetcode submit region begin(Prohibit modification and deletion)
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] = text1[0..i-1]和text2[0..j-1]的最长公共子序列

	m, n := len(text1), len(text2)
	dp := *Make2DArray(m+1, n+1)
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func Make2DArray(m int, n int) *[][]int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	return &dp
}

//leetcode submit region end(Prohibit modification and deletion)

func main() {
	fmt.Printf("%v", longestCommonSubsequence("abcde", "ace"))
}
