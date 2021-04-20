package merge

// mergeSort
func mergeSort(m []int) {
	if len(m) < 2 {
		return
	}
	mid := len(m) / 2
	mergeSort(m[:mid])
	mergeSort(m[mid:])
	merge(m, mid)
}

func merge(nums []int, mid int){
	i,j := 0, mid
	tmp := make([]int, 0, len(nums))
	for i < mid && j < len(nums) {
		if i < mid && nums[i] < nums[mid] {
			tmp = append(tmp, nums[i])
			i++
		}
		if j < len(nums) && nums[j] < nums[i] {
			tmp = append(tmp, nums[j])
			j++
		}
	}
	if i < mid {
		tmp = append(tmp, nums[i:mid]...)
	}
	if j < len(nums) {
		tmp = append(tmp, nums[j:]...)
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
	return
}