package main

import (
	"errors"
	"fmt"
)

func main() {
	slice := []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	res, err := SplitSlice(slice, 2)
	if err != nil {
		panic("could not split")
	}
	fmt.Printf("%v ", res)
}

func SplitSlice(numbers []int64, batchsize int) ([][]int64, error) {
	var begin, end int
	if batchsize < 0 || numbers == nil {
		return nil, errors.New("invalid batch size or empty slice")
	}
	quantity := len(numbers) / int(batchsize)
	if len(numbers)%int(batchsize) != 0 {
		quantity += 1
	}
	end = batchsize
	batches := make([][]int64, quantity)
	for i := 0; i < quantity; i++ {
		if end > len(numbers) {
			batches[i] = numbers[begin:]
			break
		}
		batches[i] = numbers[begin:end]
		begin = end
		end += int(batchsize)
	}
	return batches, nil
}

/* groupCity := map[int][]string{
	10:   []string{"Деревня", "Село"},        // города с населением 10-99 тыс. человек
	100:  []string{"Город", "Большой город"}, // города с населением 100-999 тыс. человек
	1000: []string{"Миллионик"},              // города с населением 1000 тыс. человек и более
}
cityPopulation := map[string]int{
	"Село":      100,
	"Миллионик": 500,
}
for _, val := range groupCity[10] {
	if _, found := cityPopulation[val]; found {
		delete(cityPopulation, val)
	}
}
for _, val := range groupCity[1000] {
	if _, found := cityPopulation[val]; found {
		delete(cityPopulation, val)
	}
} */
