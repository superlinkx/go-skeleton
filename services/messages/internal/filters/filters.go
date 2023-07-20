package filters

import "golang.org/x/exp/constraints"

func FilterToOddIntegers[T constraints.Integer](integers []T) []T {
	var oddIntegers = make([]T, 0, len(integers)/2)

	for _, id := range integers {
		if id%2 == 1 {
			oddIntegers = append(oddIntegers, id)
		}
	}

	return oddIntegers
}
