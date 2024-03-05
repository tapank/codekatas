package pascal

// Triangle returns a Pascal's triangle with n rows.
// Each row is represented by a slice of integers.
// For example, the 5th row of Pascal's triangle is [1, 4, 6, 4, 1].
// If n is less than 0, Triangle returns nil.
// If n is 0, Triangle returns an empty slice.
func Triangle(n int) [][]int {
	if n < 0 {
		return nil
	}
	triangle := make([][]int, n)
	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle
}
