package longestcommonsubsequence

func LCS(in1, in2 []string) (int, string) {
	if in1 == nil || in2 == nil {
		return 0, ""
	}
	m := len(in1)
	n := len(in2)
	if m == 0 || n == 0 {
		return 0, ""
	}
	c := make([][]int, m+1)
	for i := range c {
		c[i] = make([]int, n+1)
		c[i][0] = 0
	}
	for i := 0; i < n+1; i++ {
		c[0][i] = 0
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if in1[i-1] == in2[j-1] {
				c[i][j] = c[i-1][j-1] + 1
			} else {
				c[i][j] = Max(c[i][j-1], c[i-1][j])
			}
		}
	}
	return c[m][n], backtrack(c, in1, in2, m, n)
}

func backtrack(c [][]int, in1, in2 []string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	} else if in1[i-1] == in2[j-1] {
		return backtrack(c, in1, in2, i-1, j-1) + in1[i-1]
	} else {
		if c[i][j-1] > c[i-1][j] {
			return backtrack(c, in1, in2, i, j-1)

		} else {
			return backtrack(c, in1, in2, i-1, j)
		}
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
