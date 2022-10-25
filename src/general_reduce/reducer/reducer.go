package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ram_error = make(map[string]int)
var ram_speed = make(map[string]float64)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)

		if len(fields) == 3 {
			if fields[0] == "error" {
				error_num, err := strconv.ParseInt(fields[2], 10, 64)
				if err != nil {
					continue
				}

				ram_error[fields[1]] = int(error_num)
			} else if fields[0] == "h" {
				h_speed, err := strconv.ParseFloat(fields[2], 64)
				if err != nil {
					continue
				}

				ram_speed[fields[1]] = h_speed
			} else {
				continue
			}

		} else {
			continue
		}
	}

	for key := range ram_error {
		fmt.Printf("%s :: %v - %.7f\n", key, ram_error[key], ram_speed[key])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
