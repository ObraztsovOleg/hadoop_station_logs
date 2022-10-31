package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ram = make(map[string]float64)

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
		} else {
			continue
		}
	}

	for el := range ram {
		fmt.Printf("h\t%s\t%f\n", el, ram[el])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
