package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const alphaRoman = "IVXLCDM"

func ArabicToRoman(number int) (string, error) {
	if number < 1 {
		return "", errors.New("неверный расчет: результат меньше 1")
	}

	conversions := []struct {
		value   int
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.numeral
			number -= conversion.value
		}
	}
	return roman, nil
}

func DecodeToArabic(roman string) int {
	translateRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var decNum, tmpNum int
	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]
		decDigit := translateRoman[romanDigit]
		if decDigit < tmpNum {
			decNum -= decDigit
		} else {
			decNum += decDigit
			tmpNum = decDigit
		}
	}
	return decNum
}

func CheckString(input1, input2 string) (interface{}, error) {
	validArabic := map[string]bool{"1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true, "10": true}
	validRoman := map[string]bool{"I": true, "II": true, "III": true, "IV": true, "V": true, "VI": true, "VII": true, "VIII": true, "IX": true, "X": true}

	_, arabic1 := validArabic[input1]
	_, arabic2 := validArabic[input2]
	_, roman1 := validRoman[input1]
	_, roman2 := validRoman[input2]

	if arabic1 && arabic2 {
		num1, _ := strconv.Atoi(input1)
		num2, _ := strconv.Atoi(input2)
		return []int{num1, num2}, nil
	} else if roman1 && roman2 {
		return []string{input1, input2}, nil
	} else {
		return nil, errors.New("Некорректные значения. Введите 'I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X' или цифры от 1 до 10")
	}
}

func Calculate(x int, operator string, y int) int {
	var result int
	switch operator {
	case "+":
		result = x + y
	case "-":
		result = x - y
		if y > x {
			fmt.Println("Программа не работает с результатами меньше 0")
			os.Exit(1)
		}
	case "*":
		result = x * y
	case "/":
		if y == 0 {
			fmt.Println("Делить на 0 нельзя!")
			os.Exit(1)
		}
		result = x / y
	default:
		fmt.Println("Неверный оператор:", operator)
		os.Exit(1)
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите выражение: ")
	expression, _ := reader.ReadString('\n')

	expression = strings.TrimSpace(expression)

	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Строка не является математической операцией.")
		os.Exit(1)
	}

	x, operator, y := tokens[0], tokens[1], tokens[2]

	result, err := CheckString(x, y)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		switch res := result.(type) {
		case []int:
			num1, num2 := res[0], res[1]

			resultArabic := Calculate(num1, operator, num2)
			fmt.Println("Результат выражения:", resultArabic)
		case []string:
			str1, str2 := res[0], res[1]
			decode1, decode2 := DecodeToArabic(str1), DecodeToArabic(str2)
			calcRoman := Calculate(decode1, operator, decode2)
			resultRoman, err := ArabicToRoman(calcRoman)
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Результат выражения:", resultRoman)
			}
		}
	}
}
