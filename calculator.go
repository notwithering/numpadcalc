package main

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/atotto/clipboard"
)

var keyCh = make(chan xproto.Keycode)

func evaluate(expr string) (*big.Float, error) {
	if !strings.ContainsAny(expr, "+-*/") {
		f, _, err := big.ParseFloat(expr, 10, 100, big.ToPositiveInf)
		if err != nil {
			return nil, err
		}
		return f, nil
	}

	var tokens []string
	var currentToken string
	for _, char := range expr {
		if char == ' ' {
			continue
		}
		if char == '+' || char == '-' || char == '*' || char == '/' {
			if currentToken != "" {
				tokens = append(tokens, currentToken)
			}
			tokens = append(tokens, string(char))
			currentToken = ""
		} else {
			currentToken += string(char)
		}
	}
	if currentToken != "" {
		tokens = append(tokens, currentToken)
	}

	var values []*big.Float
	var operators []string
	for i, token := range tokens {
		if i%2 == 0 {
			f, _, err := big.ParseFloat(token, 10, 100, big.ToPositiveInf)
			if err != nil {
				return nil, err
			}
			values = append(values, f)
		} else {
			operators = append(operators, token)
		}
	}

	for i := 0; i < len(operators); i++ {
		if operators[i] == "*" || operators[i] == "/" {
			op := operators[i]
			a := values[i]
			b := values[i+1]
			var result *big.Float
			if op == "*" {
				result = new(big.Float).Mul(a, b)
			} else {
				result = new(big.Float).Quo(a, b)
			}
			values = append(values[:i], append([]*big.Float{result}, values[i+2:]...)...)
			operators = append(operators[:i], operators[i+1:]...)
			i--
		}
	}

	result := values[0]
	for i := 1; i < len(values); i++ {
		op := operators[i-1]
		operand := values[i]
		if op == "+" {
			result.Add(result, operand)
		} else {
			result.Sub(result, operand)
		}
	}

	return result, nil
}

func calculator() {
	var formula strings.Builder

	for keycode := range keyCh {
		key := keyMap[keycode]

		if formula.Len() == 0 {
			formula.WriteByte('0')
			if isOperator(key) {
				s, _ := clipboard.ReadAll()
				f, _, err := new(big.Float).Parse(s, 10)
				if err == nil {
					formula.WriteString(f.String())
				}
			}
		}

		if key == "enter" {
			n, err := evaluate(formula.String())
			fmt.Println(n, err)

			formula.Reset()

			if err != nil {
				clipboard.WriteAll(err.Error())
				continue
			}

			clipboard.WriteAll(fmt.Sprint(n))
			continue
		}

		formula.WriteString(key)

		fmt.Println(formula.String())
	}
}

func isOperator(s string) bool {
	switch s {
	case "+", "-", "*", "/":
		return true
	}
	return false
}
