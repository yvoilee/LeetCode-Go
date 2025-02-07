package leetcode

/*
从起点开始接下来有 n 个方块，相邻方块间的距离都为 1，每个方块上有增加体力的食用蘑菇或减少体力的毒蘑菇，蘑菇带来的体力改变是已知的，记为change[i]。一个人初始体力为 m，每次可以往前跳任意个方块，体力耗尽就会死掉。
假设每跳一次消耗的体力是跳的距离的平方。问这个人能否跳到终点，如果能，求可能剩余的最大体力。
*/

func jumpGamePlus(change []int, m int) int {
	dp := make([]int, len(change))
	dp[0] = m
	for i := 1; i < len(change); i++ {
		for j := 0; j < i; j++ {
			if dp[j] == 0 {
				continue
			}
			val := dp[j] + change[i] - (i-j)*(i-j)
			dp[i] = max(dp[i], val)
		}
	}
	return dp[len(change)-1]
}
