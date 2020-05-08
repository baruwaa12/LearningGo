package sublist


// Collection array of integers
type Collection []int

// Relation type string
type Relation string

func isEqual(first, second Collection) bool {
	if len(first) != len(second) {
		return false
	}

	for index, value := range first {
		if value != second[index] {
			return false
		}
	}

	return true
}

func isSublist(first, second Collection) bool {
	firstSize, secondSize := len(first), len(second)

	for index := 0; index <= secondSize-firstSize; index++ {
		if isEqual(first, second[index:index+firstSize]) {
			return true
		}
	}

	return false
}

// Sublist Function
func Sublist(first, second Collection) Relation {
	switch {
	case isEqual(first, second):
		return "equal"
	case isSublist(first, second):
		return "sublist"
	case isSublist(second, first):
		return "superlist"
	}

	return "unequal"
}