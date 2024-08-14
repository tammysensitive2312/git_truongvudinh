package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func determinant(matrix [3][3]float64) float64 {
	// store index of the matrix
	indices := [6][3]int{
		{0, 1, 2},
		{1, 2, 0},
		{2, 0, 1},
		{2, 1, 0},
		{1, 0, 2},
		{0, 2, 1},
	}

	// variable to store value of matrix determinant
	det := 0.0
	for i := 0; i < 3; i++ {
		// Calculate positive terms
		product := 1.0
		for j := 0; j < 3; j++ {
			product *= matrix[j][indices[i][j]]
		}
		det += product

		// Calculate negative terms
		product = 1.0
		for j := 0; j < 3; j++ {
			product *= matrix[j][indices[i+3][j]]
		}
		det -= product
	}

	return det
}

type Words struct {
	exist, length int
}

func removeSpecChar(word string) string {
	var builder strings.Builder
	for _, char := range word {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			builder.WriteRune(char)
		}
	}
	return builder.String()
}

func solve(filename string) {
	// read file and handle errors
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist: ", filename)
		} else {
			fmt.Println("Error reading file: ", err)
		}
	}

	// split data into a slice of string
	words := strings.Fields(string(data))
	fmt.Println(len(words))

	// initialize a map
	wordCount := make(map[string]Words)

	// Lọc bỏ dấu câu và đếm tần suất từ
	for _, word := range words {
		newWord := removeSpecChar(word)
		if newWord != "" {
			// Lấy giá trị hiện tại của từ từ map
			wordInfo, exists := wordCount[newWord]
			if exists {
				// Nếu từ đã tồn tại, tăng số lần xuất hiện
				wordInfo.exist++
			} else {
				// Nếu từ chưa tồn tại, khởi tạo giá trị mới với độ dài từ
				wordInfo = Words{exist: 1, length: len(newWord)}
			}
			// Cập nhật map với giá trị mới
			wordCount[newWord] = wordInfo
		}
	}

	// In kết quả đếm từ
	for k, v := range wordCount {
		fmt.Printf("%s: %d : %d\n", k, v.exist, v.length)
	}
}

func Fib(n int) []int {
	var ans = []int{}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ans = append(ans, a)
		a, b = b, a+b
	}
	return ans
}

func main() {
	/**
	matrix := [3][3]float64{
		{1.9, 2, 3},
		{4, 2, 6},
		{7, 8.8, 4},
	}

	det := determinant(matrix)
	fmt.Printf("%f", det)
	*/

	//solve("data.txt")

	fmt.Println(Fib(6))
}
