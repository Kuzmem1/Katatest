package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	m1 := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10}

	m := map[string]int{
		"1":  1,
		"2":  2,
		"3":  3,
		"4":  4,
		"5":  5,
		"6":  6,
		"7":  7,
		"8":  8,
		"9":  9,
		"10": 10,
	}
	var num1, num2, operation string
	var wrong string
	fmt.Scanln(&num1, &operation, &num2, &wrong)
	err0 := errors.New("Одно из чисел меньше 1(I) или больше 10 (X) \n")
	Roman := "I II III IV V VI VII VIII IX X"
	Arabic := "1 2 3 4 5 6 7 8 9 10"
	err1 := errors.New("Римское число не может быть отрицательным \n")
	err2 := errors.New("Нужно выражение формата A + B \n")
	err3 := errors.New("В одном выражении не могут быть римские и арабские числа \n")
	lookfor := num2
	contain := strings.Contains(Roman, lookfor)

	lookfor1 := num2
	contain1 := strings.Contains(Arabic, lookfor1)
	value, _ := m[num1]
	value1, _ := m[num2]
	value2, _ := m1[num1]
	value3, _ := m1[num2]

	if wrong == "" {
		if value >= 1 && value <= 10 {

			if value1 >= 1 && value1 <= 10 {
				switch operation {

				case "-":
					fmt.Printf("%d \n", value-value1)

				case "+":
					fmt.Printf("%d \n", value+value1)

				case "*":
					fmt.Printf("%d \n", value*value1)

				case "/":
					fmt.Printf("%d \n", value/value1)

				default:
					fmt.Print("Неверное выражение")
				}
			} else if contain == true {
				fmt.Print(err3)
			} else if value1 > 10 || value1 < 1 {
				fmt.Print(err0)
			}
		} else if _, err := strconv.Atoi(num1); err == nil {
			fmt.Print(err0)
		} else {
			if value2 >= 1 && value2 <= 10 {

				if value3 >= 1 && value3 <= 10 {

					m2 := map[int]string{
						1: "I", 11: "XI", 21: "XXI", 31: "XXXI", 41: "XLI", 51: "L", 61: "LXI", 71: "LXXI", 81: "LXXXI", 91: "XCI",
						2: "II", 12: "XII", 22: "XXII", 32: "XXXII", 42: "XLII", 52: "LII", 62: "LXII", 72: "LXXII", 82: "LXXXII", 92: "XCII",
						3: "III", 13: "XIII", 23: "XXIII", 33: "XXXIII", 43: "XLIII", 53: "LIII", 63: "LXIII", 73: "LXXIII", 83: "LXXXIII", 93: "XCIII",
						4: "IV", 14: "XIV", 24: "XXIV", 34: "XXXIV", 44: "XLIV", 54: "LIV", 64: "LXIV", 74: "LXXIV", 84: "LXXXIV", 94: "XCIV",
						5: "V", 15: "XV", 25: "XXV", 35: "XXXV", 45: "XLV", 55: "LV", 65: "LXV", 75: "LXXV", 85: "LXXXV", 95: "XCV",
						6: "VI", 16: "XVI", 26: "XXVI", 36: "XXXVI", 46: "XLVI", 56: "LVI", 66: "LXVI", 76: "LXXVI", 86: "LXXXVI", 96: "XCVI",
						7: "VII", 17: "XVII", 27: "XXVII", 37: "XXXVII", 47: "XLVII", 57: "LVII", 67: "LXVII", 77: "LXXVII", 87: "LXXXVII", 97: "XCVII",
						8: "VIII", 18: "XVIII", 28: "XXVIII", 38: "XXXVIII", 48: "XLVIII", 58: "LVIII", 68: "LXVIII", 78: "LXXVIII", 88: "LXXXVIII", 98: "XCVIII",
						9: "IX", 19: "XIX", 29: "XXIX", 39: "XXXIX", 49: "XLIX", 59: "LIX", 69: "LXIX", 79: "LXXIX", 89: "LXXXIX", 99: "XCIX",
						10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
					}
					Value4, _ := m2[value2-value3]
					Value5, _ := m2[value2+value3]
					Value6, _ := m2[value2*value3]
					Value7, _ := m2[value2/value3]

					switch operation {

					case "-":
						if value2-value3 < 1 {
							fmt.Println(err1)
						} else {
							fmt.Printf("%s \n", Value4)
						}
					case "+":
						fmt.Printf("%s \n", Value5)

					case "*":
						fmt.Printf("%s \n", Value6)

					case "/":
						fmt.Printf("%s \n", Value7)

					default:
						fmt.Print("Неверное выражение")
					}
				} else if contain1 == true {
					fmt.Print(err3)

				} else {
					fmt.Print(err2)
				}
			} else {
				fmt.Print(err0)
			}
		}
	} else {
		fmt.Print(err2)
	}
}