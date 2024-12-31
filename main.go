package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}
	sumx, sumy, sumxy, sumx2, sumy2, length, err := Read_File()
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	m, b := Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2, length)
	r := Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2, (length))
	fmt.Printf("Linear Regression Line: y =%.6fx + %.6f \n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}

func Read_File() (float64, float64, float64, float64, float64, int, error) {
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
	length := 0
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
	length = int(x)
	return float64(sumx), sumy, sumxy, sumx2, sumy2, length, nil
}

func Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2 float64, legth int) (float64, float64) {
	var m float64
	var b float64

	m = ((float64(legth) * sumxy) - (sumx * sumy)) / ((float64(legth) * sumx2) - (sumx * sumx))

	b = ((sumy) - (m * sumx)) / float64(legth)
	return m, b
}

func Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2 float64, legth int) float64 {
	var r float64
	holder1 := (float64(legth) * sumxy) - (sumx * sumy)
	holder2 := ((float64(legth) * sumx2) - (sumx * sumx)) * ((float64(legth) * sumy2) - (sumy * sumy))

	r = holder1 / math.Sqrt(holder2)
	return r
}
