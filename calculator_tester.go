package main

import (
	"fmt"
	"log"
	"strconv"
)

var romans = map[string]int{
	"X":    10,
	"IX":   9,
	"VIII": 8,
	"VII":  7,
	"VI":   6,
	"V":    5,
	"IV":   4,
	"III":  3,
	"II":   2,
	"I":    1}

func main() {
	fmt.Println("Калькулятор запущен")
	fmt.Println("-------------------")

	var num1, num2 string

	var operator string

	fmt.Print("Введите первое число, затем знак сочетания (например: +, -, *, /), затем второе число: ")
	fmt.Scan(&num1, &operator, &num2)

	// КОНВЕРТАТОР ПЕРЕМЕННЫХ
	proobraz1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal(err)
	}

	proobraz2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal(err)
	}

	// БЛОК ОГРАНИЧЕНИЯ ВВОДА ЧИСЕЛ БОЛЬШЕ 10 ИЛИ МЕНЬШЕ 0
	if proobraz1 > 10 {
		fmt.Print("Ошибка: недопустимое число №1")
	}
	if proobraz2 > 10 {
		fmt.Print("Ошибка: недопустимое число №2")
	}

	if proobraz1 < 0 {
		fmt.Print("Ошибка: недопустимое число №1")
	}
	if proobraz2 < 0 {
		fmt.Print("Ошибка: недопустимое число №2")
	}

	// БЛОК ОПЕРАЦИЙ

	result := 0

	switch operator {

	case "+":
		result = proobraz1 + proobraz2
		fmt.Println("Вывод:", result)

	case "-":
		result = proobraz1 - proobraz2
		fmt.Println("Вывод:", result)

	case "*":
		result = proobraz1 * proobraz2
		fmt.Println("Вывод:", result)

	case "/":
		if proobraz2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}

		result = proobraz1 / proobraz2
		fmt.Println("Вывод:", result)

	default:
		fmt.Println("Ошибка: неверный оператор")

		return
	}

	// ИТОГОВЫЙ ВЫВОД

	fmt.Println("-------------------")

}
