package sorting

func merge(p, q []int) []int {
	sorted := make([]int, len(p)+len(q))
	pi, qi, i := 0, 0, 0
	for pi < len(p) && qi < len(q) {
		if p[pi] <= q[qi] {
			sorted[i] = p[pi]
			pi++
		} else {
			sorted[i] = q[qi]
			qi++
		}
		i++
	}

	if qi < len(q) {
		copy(sorted[i:], q[qi:])
	}

	if pi < len(p) {
		copy(sorted[i:], p[pi:])
	}

	return sorted
}

func MergeSort(instance []int) []int {
	l, r := 0, len(instance)
	mid := (l + r) / 2

	if len(instance) == 1 {
		return instance
	}

	left := MergeSort(instance[l:mid])
	right := MergeSort(instance[mid:r])

	return merge(left, right)
}
