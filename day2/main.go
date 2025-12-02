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

		scanner.Scan()
		line := scanner.Text()
		id_ranges := strings.Split(line, ",")

		invalid_id_exactly_twice_sum := 0
		invalid_id_atleast_twice_sum := 0

		for _,id_range := range id_ranges {
			boundaries := strings.Split(id_range, "-")

			left_boundary,_ := strconv.Atoi(boundaries[0])
			righ_boundary,_ := strconv.Atoi(boundaries[1])

			// part 1
			for id := left_boundary; id <= righ_boundary; id++ {
				middle_of_number := len(strconv.Itoa(id)) / 2
				left_sequence := strconv.Itoa(id)[0:middle_of_number]
				right_sequence := strconv.Itoa(id)[middle_of_number:]

				if (left_sequence == right_sequence) {
					invalid_id_exactly_twice_sum += id
				}
			}

			// part 2
			for id := left_boundary; id <= righ_boundary; id++ {
				equal_len_substrings := findEqualLenSubStrings(id)

				invalid := false
				for _, equal_len_substring := range equal_len_substrings {
					first_substring := strconv.Itoa(id)[:equal_len_substring]
       
					if (strings.Count(strconv.Itoa(id), first_substring) == len(strconv.Itoa(id)) / equal_len_substring) {
						invalid = true
					}
				}

				if invalid {
					invalid_id_atleast_twice_sum += id
				}
			}

		}
		
	fmt.Println(invalid_id_exactly_twice_sum)
	fmt.Println(invalid_id_atleast_twice_sum)
}

func findEqualLenSubStrings(id int) []int {
	equal_len_substrings := []int{}

	for i := 1; i < len(strconv.Itoa(id)); i++ {
		if len(strconv.Itoa(id)) % i == 0 {
			equal_len_substrings = append(equal_len_substrings, i)
		}
	}

	return equal_len_substrings
}
