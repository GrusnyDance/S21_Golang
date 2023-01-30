package main

import (
	"bufio"
	"fmt"
	"github.com/montanaflynn/stats"
	"os"
	"strconv"
	"strings"
)

func main() {
	m := make(map[string]func([]int), 4)
	m["mean"] = mean
	m["median"] = median
	m["mode"] = mode
	m["sd"] = sd

	scanner := bufio.NewScanner(os.Stdin) // Creating scanner object
	for scanner.Scan() {
		var slice []int
		strSlice := strings.Fields(scanner.Text())
		for _, elem := range strSlice {
			if val, err := strconv.Atoi(elem); err == nil {
				slice = append(slice, val)
			} else {
				fmt.Println("Your data is spoiled, bye")
				return
			}
		}

		if len(os.Args) == 1 {
			mean(slice)
			median(slice)
			mode(slice)
			sd(slice)
		} else {
			for _, elem := range os.Args[1:] {
				if f, ok := m[elem]; ok {
					f(slice)
				} else {
					fmt.Println("Your args are wrong, bye")
					return
				}
			}
		}

	}
}

func mean(slice []int) {
	data := stats.LoadRawData(slice)
	ret, err := stats.Mean(data)
	if err != nil {
		fmt.Println("Smth is wrong with your data")
	} else {
		fmt.Printf("%.2f\n", ret)
	}
}

func median(slice []int) {
	data := stats.LoadRawData(slice)
	ret, err := stats.Median(data)
	if err != nil {
		fmt.Println("Smth is wrong with your data")
	} else {
		fmt.Printf("%.2f\n", ret)
	}
}

func mode(slice []int) {
	data := stats.LoadRawData(slice)
	ret, err := stats.Mode(data)
	if err != nil {
		fmt.Println("Smth is wrong with your data")
	} else {
		if len(ret) == 0 {
			min, _ := stats.Min(data)
			fmt.Printf("%.2f\n", min)
		} else if len(ret) > 1 {
			min, _ := stats.Min(ret)
			fmt.Printf("%.2f\n", min)
		} else {
			fmt.Printf("%.2f\n", ret[0])
		}
	}
}

func sd(slice []int) {
	data := stats.LoadRawData(slice)
	ret, err := stats.StandardDeviation(data)
	if err != nil {
		fmt.Println("Smth is wrong with your data")
	} else {
		fmt.Printf("%.2f\n", ret)
	}
}
