package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var ram_error = make(map[string]int)
var increased_errors = make(map[string]int)
var ram_speed = make(map[string]float64)

const CC = 0.95

var (
	avg_error          float64 = 0.0
	avg_speed          float64 = 0.0
	avg_rerror         float64 = 0.0
	avg_rspeed         float64 = 0.0
	sum_mult                   = 0.0
	speed_sum_sq               = 0.0
	rspeed_sum_sq              = 0.0
	error_sum_sq               = 0.0
	error_sum                  = 0.0
	weighted_speed             = 0.0
	standart_deviation         = 0.0
	rcount             int     = 0
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

	keys := make([]string, 0, len(ram_error))
	for k := range ram_error {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var prev_error int64 = 0

	for _, key := range keys {
		// fmt.Println(key, ram_error[key], ram_speed[key])
		avg_error = avg_error + float64(ram_error[key])
		avg_speed = avg_speed + ram_speed[key]

		var dirivative float64 = float64(int64(ram_error[key])-prev_error) / float64(300)

		if dirivative > 0.0 {
			// fmt.Println(ram_error[key], prev_error)
			avg_rerror = avg_rerror + float64(ram_error[key])
			avg_rspeed = avg_rspeed + ram_speed[key]
			rcount = rcount + 1
		}
		prev_error = int64(ram_error[key])
	}

	// fmt.Println(rcount, avg_rspeed, avg_rerror)
	avg_error = avg_error / float64(len(ram_error))
	avg_speed = avg_speed / float64(len(ram_error))
	avg_rerror = avg_rerror / float64(rcount)
	avg_rspeed = avg_rspeed / float64(rcount)

	for _, key := range keys {
		sum_mult = sum_mult + (float64(ram_error[key])-avg_error)*(ram_speed[key]-avg_speed)
		speed_sum_sq = speed_sum_sq + (ram_speed[key]-avg_speed)*(ram_speed[key]-avg_speed)
		error_sum_sq = error_sum_sq + (float64(ram_error[key])-avg_error)*(float64(ram_error[key])-avg_error)

		var dirivative float64 = float64(int64(ram_error[key])-prev_error) / float64(300)

		if dirivative > 0.0 {
			error_sum = error_sum + float64(ram_error[key])
			weighted_speed = weighted_speed + float64(ram_error[key])*ram_speed[key]
			rspeed_sum_sq = rspeed_sum_sq + (ram_speed[key]-avg_rspeed)*(ram_speed[key]-avg_rspeed)
		}
		prev_error = int64(ram_error[key])
	}

	standart_deviation = math.Pow(rspeed_sum_sq/float64(rcount), 1.0/2)

	fmt.Printf("correlation cooficient: %f\nweighted average: %f\nconfidence interval:\n\t-min: %f\n\t-max: %f\n",
		sum_mult/math.Pow(error_sum_sq*speed_sum_sq, 1.0/2),
		weighted_speed/error_sum,
		(avg_rspeed - CC*standart_deviation/math.Pow(float64(rcount), 1.0/2)),
		(avg_rspeed + CC*standart_deviation/math.Pow(float64(rcount), 1.0/2)))

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
