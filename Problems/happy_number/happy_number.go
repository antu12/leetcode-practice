package happy_number

func happyNumber(n int) bool {
	if n < 1 {
		return false
	}
	temp := make(map[int]int)
	for true {
		sum := 0
		for n > 0 {
			t := (n % 10)
			sum += t * t
			n = n / 10
		}
		n = sum
		if n == 1 {
			return true
		}
		if _, ok := temp[n]; ok {
			return false
		}
		temp[n]++
	}
	return false
}
