package main

import (
	"errors"
	"fmt"
)

func DeleteSlice[T any](slice []T, idx int) ([]T, error) {
	if idx < 0 || idx > len(slice) {
		return nil, errors.New("下标错误")
	}
	s1 := make([]T, 0, len(slice)-1)
	for i := 0; i < idx; i++ {
		s1 = append(s1, slice[i])
	}
	for i := idx + 1; i < len(slice); i++ {
		s1 = append(s1, slice[i])
	}
	return s1, nil
}

func main() {
	//fmt.Printf("Slice after deletion: %v\n", DeleteSlice([]int{1, 3, 5, 7, 9}, 2))
	val, _ := DeleteSlice[int]([]int{1, 3, 5, 7}, 2)
	fmt.Printf("%v \n", val)
}
