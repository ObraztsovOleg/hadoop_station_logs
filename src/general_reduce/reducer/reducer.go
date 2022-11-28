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

	keys := make([]int, 0, len(ram_error))
	for k := range ram_error {
		enum, err := strconv.Atoi(k)
		if err != nil {
			continue
		}
		keys = append(keys, enum)
	}
	sort.Sort(sort.IntSlice(keys))

	var prev_error int64 = 0

	for _, key := range keys {
		s := strconv.Itoa(key)
		avg_error = avg_error + float64(ram_error[s])
		avg_speed = avg_speed + ram_speed[s]

		var dirivative = int64(ram_error[s]) - prev_error

		if dirivative > 0 {
			avg_rerror = avg_rerror + float64(ram_error[s])
			avg_rspeed = avg_rspeed + ram_speed[s]
			rcount = rcount + 1
		}
		prev_error = int64(ram_error[s])
	}

	prev_error = 0

	// fmt.Println(rcount, avg_rspeed, avg_rerror)
	avg_error = avg_error / float64(len(ram_error))
	avg_speed = avg_speed / float64(len(ram_error))
	avg_rerror = avg_rerror / float64(rcount)
	avg_rspeed = avg_rspeed / float64(rcount)

	for _, key := range keys {
		s := strconv.Itoa(key)
		sum_mult = sum_mult + (float64(ram_error[s])-avg_error)*(ram_speed[s]-avg_speed)
		speed_sum_sq = speed_sum_sq + (ram_speed[s]-avg_speed)*(ram_speed[s]-avg_speed)
		error_sum_sq = error_sum_sq + (float64(ram_error[s])-avg_error)*(float64(ram_error[s])-avg_error)

		var dirivative = int64(ram_error[s]) - prev_error

		if dirivative > 0 {
			error_sum = error_sum + float64(ram_error[s])
			weighted_speed = weighted_speed + float64(ram_error[s])*ram_speed[s]
			rspeed_sum_sq = rspeed_sum_sq + (ram_speed[s]-avg_rspeed)*(ram_speed[s]-avg_rspeed)
		}
		prev_error = int64(ram_error[s])
	}

	fmt.Println(error_sum)
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
