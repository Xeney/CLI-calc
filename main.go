package main

import (
	"bufio"
	"calc/calc"
	"fmt"
	"os"
)

func main() {
	calculator := calc.Calculator{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Добро пожаловать в CLI-калькулятор!")
	fmt.Println("Введите выражение в формате 'число оператор число' (например, '2 + 2').")
	fmt.Println("Для просмотра истории введите 'history'. Для выхода введите 'exit'.")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		// Выход из программы
		if input == "exit" {
			fmt.Println("Выход из программы.")
			break
		}

		// Просмотр истории
		if input == "history" {
			calculator.GetHistory()
			continue
		}

		// Выполнение операции
		calculator.PerformCalculator(input)
	}
}
