package twosum

func twoSum(nums []int, target int) []int {
	for firstIdx, firstNum := range nums {
		for secondIdx, secondNum := range nums[firstIdx+1:] {
			if firstNum+secondNum == target {
				return []int{firstIdx, secondIdx + firstIdx + 1}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	inputMap := make(map[int]int)
	for idx, value := range nums {
		requireNum := target - value
		prevIdx, exist := inputMap[requireNum]
		if exist {
			return []int{prevIdx, idx}
		}

		inputMap[value] = idx
	}
	return []int{}
}
