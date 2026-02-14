package to_sum

func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, num := range nums {
		if j, ok := temp[num]; ok {
			return []int{j, i}
		}
		temp[target-num] = i
	}
	return nil
}
