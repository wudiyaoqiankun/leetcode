package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}
	dq := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dq[i] = make([]bool, n+1)
	}
	dq[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 {
				dq[i][j] = dq[i][j] || (dq[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dq[i][j] = dq[i][j] || (dq[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}
	return dq[m][n]
}
