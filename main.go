package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toRoman(number int) string {
	nums := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	symb := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var result string = ""

	for number > 0 {
		for key, _ := range nums {
			if number >= nums[key] {
				result += symb[key]
				number -= nums[key]
				break
			}
		}
	}
	return result
}

func main() {

	//make map
	fromRomanNum := map[string]int8{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	toRomanNum := make(map[int8]string)
	for key, value := range fromRomanNum {
		toRomanNum[value] = key
	}
	//---

	//init vars
	var firstValue int8
	var secondValue int8
	var RomanNum bool
	var question string
	var operand string
	//

	reader := bufio.NewReader(os.Stdin)
	//get question
	fmt.Println("Введите выражение: ")
	question, _ = reader.ReadString('\n')

	question = strings.ToUpper(strings.ReplaceAll(strings.TrimSpace(question), " ", ""))
	//---

	//get operand
	if i := strings.Count(question, "+"); i == 1 && len(question) >= 3 {
		operand = "+"
	} else if i := strings.Count(question, "-"); i == 1 && len(question) >= 3 {
		operand = "-"
	} else if i := strings.Count(question, "*"); i == 1 && len(question) >= 3 {
		operand = "*"
	} else if i := strings.Count(question, "/"); i == 1 && len(question) >= 3 {
		operand = "/"
	} else {
		fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
		return
	}

	{
		add := strings.Contains(question, "+")
		sub := strings.Contains(question, "-")
		mul := strings.Contains(question, "*")
		div := strings.Contains(question, "/")

		if (add && sub || add && mul || add && div) || (sub && mul || sub && div) || mul && div {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
	}
	//---

	//get values
	{
		values := strings.SplitN(question, operand, 2)

		RomanNum = strings.ContainsAny(values[0], "IVX")
		if strings.ContainsAny(values[1], "IVX") != RomanNum {
			fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
			return
		}

		//FirstValue
		if RomanNum {
			firstValue = fromRomanNum[values[0]]
			//fmt.Println("getRoman", firstValue, values[0])
			if firstValue == 0 || firstValue > 10 {
				fmt.Println("Вывод ошибки, так как формат 1го операнда не удовлетворяет заданию - калькулятор должен принимать от 1 до 10 включительно и не более.")
				return
			}
		} else {
			number, err := strconv.Atoi(values[0])
			if err == nil && number <= 10 {
				firstValue = int8(number)
			} else {
				fmt.Println("Вывод ошибки, так как формат 1го операнда не удовлетворяет заданию - калькулятор должен принимать от 1 до 10 включительно и не более.")
				return
			}
		}

		//SecondValue
		if RomanNum {
			secondValue = fromRomanNum[values[1]]
			if secondValue == 0 || secondValue > 10 {
				fmt.Println("Вывод ошибки, так как формат 2го операнда не удовлетворяет заданию - калькулятор должен принимать от 1 до 10 включительно и не более.")
				return
			}
		} else {
			number, err := strconv.Atoi(values[1])
			if err == nil && number <= 10 {
				secondValue = int8(number)
			} else {
				fmt.Println("Вывод ошибки, так как формат 2го операнда не удовлетворяет заданию - калькулятор должен принимать от 1 до 10 включительно и не более.")
				return
			}
		}

	}
	//---

	//calculation
	if operand == "+" {
		firstValue = firstValue + secondValue
	} else if operand == "-" {
		firstValue = firstValue - secondValue
	} else if operand == "*" {
		firstValue = firstValue * secondValue
	} else if operand == "/" {
		firstValue = firstValue / secondValue
	} else {
		fmt.Println("Вывод неожиданной ошибки")
		return
	}

	if RomanNum {
		if firstValue <= 0 {
			fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			return
		}

		fmt.Println(toRoman(int(firstValue)))
	} else {
		println(firstValue)
	}
	//---
}
