package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	time  = int64(0)
	speed = float64(0)
	count = int64(0)
)

var ram = make(map[string]float64)
var ram_count = make(map[string]int)
var max = 0.0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)

		if len(fields) == 3 {
			stdin_speed, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				continue
			}

			ram[fields[1]] = ram[fields[1]] + stdin_speed
			ram_count[fields[1]] = ram_count[fields[1]] + 1
		} else {
			continue
		}
	}

	for el := range ram {
		avg_speed := ram[el] / float64(ram_count[el])

		if max < avg_speed {
			max = avg_speed
		}
	}

	fmt.Printf("Max avg speed  - %f\n", max)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
