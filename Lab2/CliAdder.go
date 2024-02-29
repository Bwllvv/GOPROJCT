package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println(adder())
}

func adder() int {
	num1 := flag.Int("lev1", 0, "Number1")
	num2 := flag.Int("lev2", 0, "Number2")
	flag.Parse()
	intNum1 := *num1
	intNum2 := *num2
	if intNum1 == 0 && intNum2 == 0 {
		fmt.Println("\nnm 1:")
		fmt.Scanf("%d", &intNum1)
		fmt.Println("nm2 2:")
		fmt.Scanf("%d", &intNum2)
	}
	result := intNum1 + intNum2
	return result
}
