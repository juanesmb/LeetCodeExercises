package twosum

import "fmt"

func Init() {
	input := []int{2, 7, 11, 15}
	result := twoSum(input, 9)

	fmt.Sprintln(result)
}

func twoSum(nums []int, target int) []int {
	fmt.Sprintln("nums: ", nums)
	fmt.Sprintln("target: ", target)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result := []int{i, j}
				fmt.Sprintln("result: ", result)

				return result
			}
		}
	}

	fmt.Sprintln("No result was found")

	return nil
}
