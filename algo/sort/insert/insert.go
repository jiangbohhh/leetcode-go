package insert

// insertSort
func insertSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		val := nums[i]
		j := i - 1
		for ; j >= 0; j-- {
			if nums[j] > val {
				nums[j+1] = nums[j]
			}else {
				break
			}
		}
		nums[j+1] = val
	}
}
