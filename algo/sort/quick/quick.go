package quick

func quickSort(m []int) {
	quick(m, 0, len(m)-1)
}

// quick
func quick(m []int, start, end int) {
	if start >= end {
		return
	}
	divide := partition(m, start, end)
	quick(m, start, divide-1)
	quick(m, divide+1, end)
}

// partition
func partition(m []int, start, end int) int {
	pivot := m[end]
	i := start
	for j := start; j < end; j++ {
		if m[j] < pivot {
			m[j], m[i] = m[i], m[j]
			i += 1
		}
	}
	m[i], m[end] = m[end], m[i]
	return i
}
