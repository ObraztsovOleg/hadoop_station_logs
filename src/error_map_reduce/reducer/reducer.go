package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ram = make(map[string]int)

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
		fmt.Printf("error\t%s\t%v\n", key, ram[key])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
