package zerocross

// Both function returns places of zero crossing in data slice, both from negative and positive value.
// Slice 'f' containes places in the original slice where signal cross zero from both directions.
// Slice 'k' containes directions of zero crossing by original slice. '-1' signal
// crossed from positive to negative, and vice versa.
func Both(n []float64) ([]int, []int) {
	niz := make([]int, len(n))
	k := make([]int, len(n)-1)
	f := []int{}

	for i, v := range n {
		if v > 0 {
			niz[i] = 1
		} else {
			niz[i] = 0
		}
	}

	k1 := niz[1:]
	k2 := niz[0 : len(niz)-1]
	for i := 0; i < len(niz)-1; i++ {
		k[i] = k1[i] - k2[i]
		if k[i] != 0 {
			f = append(f, i)
		}
	}
	return f, k
}

// ToPositive function returns places of zero crossing in data slice, to positive value.
// Slice 'f' containes places in the original slice where signal cross zero to positive value.
// Slice 'k' containes directions of zero crossing by original slice. '-1' signal
// crossed from positive to negative, and vice versa.
func ToPositive(n []float64) ([]int, []int) {
	niz := make([]int, len(n))
	k := make([]int, len(n)-1)
	f := []int{}

	for i, v := range n {
		if v > 0 {
			niz[i] = 1
		} else {
			niz[i] = 0
		}
	}

	k1 := niz[1:]
	k2 := niz[0 : len(niz)-1]
	for i := 0; i < len(niz)-1; i++ {
		k[i] = k1[i] - k2[i]
		if k[i] > 0 {
			f = append(f, i)
		}
	}
	return f, k
}

// ToNegative function returns places of zero crossing in data slice, to negative value.
// Slice 'f' containes places in the original slice where signal cross zero to negative value.
// Slice 'k' containes directions of zero crossing by original slice. '-1' signal
// crossed from positive to negative, and vice versa.
func ToNegative(n []float64) ([]int, []int) {
	niz := make([]int, len(n))
	k := make([]int, len(n)-1)
	f := []int{}

	for i, v := range n {
		if v > 0 {
			niz[i] = 1
		} else {
			niz[i] = 0
		}
	}

	k1 := niz[1:]
	k2 := niz[0 : len(niz)-1]
	for i := 0; i < len(niz)-1; i++ {
		k[i] = k1[i] - k2[i]
		if k[i] < 0 {
			f = append(f, i)
		}
	}
	return f, k
}
