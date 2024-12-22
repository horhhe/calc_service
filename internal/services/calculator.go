package services

import (
	"errors"
	"strconv"
	"strings"
)

// Calc вычисляет арифметическое выражение
func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0, errors.New("пустое выражение")
	}

	precedence := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	output := []string{}
	operators := []rune{}

	var numberBuffer strings.Builder
	for _, ch := range expression {
		if (ch >= '0' && ch <= '9') || ch == '.' {
			numberBuffer.WriteRune(ch)
		} else {
			if numberBuffer.Len() > 0 {
				output = append(output, numberBuffer.String())
				numberBuffer.Reset()
			}
			if ch == '(' {
				operators = append(operators, ch)
			} else if ch == ')' {
				for len(operators) > 0 && operators[len(operators)-1] != '(' {
					output = append(output, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				if len(operators) == 0 {
					return 0, errors.New("несоответствующая скобка")
				}
				operators = operators[:len(operators)-1]
			} else if precedence[ch] > 0 {
				for len(operators) > 0 && precedence[operators[len(operators)-1]] >= precedence[ch] {
					output = append(output, string(operators[len(operators)-1]))
					operators = operators[:len(operators)-1]
				}
				operators = append(operators, ch)
			} else {
				return 0, errors.New("недопустимый символ: " + string(ch))
			}
		}
	}
	if numberBuffer.Len() > 0 {
		output = append(output, numberBuffer.String())
		numberBuffer.Reset()
	}

	for len(operators) > 0 {
		if operators[len(operators)-1] == '(' {
			return 0, errors.New("несоответствующая скобка")
		}
		output = append(output, string(operators[len(operators)-1]))
		operators = operators[:len(operators)-1]
	}

	return evaluatePostfix(output)
}

func evaluatePostfix(tokens []string) (float64, error) {
	stack := []float64{}

	for _, token := range tokens {
		if isNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, err
			}
			stack = append(stack, num)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("недостаточно операндов для операции: " + token)
			}
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, errors.New("деление на ноль")
				}
				stack = append(stack, a/b)
			default:
				return 0, errors.New("недопустимая операция: " + token)
			}
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("неверное выражение")
	}

	return stack[0], nil
}

func isNumber(token string) bool {
	_, err := strconv.ParseFloat(token, 64)
	return err == nil
}
