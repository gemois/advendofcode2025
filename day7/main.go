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
}
