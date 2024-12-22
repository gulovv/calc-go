package calculation

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
    "unicode"
)

func Calc(expression string) (float64, error) {
    expression = strings.ReplaceAll(expression, " ", "")

    if !ValidateExpression(expression) { 
        return 0, ErrInvalidExpression
    }

    var values []float64
    var operators []rune

    applyOperator := func() error {
        if len(values) < 2 || len(operators) == 0 {
            return errors.New("некорректное выражение")
        }

        op := operators[len(operators)-1]
        operators = operators[:len(operators)-1]

        right := values[len(values)-1]
        left := values[len(values)-2]
        values = values[:len(values)-2]

        var result float64
        switch op {
        case '+':
            result = left + right
        case '-':
            result = left - right
        case '*':
            result = left * right
        case '/':
            if right == 0 {
                return ErrDivisionByZero
            }
            result = left / right
        }

        values = append(values, result)
        return nil
    }

    precedence := func(op rune) int {
        switch op {
        case '+', '-':
            return 1
        case '*', '/':
            return 2
        }
        return 0
    }

    for i := 0; i < len(expression); i++ {
        ch := rune(expression[i])

        if unicode.IsDigit(ch) || ch == '.' {
            start := i
            for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || rune(expression[i]) == '.') {
                i++
            }
            value, err := strconv.ParseFloat(expression[start:i], 64)
            if err != nil {
                return 0, errors.New("некорректное число")
            }
            values = append(values, value)
            i-- 
        } else if ch == '(' {
            operators = append(operators, ch) 
        } else if ch == ')' {
            
            for len(operators) > 0 && operators[len(operators)-1] != '(' {
                if err := applyOperator(); err != nil {
                    return 0, err
                }
            }
            if len(operators) == 0 || operators[len(operators)-1] != '(' {
                return 0, errors.New("некорректное выражение")
            }
            operators = operators[:len(operators)-1] 
        } else if ch == '+' || ch == '-' || ch == '*' || ch == '/' {
           
            for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(ch) {
                if err := applyOperator(); err != nil {
                    return 0, err
                }
            }
            operators = append(operators, ch)
        } else {
            return 0, fmt.Errorf("некорректный символ: %c", ch)
        }
    }

    
    for len(operators) > 0 {
        if err := applyOperator(); err != nil {
            return 0, err
        }
    }

    if len(values) != 1 {
        return 0, errors.New("некорректное выражение")
    }
    return values[0], nil
}