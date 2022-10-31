package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var ram_error = make(map[string]int)
var ram_speed = make(map[string]float64)

var (
	avg_error      = 0.0
	avg_speed      = 0.0
	sum_mult       = 0.0
	speed_sum_sq   = 0.0
	error_sum_sq   = 0.0
	error_sum      = 0.0
	weighted_speed = 0.0
)

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
		avg_error = avg_error + float64(ram_error[key])
		avg_speed = avg_speed + ram_speed[key]
	}

	avg_error = avg_error / float64(len(ram_error))
	avg_speed = avg_speed / float64(len(ram_error))

	for key := range ram_error {
		speed_sum_sq = speed_sum_sq + (ram_speed[key]-avg_speed)*(ram_speed[key]-avg_speed)
		error_sum_sq = error_sum_sq + (float64(ram_error[key])-avg_error)*(float64(ram_error[key])-avg_error)
		sum_mult = sum_mult + (float64(ram_error[key])-avg_error)*(ram_speed[key]-avg_speed)
		error_sum = error_sum + float64(ram_error[key])
		weighted_speed = weighted_speed + float64(ram_error[key])*ram_speed[key]
	}

	fmt.Printf("correlation cooficient: %f\tweighted average: %f\n",
		sum_mult/math.Pow(error_sum_sq*speed_sum_sq, 1.0/2),
		weighted_speed/error_sum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
