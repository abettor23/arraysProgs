package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverser(arr [10]int) [10]int {
	var reversedArr [10]int
	for i, j := len(arr)-1, 0; i >= 0 && j < len(reversedArr); i, j = i-1, j+1 {
		reversedArr[j] = arr[i]
	}
	return reversedArr
}

func main() {
	fmt.Println("Введите 10 целых чисел, а я их переверну!")

	var numbers [10]int

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

	result := reverser(numbers)

	fmt.Println("Ваши десять чисел:", numbers)
	fmt.Println("Ваши числа, но перевёрнутые:", result)
}
