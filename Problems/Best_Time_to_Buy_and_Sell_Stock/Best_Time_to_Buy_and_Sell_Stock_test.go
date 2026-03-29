package Best_Time_to_Buy_and_Sell_Stock

import "testing"

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	tests := []struct {
		prices   []int
		expected int
	}{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}

	for _, test := range tests {
		result := BestTimeToBuyAndSellStock(test.prices)
		if result != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, result)
		}
	}
}
