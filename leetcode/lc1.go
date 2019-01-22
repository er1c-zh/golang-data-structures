package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int][]int, 0)
	for idx, v := range nums {
		m[v] = append(m[v], idx)
	}

	for idx, v := range nums {
		t := target - v
		c, ok := m[t]

		if ok == false {
			continue
		}

		switch v == t {
		case true:
			if len(c) >= 2 {
				return c[0:2]
			}
		case false:
			if len(c) >= 1 {
				r := make([]int, 2)
				r[0] = idx
				r[1] = c[0]
				return r
			}
		}
	}

	return make([]int, 0)
}
