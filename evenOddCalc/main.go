package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите 10 целых чисел, а я из них выведу количество чётных и нечётных:")

	var numbers [10]int
	var evens int
	var odds int

	for i := 0; i < len(numbers); i++ {
		var tmp int
		for {
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			var err error
			tmp, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Ошибка: введите корректное целое число.")
				continue
			} else {
				break
			}
		}
		numbers[i] = tmp
	}

	for i, _ := range numbers {
		if numbers[i]%2 == 0 {
			evens++
		} else {
			odds++
		}
	}
	fmt.Println("Все ваши числа:", numbers)
	fmt.Println("Количество чётных чисел:", evens)
	fmt.Println("Количество нечётных чисел:", odds)
}
