package codewars

// FindOdd Given an array of integers, find the one that appears an odd number of times.
// There will always be only one integer that appears an odd number of times.
func FindOdd(seq []int) int {
	odd := false
	oddIndex := 0
	var oddNumber int

	if len(seq) == 0 {
		return 0
	}

	for index, number := range seq {
		if odd && oddNumber == number {
			copy(seq[oddIndex:], seq[oddIndex+1:])
			copy(seq[index-1:], seq[index:])
			seq[len(seq)-1] = 0
			seq[len(seq)-2] = 0
			seq = seq[:len(seq)-2]
			oddIndex = 0
			oddNumber = 0
			odd = false
			break
		} else if !odd {
			odd = true
			oddNumber = number
			oddIndex = index
		}
	}

	if odd {
		return oddNumber
	}

	return FindOdd(seq)
}

// FindOddBest Best Practices
func FindOddBest(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}
