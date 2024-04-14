package spiralmatrix

// matrix indices increment direction
var dir = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func SpiralMatrix(size int) [][]int {
	// validate input
	if size < 0 {
		return nil
	}

	// create matrix
	m := make([][]int, size)
	if size == 0 {
		return m
	}
	for i := range m {
		m[i] = make([]int, size)
	}

	// fill matrix
	dirIndex := 0
	n := 1
	var r, c int
	for {
		// set value
		m[r][c] = n
		n++

		// evaluate next indices in the same direction
		rn, cn := r+dir[dirIndex][0], c+dir[dirIndex][1]
		if rn >= size || rn < 0 || cn >= size || cn < 0 || m[rn][cn] != 0 {
			dirIndex = (dirIndex + 1) % 4
			r, c = r+dir[dirIndex][0], c+dir[dirIndex][1]
		} else {
			r, c = rn, cn
		}

		// if we fail even after changing direction, then exit
		if r >= size || r < 0 || c >= size || c < 0 || m[r][c] != 0 {
			break
		}
	}
	return m
}
