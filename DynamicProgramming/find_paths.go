package DynamicPrograming

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	//出界的格子始终只有2(m+n)
	//问题转换为在 少于max move - 1移动次数下有多少种方式到达这些格子
	//再化为少于max move -2 次数下达到上一轮格子的相邻格子
	//最短为曼哈顿距离，最短大于max可以排除
}
