package calc

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation struct {
	Expression string
	Result     float64
}

type Calculator struct {
	History []Operation
}

func (calc *Calculator) AddHistory(expression string, result float64) {
	calc.History = append(calc.History, Operation{
		Expression: expression,
		Result:     result,
	})
}

func (calc *Calculator) GetHistory() {
	if len(calc.History) == 0 {
		fmt.Println("История пуста")
		return
	}
	fmt.Println("История операций:")
	for _, v := range calc.History {
		fmt.Printf("%v = %v\n", v.Expression, v.Result)
	}
}

func (calc *Calculator) PerformCalculator(expression string) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		fmt.Println("Соблюдайте синтаксис!\n{Число операция число} : {2 + 2}")
		return
	}
	num1, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		fmt.Println("Ошибка конвертирования!", err)
		return
	}
	num2, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		fmt.Println("Ошибка конвертирования!", err)
		return
	}

	var result float64
	switch parts[1] {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	}

	calc.AddHistory(expression, result)
	fmt.Printf("Результат: %s = %.2f\n", expression, result)
}
