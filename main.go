package main

import (
	"fmt"

	"github.com/WindomZ/interesting-solutions/solutions"
)

func main() {
	matrixStr := "HTQZAFEOHPOMLMBCSOLQHXDKO"
	matrixX, matrixY := 5, 5 // 5*5 matrix

	// Solution1
	fmt.Println("========== Solution1 ==========")
	result1 := solutions.Solution1([]string{"HELLO"}, matrixStr, matrixX, matrixY)
	for word, letters := range result1 {
		fmt.Println("Word:", word)
		for _, letter := range letters {
			fmt.Println(word, "matched letters:", letter.ReadableString())
		}
	}

	// Solution2
	fmt.Println("========== Solution2 ==========")
	result2 := solutions.Solution2([]string{"MOMOSO"}, matrixStr, matrixX, matrixY)
	for word, letters := range result2 {
		for _, letter := range letters {
			fmt.Println(word, "matched letters:", letter.ReadableString())
		}
	}

	// Solution3
	fmt.Println("========== Solution3 ==========")
	result3 := solutions.Solution3([]string{"MOMOSO"}, matrixStr, matrixX, matrixY)
	for word, letters := range result3 {
		for _, letter := range letters {
			fmt.Println(word, "matched letters:", letter.ReadableString())
		}
	}
}
