package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()
		rows_fields := strings.Split(row, ",")
		for index, elem := range rows_fields {
			rows_fields[index] = strings.Trim(elem, "\t ")
		}

		if len(rows_fields) == 3 {
			day, err := strconv.ParseInt(rows_fields[0], 10, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}

			sec, err := strconv.ParseInt(rows_fields[1], 10, 64)
			if err != nil {
				fmt.Println(err)
				continue
			}

			speed, err := strconv.ParseFloat(rows_fields[2], 64)
			if err != nil {
				fmt.Println(err)
				continue
			}

			time := (day-1)*86400 + sec

			fmt.Printf("%v\t%v\t%.7f\n", day, time, speed)
		} else {
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	}
}
