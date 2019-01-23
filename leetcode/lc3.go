package leetcode

func lengthOfLongestSubstring(s string) int {
	ba := []byte(s)
	if len(ba) == 0 {
		return 0
	}
	m := make(map[byte]int, 128)
	r := 0
	c := 1

	iSlow := 0
	iFast := 1
	m[ba[iSlow]] = iSlow

	for iFast < len(ba) {
		cByte := ba[iFast]
		idx, ok := m[cByte]
		if ok == true && idx >= iSlow {
			r = max(r, c)
			c -= idx - iSlow + 1
			iSlow = idx + 1

			if len(ba) - iSlow +1 < r {
				return r
			}
		}
		m[cByte] = iFast
		c++
		iFast++
	}

	r = max(r, c)

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}