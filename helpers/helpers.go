package helpers

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Read_File() (float64, float64, float64, float64, float64, float64, error) {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	defer file.Close()
	var sumx float64 = 0
	var sumy float64 = 0
	var sumxy float64
	var sumx2 float64
	var sumy2 float64
	length := 0.0
	var x float64 = 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()
		y, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error Parsing :", err)
			continue
		}

		sumy += y
		sumxy += y * x
		sumx2 += x * x
		sumy2 += y * y
		sumx += x
		x++

	}
	if err := scanner.Err(); err != nil {
		return 0, 0, 0, 0, 0, 0, err
	}
	length = x
	return float64(sumx), sumy, sumxy, sumx2, sumy2, length, nil
}

func Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2 ,legth float64) (float64, float64) {
	var m float64
	var b float64

	m = ((legth * sumxy) - (sumx * sumy)) / ((legth * sumx2) - (sumx * sumx))

	b = ((sumy) - (m * sumx)) / legth
	return m, b
}

func Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2 ,legth float64) float64 {
	var r float64
	holder1 := (legth * sumxy) - (sumx * sumy)
	holder2 := ((legth * sumx2) - (sumx * sumx)) * ((legth * sumy2) - (sumy * sumy))

	r = holder1 / math.Sqrt(holder2)
	return r
}
