package main

func MaxDifference(numbers []int) int {
	if len(numbers) < 2 {
		return 0
	}
	maxNum := numbers[0]
	minNum := numbers[0]

	for _, x := range numbers {
		if maxNum < x {
			maxNum = x
		}
		if minNum > x {
			minNum = x
		}
	}
	return maxNum - minNum

}
