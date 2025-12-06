package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
		"strings"
)

func main() {
    file,_ := os.Open(os.Args[1])
    defer file.Close()
		scanner := bufio.NewScanner(file)

		var all_fields [][]string

		for scanner.Scan() {
			line := scanner.Text()
			fields := strings.Fields(line)
			all_fields = append(all_fields, fields)
		}

		operators := all_fields[len(all_fields) - 1]
		number_lines := all_fields[:len(all_fields) - 1]
		columns := len(all_fields[0])

		total_sum := 0
		for column := 0; column < columns; column++ {
			operator := operators[column]

			column_sum := 0
			if operator == "*" {
				column_sum = 1
			}

			for number_line := 0; number_line < len(number_lines); number_line++ {
				number_str := number_lines[number_line][column]
				number,_ := strconv.Atoi(number_str)
				if operator == "+" {
					column_sum += number
				} else {
					column_sum *= number
				}
			}

			total_sum += column_sum
		}

		fmt.Println(total_sum)
}
