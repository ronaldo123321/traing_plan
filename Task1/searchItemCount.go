package main

import "fmt"

func searchItem(nums []int) {
	countMap := make(map[int]int)
	for _, item := range nums {
		countMap[item]++
	}
	for num, count := range countMap {
		if count == 1 {
			fmt.Println(num)
			return
		}
	}
}

func main() {
	nums := []int{3, 3, 4, 2, 4, 2, 1}
	searchItem(nums)
}
