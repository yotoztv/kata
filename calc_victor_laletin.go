package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rome_num_10(str1 string) int {
	if str1 == "I" {
		return 1
	} else if str1 == "II" {
		return 2
	} else if str1 == "III" {
		return 3
	} else if str1 == "IV" {
		return 4
	} else if str1 == "V" {
		return 5
	} else if str1 == "VI" {
		return 6
	} else if str1 == "VII" {
		return 7
	} else if str1 == "VIII" {
		return 8
	} else if str1 == "IX" {
		return 9
	} else if str1 == "X" {
		return 10
	} else {
		return 0
	}
}

func it_num(str1 string) (int, int) {
	x, err1 := strconv.Atoi(str1)

	r := rome_num_10(str1)

	if err1 == nil && x > 0 && x < 11 {
		return x, 0
	} else if r > 0 {
		return r, 1
	} else {
		return 0, 0
	}
}

func str_rome(x int) string {

	str1 := ""

	nu_1 := strings.Split("0 I II III IV V VI VII VIII IX 77", " ")
	nu_10 := strings.Split("0 X XX XXX XL L LX LXX LXXX XC 77", " ")
	// элементы "0" добавлены для удобства соответствия 1-1 9-9,
	// а "77" для наглядности бага ошибок, если выйду за диапазон

	if x == 100 {
		return "C"
	}
	if x/10 > 0 {
		str1 = str1 + nu_10[x/10]
		x = x % 10
	}
	if x/1 > 0 {
		str1 = str1 + nu_1[x]
	}
	return str1
}

func main() {

	var input_str string
	var result int
	var size int
	var a, b, rome_a, rome_b int

	input_str, _ = bufio.NewReader(os.Stdin).ReadString('\n')

	input_str = strings.Trim(input_str, "\n")
	input_str = strings.TrimSpace(input_str)

	box := strings.Split(input_str, " ")

	size = len(box)

	if size != 3 {
		fmt.Println("Элементов должно быть 3, и они должны быть разделены одиночными пробелами")
		os.Exit(1)
	}

	a, rome_a = it_num(box[0])
	b, rome_b = it_num(box[2])

	if !(a != 0 && b != 0) {
		fmt.Println("Аргументы должны быть в интервале от 1 до 10 включительно")
		os.Exit(1)
	}

	if rome_a != rome_b {
		fmt.Println("Аргументы должны быть либо римскими либо арабскими оба два")
		os.Exit(1)
	}

	if box[1][0] == '+' && len(box[1]) == 1 {
		result = a + b
	} else if box[1][0] == '-' {
		result = a - b
	} else if box[1][0] == '*' {
		result = a * b
	} else if box[1][0] == '/' {
		result = a / b
	} else {
		fmt.Println("Не верная математическая операция, допускается: +, -, *, /")
		os.Exit(1)
	}

	if result <= 0 && rome_b > 0 {
		fmt.Println("В Римских цифрах может быть только положительный результат от  I(1) и выше")
		os.Exit(1)
	}

	if rome_a == 0 {
		fmt.Println(result)
	} else {
		fmt.Println(str_rome(result))
	}

}
