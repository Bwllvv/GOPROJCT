package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	showLines([]int{1, 2, 3})
}

func showLine(line int) string {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		return "File corrupted"
	}
	dataList := strings.Split(string(data), "\n")
	return dataList[line-1]
}

func showLines(lines []int) {
	for _, i := range lines {
		fmt.Printf("Line %d: %s", i, showLine(i))
		fmt.Println()
	}
}
