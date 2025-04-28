package main

import "fmt"

func searchAllList(nums []int) {

	var backtrack func(int)
	n := len(nums)
	backtrack = func(first int) {

		if first == n {
			fmt.Println(append([]int{}, nums...))
			return
		}

		for i := first; i < n; i++ {
			nums[first], nums[i] = nums[i], nums[first]
			backtrack(first + 1)
			nums[first], nums[i] = nums[i], nums[first]
		}

	}
	backtrack(0)

}

func main() {
	nums := []int{2, 3, 4}
	searchAllList(nums)
}
