package json_parser

// CompareRuneSlices n is the upperbound to compare r1 and r2. r1 and r2 are the runes to compare
func CompareRuneSlices(r1 []rune, r2 []rune, n int) bool {
	if n > len(r2) || n > len(r1) {
		return false
	}

	for i := 0; i < n; i++ {
		if r1[i] != r2[i] {
			return false
		}
	}
	return true
}
