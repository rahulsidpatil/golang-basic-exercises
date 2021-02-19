package main

func main() {

	nums := []int{1, 2, 3}
	Show("nums", nums)

	_ = append(nums, 4)
	Show("nums", nums)

	nums = append(nums, 4)
	Show("nums", nums)

	nums = append(nums, 9)
	Show("nums", nums)

	nums = append(nums, 4)
	Show("nums", nums)

	// or:
	// nums = append(nums, 4, 9)
	// Show("nums", nums)

	nums = []int{1, 2, 3}
	tens := []int{12, 13}

	nums = append(nums, tens...)
	Show("nums", nums)
}
