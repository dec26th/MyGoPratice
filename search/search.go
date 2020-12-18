package search

type Filter func(s []byte) bool

func getDoubleLoop(target []byte) Filter {
	return func(s []byte) bool {
		for i := 0; i < len(s); i++ {
			if !isInSlice(s[i], target) {
				return false
			}
		}
		return true
	}
}

func isInSlice(b byte, s []byte) bool {
	for i := 0; i < len(s); i++ {
		if b != s[i] {
			return false
		}
	}

	return true
}

func getMap(target []byte) Filter {
	return func(s []byte) bool {
		mapper := make(map[byte]int8, len(s))
		for i := 0; i < len(s); i++ {
			mapper[s[i]] = 1
		}

		for i := 0; i < len(target); i++ {
			if _, ok := mapper[target[i]]; !ok {
				return false
			}
		}

		return true
	}
}
