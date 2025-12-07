package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open(os.Args[1])
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]string
	for scanner.Scan() {
		line := scanner.Text()
		line_locations := strings.Split(line, "")
		grid = append(grid, line_locations)
	}

	fmt.Println(grid)

	can_be_accesed := 0
	paper_roll := "@"

	rows_lenght := len(grid)
	columns_lenght := len(grid[0])

	for row := 0; row < rows_lenght; row++ {
		for column := 0; column < columns_lenght; column++ {
			if grid[row][column] == "." {
				continue
			}

			adjacent_forklifts_count := 0

			if column-1 >= 0 && grid[row][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if column+1 < columns_lenght && grid[row][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && grid[row-1][column] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && grid[row+1][column] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && column-1 >= 0 && grid[row-1][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && column+1 < columns_lenght && grid[row-1][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && column-1 >= 0 && grid[row+1][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && column+1 < columns_lenght && grid[row+1][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if adjacent_forklifts_count < 4 {
				can_be_accesed++
			}
		}
	}

	fmt.Println(can_be_accesed)

	// part 2
	fmt.Println(find_all_accesible_paper_rolls(grid))
}

func find_all_accesible_paper_rolls(grid [][]string) int {
	can_be_accesed := 0
	paper_roll := "@"

	rows_lenght := len(grid)
	columns_lenght := len(grid[0])

	grid_removed := make([][]string, len(grid))

	for i := range grid {
		grid_removed[i] = make([]string, len(grid[i]))
		copy(grid_removed[i], grid[i])
	}

	for row := 0; row < rows_lenght; row++ {
		for column := 0; column < columns_lenght; column++ {
			if grid[row][column] == "." {
				continue
			}

			adjacent_forklifts_count := 0

			if column-1 >= 0 && grid[row][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if column+1 < columns_lenght && grid[row][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && grid[row-1][column] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && grid[row+1][column] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && column-1 >= 0 && grid[row-1][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row-1 >= 0 && column+1 < columns_lenght && grid[row-1][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && column-1 >= 0 && grid[row+1][column-1] == paper_roll {
				adjacent_forklifts_count++
			}

			if row+1 < rows_lenght && column+1 < columns_lenght && grid[row+1][column+1] == paper_roll {
				adjacent_forklifts_count++
			}

			if adjacent_forklifts_count < 4 {
				can_be_accesed++
				grid_removed[row][column] = "."
			}
		}
	}

	if can_be_accesed == 0 {
		return 0
	}

	return can_be_accesed + find_all_accesible_paper_rolls(grid_removed)
}
