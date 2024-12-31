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
	sumx, sumy, sumxy, sumx2, sumy2,length, err := Read_File() 
	if err != nil {
		log.Fatal("Error reading file:", err)

	}
	fmt.Print(sumx)
	a , b := Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2,float64(length))
	r := Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2,float64(length))
	fmt.Printf("Linear Regression Line: y =%.6f + %.6f x\n",a,b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n",r)

}
func Read_File() ( float64, float64, float64,  float64,  float64,int,  error) {
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil{
		return 0,0,0,0,0,0,err
	}
	length := 0
	var sumx float64
	var sumy float64
	var sumxy float64
	var sumx2 float64
	var sumy2 float64
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		length++
		line := scanner.Text()
		number, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error Parsing :", err)
			continue
		}

		sumy += number
		sumxy += number * float64(sumx)
		sumx2 += float64(sumx) * float64(sumx)
		sumy2 += number * number
		//fmt.Println(number)
		sumx++


	}
	if err := scanner.Err(); err != nil {
		return 0, 0, 0, 0, 0,0, err
	}
	return float64(sumx), sumy, sumxy, sumx2, sumy2,length, nil
}

func Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2,length float64) (float64, float64) {
	var a float64
	var b float64


	b = ((length * sumxy) - (sumx * sumy)) / (length * sumx2) - (sumx * sumx)
	a = (sumy / length) - ((b * sumx) / length)
	return b, a
}

func Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2,length float64) float64 {
	
	var r float64
	holder1 := (length * sumxy) - (sumx * sumy)
	holder2 := ((length * sumx2) - (sumx * sumx)) * ((length * sumy2) - (sumy * sumy))
	r = holder1 / math.Sqrt(holder2)
	return r
}
