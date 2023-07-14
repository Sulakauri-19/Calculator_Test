package main

import (
	"fmt"
	"strconv" //  для преобразования чисел из строк в целые числа
)

var romans = map[rune]int{ // отображает римские цифры на соответствующие им значения
	'I': 1,
	'V': 5,
	'X': 10}

var intToRomans = map[int]string{ // обратное отображение исх. значений в римские цифры
	100: "C",
	90:  "XC",
	50:  "L",
	10:  "X",
	9:   "IX",
	5:   "V",
	4:   "IV",
	1:   "I"} // все числа вне данного набора приведут к ошибке

var ints = []int{100, 90, 50, 10, 9, 5, 4, 1} // массив для преобразования арабских цифр(чисел) в арабские для itor

func in(char rune, chars []rune) bool { // здесь проверяем, является ли текущий символ римской цифрой
	for _, ch := range chars {
		if char == ch {
			return true
		}
	}
	return false
}

func parseRome(snum string) int { // возвращает число по римской системе исчесления
	num := 0
	chars := []rune{'V', 'X'}

	for i := 0; i < len(snum); i++ {
		if i+1 >= len(snum) {
			num += romans[rune(snum[i])]
		} else {
			if in(rune(snum[i+1]), chars) && rune(snum[i]) == 'I' {
				num -= romans[rune(snum[i])]
			} else {
				num += romans[rune(snum[i])]
			}
		}
	}

	right_rome := itor(num)

	if snum != right_rome {
		panic("Ошибка: неправильно введено римское число")
	}

	return num
}

func parse(snum string) (int, bool) { // преобразовывает в целое число; иначе - в римское
	num, err := strconv.Atoi(snum)

	if err != nil {
		num = parseRome(snum)
		return num, true
	}

	return num, false
}

func itor(num int) string { // итерируется по массиву ints и определяет, добавляет римскую цифру к результату
	res := ""

	if num < 1 {
		panic("Ошибка: римские числа не могут быть < 1")

	}

	for _, v := range ints {
		count := num / v
		for i := 0; i < count; i++ {
			res += intToRomans[v]
		}
		num -= v * count
	}

	return res
}

func main() {
	fmt.Println("Калькулятор запущен")
	fmt.Println("-------------------")

	var snum1, snum2 string

	var operator string

	fmt.Print("Введите первое число, затем знак сочетания (например: +, -, *, /), затем второе число: ")
	fmt.Scan(&snum1, &operator, &snum2)

	num1, isRome1 := parse(snum1)
	num2, isRome2 := parse(snum2)

	if isRome1 != isRome2 {
		panic("Ошибка: нельзя вычислять арабские и римские числа одновременно")
	}

	if num1 > 10 || num2 > 10 {
		panic("Ошибка: можно только числа <= 10")
	}

	if num1 <= 0 || num2 <= 0 {
		panic("Ошибка: можно только числа > 0")
	}

	// БЛОК ОПЕРАЦИЙ

	result := 0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль")
			return
		}

		result = num1 / num2
	default:
		fmt.Println("Ошибка: неверный оператор")
		return
	}

	if isRome1 {
		fmt.Println("Вывод:", itor(result))
	} else {
		fmt.Println("Вывод:", result)
	}

	// ИТОГОВЫЙ ВЫВОД

	fmt.Println("-------------------")

}
