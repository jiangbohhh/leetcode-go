package _select

// selectSort
func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				nums[j] = nums[i]
				nums[i] = min

			}
		}
	}
}
