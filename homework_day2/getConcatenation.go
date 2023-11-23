package homework_day2

func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i, num_i := range nums {
		ans[i] = num_i
		ans[i+len(nums)] = nums[i]
	}
	return ans
}
