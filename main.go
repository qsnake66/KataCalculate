package main

import (
	"fmt"
	"strconv"
	"strings"
  "bufio"
  "os"
)

func main() {


	fmt.Println("Введите выражение: число,оператор,число. Можно вводить числа до 10. Например 2+2")
	reader := bufio.NewReader(os.Stdin)
  text, _ := reader.ReadString('\n')
	calculate(text)

}

func calculate(text string) {

	tokens := strings.Split(text,  " ")


	var isRomanNum1 bool = false
	var isRomanNum2 bool = false
	var result int
	var rimResult string

	rimNum := map[string](int){
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20, "XXI": 21, "XXII": 22, "XXIII": 23, "XXIV": 24, "XXV": 25, "XXVI": 26, "XXVII": 27, "XXVIII": 28, "XXIX": 29, "XXX": 30, "XXXI": 31, "XXXII": 32, "XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37, "XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42, "XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47, "XLVIII": 48, "XLIX": 49, "L": 50, "LI": 51, "LII": 52, "LIII": 53, "LIV": 54, "LV": 55, "LVI": 56, "LVII": 57, "LVIII": 58, "LIX": 59, "LX": 60, "LXI": 61, "LXII": 62, "LXIII": 63, "LXIV": 64, "LXV": 65, "LXVI": 66, "LXVII": 67, "LXVIII": 68, "LXIX": 69, "LXX": 70, "LXXI": 71, "LXXII": 72, "LXXIII": 73, "LXXIV": 74, "LXXV": 75, "LXXVI": 76, "LXXVII": 77, "LXXVIII": 78, "LXXIX": 79, "LXXX": 80, "LXXXI": 81, "LXXXII": 82, "LXXXIII": 83, "LXXXIV": 84, "LXXXV": 85, "LXXXVI": 86, "LXXXVII": 87, "LXXXVIII": 88, "LXXXIX": 89, "XC": 90, "XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95, "XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	}

	if len(tokens) != 3 {
		fmt.Print("Должно быть три символа. Число, оператор и число")
		return
	}
	first_Number := strings.TrimSpace(tokens[0])
	operator := strings.TrimSpace(tokens[1])
	second_Number := strings.TrimSpace(tokens[2])


	num1, err := strconv.Atoi(first_Number)

	if err != nil {
		num1 = rimNum[first_Number]
		if num1 == 0 {
			fmt.Println("Необходимо вводить только римские или арабские числа")
			return
		} else {
			isRomanNum1 = true
		}
	}

	num2, err := strconv.Atoi(second_Number)

	if err != nil {
		num2 = rimNum[second_Number]
		if num2 == 0 {
			fmt.Println("Необходимо вводить только римские или арабские числа")
			return
		} else {
			isRomanNum2 = true
		}
	}


	if (num1 < 1 || num1 > 10) || (num2 < 1 || num2 > 10) {
		fmt.Println("Числа можно ввести от 1 до 10 включитель")
		return
	}

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "/":
		result = num1 / num2
	case "*":
		result = num1 * num2
	default:
		fmt.Println("Недопустимый оператор")
	}

	if (isRomanNum1 && isRomanNum2) {
		if result > 0 {
			for key, value := range rimNum {
				if value == result {
					rimResult = key
					fmt.Println("Результат: ", rimResult)
					return
				}
			}
		} else {
			fmt.Println("Римские цифры не могут быть равны 0 или меньше 0")
		}
	} else if (isRomanNum1 == true && isRomanNum2 == false) || (isRomanNum1 == false && isRomanNum2 == true) {
		fmt.Println("Либо только римские либо арабские")
	} else {
		fmt.Println("результат: ", result)
	}

}
