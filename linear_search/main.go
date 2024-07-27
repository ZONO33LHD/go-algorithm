// 線形探索：5を探す
package main

import "fmt"

func linearSearch(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 検索する値
	target := 5
	// 線形探索
	index := linearSearch(nums, target)
	// targetのインデックス：4を出力想定
	fmt.Printf("targetのインデックス: %d\n", index)
}
