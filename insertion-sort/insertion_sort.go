package insertion

func Sort(in []int) []int {
	if in == nil {
		return nil
	}
	for i := 1; i < len(in); i++ {
		current := in[i]
		j := i - 1
		for j >= 0 && in[j] > current {
			in[j+1] = in[j]
			j = j - 1
		}
		in[j+1] = current
	}
	return in
}
