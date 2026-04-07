package evaluate_arithmetic_expression

func evaluateArithmeticExpression(expr interface{}) int {
	switch val := expr.(type) {
	case int:
		return val
	case []interface{}:
		op := val[0].(string)
		switch op {
		case "add":
			return evaluateArithmeticExpression(val[1]) + evaluateArithmeticExpression(val[2])
		case "sub":
			return evaluateArithmeticExpression(val[1]) - evaluateArithmeticExpression(val[2])
		case "mult":
			return evaluateArithmeticExpression(val[1]) * evaluateArithmeticExpression(val[2])
		case "div":
			return evaluateArithmeticExpression(val[1]) / evaluateArithmeticExpression(val[2])
		default:
			panic("unknown operator")
		}
	default:
		panic("unknown expression")
	}
}
