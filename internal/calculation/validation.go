package calculation

import "unicode"

func ValidateExpression(expression string) bool {
	for _, ch := range expression {
		if !unicode.IsDigit(ch) && !isOperator(ch) && ch != '.' && ch != '(' && ch != ')' {
			return false
		}
	}
	return true
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}