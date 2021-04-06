package bubble

// bubbleSort
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j+1], nums[j] = nums[j], nums[j+1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
