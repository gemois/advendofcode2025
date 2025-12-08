package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var manifold []string
	for scanner.Scan() {
		manifold = append(manifold, scanner.Text())
	}

	times_splitted := 0

	start := strings.Index(manifold[0], "S")
	current_beam_positions := []int{ start }

	for row := 1; row < len(manifold); row++ {
		line := manifold[row]

		for i := 0; i < len(current_beam_positions); {
			position := current_beam_positions[i]

			if line[position] == '^' {
				times_splitted++

				current_beam_positions = append(current_beam_positions[:i], current_beam_positions[i + 1:]...)

				if position - 1 >= 0 && !slices.Contains(current_beam_positions, position - 1) {
					current_beam_positions = append(current_beam_positions, position - 1)
				}

				if position + 1 < len(line) && !slices.Contains(current_beam_positions, position + 1) {
					current_beam_positions = append(current_beam_positions, position + 1)
				}

				continue
			}

			i++
		}
	}

	fmt.Println(times_splitted)

	// part 2
	cached_function_results := make(map[[2]int]int)
	fmt.Println(dfs(manifold, 0, start, cached_function_results))
}

func dfs(manifold []string, row, col int, cached_function_results map[[2]int]int) int {

	if row < 0 || row >= len(manifold) || col < 0 || col >= len(manifold[row]) {
		return 0
	}

	function_params := [2]int{row, col}
	if cached_function_results[function_params] != 0 {
		return cached_function_results[function_params]
	}

	if row == len(manifold) - 1 {
		cached_function_results[function_params] = 1
		return 1
	}

	cell := manifold[row][col]
	var function_result int
	if cell == '^' {
		function_result = dfs(manifold, row + 1, col - 1, cached_function_results) + dfs(manifold, row + 1, col + 1, cached_function_results)
	} else {
		function_result = dfs(manifold, row + 1, col, cached_function_results)
	}

	cached_function_results[function_params] = function_result
	return function_result
}
