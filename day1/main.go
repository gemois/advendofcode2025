package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
)

func main() {
    file,_ := os.Open(os.Args[1])
    defer file.Close()
		scanner := bufio.NewScanner(file)

		zero_landing_count := 0
		zero_passed_count := 0

		current_pointer := 50

    for scanner.Scan() {
			line := scanner.Text()

			dial_direction := line[0]
			dial_amount,_ := strconv.Atoi(line[1:])
			dial_remainder := dial_amount % 100

			if dial_direction == 'L' {
				if current_pointer != 0 {
					zero_passed_count = zero_passed_count + ((dial_amount + (100 - current_pointer)) / 100)
				} else {
					zero_passed_count = zero_passed_count + (dial_amount / 100)
				}
				current_pointer = (current_pointer - dial_remainder) % 100
			} else {
				zero_passed_count = zero_passed_count + ((dial_amount + current_pointer) / 100)
				current_pointer = (current_pointer + dial_remainder) % 100
			}

			if current_pointer < 0 {
				current_pointer =  current_pointer + 100
			}

			if current_pointer == 0 {
				zero_landing_count = zero_landing_count + 1
			}
		}

		fmt.Println(zero_landing_count)
		fmt.Println(zero_passed_count)
}
