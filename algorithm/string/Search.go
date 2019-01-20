package string

func KMP(src string, sub string) int {
	cache := make([]int, len(sub))
	// generate cache
	for i := 0; i < len(cache); i++ {
		tmp := sub[0 : i + 1]
		prefix := getPrefix(tmp)
		suffix := getSuffix(tmp)
		for j := 0; j < len(tmp) - 1; j++ {
			if prefix[j] == suffix[j] {
				cache[i] = j + 1
			}
		}
	}

	for i := 0; i <= len(src) - len(sub); {
		isSuit := true
		for j := 0; j < len(sub); j++ {
			if src[i + j] != sub[j] {
				isSuit = false
				if j != 0 {
					i = i + j - cache[j - 1]
				} else {
					i++
				}
				break
			}
		}
		if isSuit {
			return i
		}
	}
	return -1
}

func getPrefix(src string) []string {
	result := make([]string, len(src) - 1)
	for i := 0; i < len(src) - 1; i++ {
		result[i] = src[0 : i + 1]
	}
	return result
}
func getSuffix(src string) []string {
	result := make([]string, len(src) - 1)
	for i := 0; i< len(src) - 1; i++ {
		result[i] = src[len(src) - i - 1 : ]
	}
	return result
}
