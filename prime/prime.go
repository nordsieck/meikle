package prime

// Largest prime < 2**32
const MaxPrime = 2147483647

// Get the greatest prime less than max.  If max is
// larger than MaxUint32, return the largest prime
// less than maxuint32.
func Greatest(max uint64) uint32 {
	if max >= MaxPrime {
		return MaxPrime
	}

	field := make([]bool, max)
	for i := uint64(2); i < max; i += 1 {
		if field[i] {
			continue
		}
		for j := i * 2; j < max; j += i {
			field[j] = true
		}
	}

	for i := max - 1; i >= 2; i-- {
		if !field[i] {
			return uint32(i)
		}
	}
	return 2
}
