package main

import (
	"fmt"
	"sort"

	"github.com/montanaflynn/stats"
)

func main() {

	//// start with some source data to use
	//data := []float64{1.0, 2.1, 3.2, 4.823, 4.1, 5.8}
	//
	//// you could also use different types like this
	//// data := stats.LoadRawData([]int{1, 2, 3, 4, 5})
	//// data := stats.LoadRawData([]interface{}{1.1, "2", 3})
	//// etc...
	//
	//median, _ := stats.Median(data)
	//fmt.Println(median) // 3.65
	//
	//roundedMedian, _ := stats.Round(median, 0)
	//fmt.Println(roundedMedian) // 4

	data := []float64{5, 4, 4, 5, 5, 5, 4, 5, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 4, 4, 4, 4, 3, 5, 4, 5, 4, 4, 5, 4, 4, 100, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 2232, 4550}
	fmt.Println("Data UNORDERED: ", data)

	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	fmt.Println("Data ORDERED: ", data)

	percentile90, err := stats.Percentile(data, 90)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Percentile 90 (STATS Lib)", percentile90)

	percentile95, err := stats.Percentile(data, 95)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Percentile 95 (STATS Lib)", percentile95)

	fmt.Println("Percentile 90: ", data[int(0.90*float64(len(data)))])
	fmt.Println("Percentile 95: ", data[int(0.95*float64(len(data)))])

}
