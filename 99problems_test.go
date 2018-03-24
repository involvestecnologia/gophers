package main

import "testing"

func TestLastItemOfAStringSlice(t *testing.T) {
	someSlice := []string{"a", "b", "c", "d", "e", "f"}
	value := findlastSliceItem(someSlice)

	if value != someSlice[len(someSlice)-1] {
		t.Errorf("Opa, não rolou eu estava esperando %s mas veio %s", someSlice[len(someSlice)-1], value)
	}
}

func TestLastItemButOneOfAStringSlice(t *testing.T) {
	someSlice := []string{"a", "b", "c", "d", "e", "f"}
	value := findlastButOneSliceItem(someSlice)

	if value != someSlice[len(someSlice)-2] {
		t.Errorf("Opa, não rolou eu estava esperando %s mas veio %s", someSlice[len(someSlice)-2], value)
	}
}

func TestNPositionFinder(t *testing.T) {
	someIntSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	value := findNElement(0, someIntSlice)

	if value != 1 {
		t.Errorf("Opa não rolou, eu estava esperando %d, mas veio %d", someIntSlice[0], value)
	}

}
