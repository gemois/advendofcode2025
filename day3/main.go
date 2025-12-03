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
	
		joltage_sum := 0;

    for scanner.Scan() {
			line := scanner.Text()

			batteries := []int{}
			for _, char := range line {
				batteries = append(batteries, int(char - '0'))
			}

			max_battery := batteries[0]
			max_index := 0
			for i, battery := range batteries {
				if battery> max_battery {
					max_battery = battery
					max_index = i
				}
			}

			if max_index == len(batteries) - 1 {
				left_max_battery := batteries[0]
				for i := 0; i < len(batteries) - 1; i++ {
					if batteries[i] > left_max_battery {
						left_max_battery = batteries[i]
					}
				}

				battery_joltage,_ := strconv.Atoi(strconv.Itoa(left_max_battery) + strconv.Itoa(max_battery))
				joltage_sum += battery_joltage
				continue
			}

			batteries_at_the_right := batteries[max_index+ 1:]
			right_max_battery := batteries_at_the_right[0]
			for _, battery := range batteries_at_the_right {
					if battery > right_max_battery {
						right_max_battery = battery
					}
			}

			battery_joltage,_ := strconv.Atoi(strconv.Itoa(max_battery) + strconv.Itoa(right_max_battery))
			joltage_sum += battery_joltage
		}

		fmt.Println(joltage_sum)
}
