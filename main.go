package main

import (
	"fmt"
	"log"
	"os"

	"lineare/helpers"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}
	sumx, sumy, sumxy, sumx2, sumy2, length, err := helpers.Read_File()
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	m, b := helpers.Linear_Regression_Line(sumx, sumy, sumxy, sumx2, sumy2, length)
	r := helpers.Correlation_Coefficient(sumx, sumy, sumxy, sumx2, sumy2, (length))
	fmt.Printf("Linear Regression Line: y =%.6fx + %.6f \n", m, b)
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r)
}
