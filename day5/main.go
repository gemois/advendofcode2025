package main

import (
    "bufio"
    "fmt"
    "os"
		"strconv"
		"strings"
		"sort"
)

func main() {
    file,_ := os.Open(os.Args[1])
    defer file.Close()
		scanner := bufio.NewScanner(file)

		db_separator_reached := false

		var ingrediient_ids []string
		var ranges []string
		
		for scanner.Scan() {
			line := scanner.Text()

			if (len(line) == 0) {
				db_separator_reached = true
				continue
			}

			if (db_separator_reached) {
				ingrediient_ids = append(ingrediient_ids, line)
			} else {
				ranges = append(ranges, line)
			}
		}

		// part 1
		fresh_ingredients := 0
		for _, ingrediient_id_str := range ingrediient_ids {
			ingrediient_id, _ := strconv.Atoi(ingrediient_id_str)
			for _, r := range ranges {
				the_range := strings.Split(r, "-")
				left, _ := strconv.Atoi(the_range[0])
				right, _ := strconv.Atoi(the_range[1])

				if ingrediient_id >= left && ingrediient_id <= right {
					fresh_ingredients++
					break
				}
			}
		}

		fmt.Println(fresh_ingredients)

		// part 2
		sort.Slice(ranges, func(i, j int) bool {
			current, _ := strconv.Atoi(strings.Split(ranges[i], "-")[0])
			next, _ := strconv.Atoi(strings.Split(ranges[j], "-")[0])
			return current < next
		})

		var merged_ranges []string

		var current_left int
		var current_right int

		for i, r := range ranges {
			parts := strings.Split(r, "-")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])

			if i == 0 {
				current_left = left
				current_right = right
				continue
			}

			if left > current_right + 1 {
				merged_ranges = append(merged_ranges, fmt.Sprintf("%d-%d", current_left, current_right))
				current_left = left
				current_right = right
				continue
			}

			if right > current_right {
				current_right = right
			}
		}

		merged_ranges = append(merged_ranges, fmt.Sprintf("%d-%d", current_left, current_right))

		all_fresh_ingredients := 0
		for _, r := range merged_ranges {
			parts := strings.Split(r, "-")
			left, _ := strconv.Atoi(parts[0])
			right, _ := strconv.Atoi(parts[1])
			all_fresh_ingredients += right - left + 1
		}

		fmt.Println(all_fresh_ingredients)
	}
