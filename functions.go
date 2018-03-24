package main

import (
	"fmt"
)

func imprimeNome(algumnome string) {
	fmt.Println(algumnome)
}

func findlastSliceItem(arr []string) string {
	return arr[len(arr)-1]
}

func findlastButOneSliceItem(arr []string) string {
	return arr[len(arr)-2]
}

func findNElement(position int, arr []int) int {
	return arr[position]
}
