/**
* クイックソート
* クイックソートは、分割統治法に基づいたアルゴリズムで、平均計算時間がO(n log n)である。
* クイックソートは、ピボットを選択し、ピボットより小さい要素と大きい要素に分割する。
* ピボットより小さい要素と大きい要素に分割する。
*/
package main

import "fmt"

func partition(nums []int, low, high int) int {
	// ピボットを選択
	pivot := nums[high]
	// ピボットより小さい要素と大きい要素に分割するためのインデックス
	i := low - 1
	// ピボットより小さい要素と大きい要素に分割する
	for j := low; j < high; j++ {
		// ピボットより小さい要素と大きい要素に分割する
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// ピボットを適切な位置に移動
	nums[i+1], nums[high] = nums[high], nums[i+1]
	// ピボットの位置を返す
	return i + 1
}

func quickSort(nums []int, low, high int) []int {
	// ピボットを適切な位置に移動
	if low < high {
		partitionIndex := partition(nums, low, high)
		quickSort(nums, low, partitionIndex-1)
		quickSort(nums, partitionIndex+1, high)
	}
	// ソートされた配列を返す
	return nums
}

func quickSortStart(nums []int) []int {
	// ソートされた配列を返す
	return quickSort(nums, 0, len(nums)-1)
}

func main() {
	// ソートされた配列を返す
	nums := []int{10, 7, 8, 9, 1, 5}
	// ソートされた配列を出力
	fmt.Println(quickSortStart(nums))
}
