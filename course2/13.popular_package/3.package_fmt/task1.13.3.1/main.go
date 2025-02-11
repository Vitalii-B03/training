package main

import (
	"fmt"
	"strconv"
)

func generateMathString(operands []int, operator string) string {
	var mathstr string
	res := 0
	switch operator {
	case "+":
		for _, value := range operands {
			res += value
		}
	case "-":
		for _, value := range operands {
			res -= value
		}
	case "*":
		res = 1
		for _, value := range operands {
			res = value * res
			fmt.Println(res)
		}
	case "/":
		res = operands[0]
		for i := 1; i < len(operands); i++ {
			res /= operands[i]
		}
	}
	if len(operands) == 1 {
		mathstr = strconv.Itoa(operands[0])
		return mathstr
	} else {
		mathstr = strconv.Itoa(operands[0])
		for i := 1; i < len(operands); i++ {
			mathstr += fmt.Sprintf(" %s %d", operator, operands[i])
		}
		return fmt.Sprintf("%s = %d", mathstr, res)
	}
}
func main() {
	fmt.Println(generateMathString([]int{20, 2, 5}, "/"))
}
