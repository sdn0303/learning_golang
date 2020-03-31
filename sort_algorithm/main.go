package sort_algorithm

func BubbleSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func SelectionSort(numbers []int) []int {
	for idx, _ := range numbers {
		min := func(nums []int) (idx int) {
			idx = 0
			for i, v := range nums {
				if nums[0] > v {
					idx = i
				}
			}
			return
		}(numbers[idx:len(numbers)])
		numbers[idx], numbers[idx+min] = numbers[idx+min], numbers[idx]
	}
	return numbers
}

func InsertSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers)-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			} else {
				break
			}
		}
	}
	return numbers
}
