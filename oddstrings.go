package main

import (
	"errors"
	"fmt"
)

func main() {
	var size int

	fmt.Println("enter slice size")
	fmt.Scan(&size)
	data := make([]int64, size, size)
	fmt.Println("enter slice elements")

	for i := 0; i < size; i++ {
		fmt.Scan(&data[i])
	}
	fmt.Println("enter batch size")
	fmt.Scan(&size)

	res, err := SplitSlice(data, size)
	if err != nil {
		fmt.Println("could not split")
		return
	}

	fmt.Printf("%v ", res)
}

func SplitSlice(numbers []int64, batchSize int) ([][]int64, error) {
	var begin, end int

	if batchSize < 0 || numbers == nil {
		return nil, errors.New("invalid batch size or empty slice")
	}

	quantity := len(numbers) / int(batchSize)
	if len(numbers)%int(batchSize) != 0 {
		quantity += 1
	}
	end = batchSize
	batches := make([][]int64, quantity)

	for i := 0; i < quantity; i++ {
		if end > len(numbers) {
			batches[i] = numbers[begin:]
			break
		}
		batches[i] = numbers[begin:end]
		begin = end
		end += int(batchSize)
	}
	return batches, nil
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	numMap := make(map[int]int, len(nums))
	var difference int

	for i, elem := range nums {
		if (target > elem && elem >= 0) || (target < 0) || elem < 0 && target > 0 || elem > target || elem < 0 && target == 0 {
			difference = target - elem
		} else {
			difference = elem - target
		}
		if i < len(nums) {
			if num, inMap := numMap[target-difference]; inMap {
				res[0], res[1] = num, i
				return res
			} else {
				numMap[difference] = i
			}
		}
	}
	return res
}
