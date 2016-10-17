package quicksort

func Sort(in []int) []int {
	if in == nil {
		return nil
	}
	quickSort(in, 0, len(in)-1)
	return in
}

func quickSort(in []int, p, r int) {
	if p < r {
		q := partition(in, p, r)
		quickSort(in, p, q-1)
		quickSort(in, q+1, r)
	}
}

func partition(in []int, p, r int) int {
	x := in[r]
	i := p - 1
	for j := p; j < r; j++ {
		if in[j] <= x {
			i += 1
			in[i], in[j] = in[j], in[i]
		}
	}
	in[i+1], in[r] = in[r], in[i+1]
	return i + 1
}
