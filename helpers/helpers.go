package helpers

/* сортировка вставкой */
func InsertionSort(inp []rune, n int) []rune {
	j := 0
	for i := 1; i < n; i++ {
		for j = i; j > 0 && inp[j] < inp[j-1]; j-- {
			inp[j-1], inp[j] = inp[j], inp[j-1]
		}

	}
	return inp
}

// min <= i <= max
func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

// конвертация строки в INT
func Atoi(s string) int {
	var (
		n uint64
		i int
		v byte
	)
	for ; i < len(s); i++ {
		d := s[i]
		if '0' <= d && d <= '9' {
			v = d - '0'
		} else if 'a' <= d && d <= 'z' {
			v = d - 'a' + 10
		} else if 'A' <= d && d <= 'Z' {
			v = d - 'A' + 10
		} else {
			n = 0
			break
		}
		n *= uint64(10)
		n += uint64(v)
	}
	return int(n)
}

// конвертация INT во FLOAT
func Float(i int) float64 {
	return float64(i)
}

// Присваиваем INT в *INT
func IntPtr(i int) *int {
	return &i
}
