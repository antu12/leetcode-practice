package evaluate_arithmetic_expression

import "testing"

func TestEvaluateArithmeticExpression(t *testing.T) {
	tests := []struct {
		operator interface{}
	}{
		{[]interface{}{"add", 1, 2}},
		{[]interface{}{"mult", 3, []interface{}{"add", 2, 4}}},
		{[]interface{}{"sub", []interface{}{"mult", 2, 3}, 4}},
		{[]interface{}{"div", []interface{}{"add", 10, 5}, 3}},
	}
	for _, tt := range tests {
		result := evaluateArithmeticExpression(tt.operator)
		t.Logf("evaluateArithmeticExpression(%v) = %v", tt.operator, result)
	}
}
