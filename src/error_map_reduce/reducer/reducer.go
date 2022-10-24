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

var ram = make(map[string]int)
var ram_count = make(map[string]int)
var max = 0

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()
		fields := strings.Fields(row)

		if len(fields) == 3 {
			err_num, err := strconv.ParseInt(fields[2], 10, 64)

			if err != nil {
				continue
			}

			ram[fields[1]] = ram[fields[1]] + int(err_num)
		} else {
			continue
		}
	}

	for key := range ram {
		fmt.Printf("%s - %v\n", key, ram[key])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
