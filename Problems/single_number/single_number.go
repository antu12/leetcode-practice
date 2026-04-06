package single_number

func singleNumber(nums []int) int {
	temp := make(map[rune]int)

	for _, item := range nums {
		if _, ok := temp[rune(item)]; ok {
			temp[rune(item)]++
		} else {
			temp[rune(item)] = 1
		}
	}

	for k, v := range temp {
		if v == 1 {
			return int(k)
		}
	}
	return 0
}
