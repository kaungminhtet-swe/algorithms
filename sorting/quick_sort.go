package sorting

func findPivot(instances []int, start, end int) int {
	pivot := instances[end]

	for i := start; i < end; i++ {
		if instances[i] < pivot {
			instances[start], instances[i] = instances[i], instances[start]
			start++
		}
	}

	instances[start], instances[end] = instances[end], instances[start]
	return start
}

func quickSort(instances []int, start, end int) {
	if start < end {
		pi := findPivot(instances, start, end)
		quickSort(instances, start, pi-1)
		quickSort(instances, pi+1, end)
	}
}
