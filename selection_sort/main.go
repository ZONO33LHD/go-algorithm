/**
* セレクションソート
* 未ソートの部分から最小値を選んで、ソート済みの部分に追加する
* 最悪計算時間: O(n^2)
* 最良計算時間: O(n^2)
* 平均計算時間: O(n^2)
* 最悪空間計算時間: O(1)
*/
package main

import "fmt"

func selectionSort(nums []int) []int {
	// リストの長さ
	lenNums := len(nums)

	// リストの長さ-1回ループ
	for i := 0; i < lenNums; i++ {
		// 最小値のインデックス
		minIdx := i
		// リストの長さ-1-i回ループ
		for j := i + 1; j < lenNums; j++ {
			// 最小値のインデックスがリストの長さ-1-iより大きい場合
			if nums[minIdx] > nums[j] {
				minIdx = j
			}
		}
		// 最小値のインデックスとリストの長さ-1-iを入れ替える
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

func main() {
	// リスト
	nums := []int{5, 3, 8, 4, 2}
	// セレクションソート
	fmt.Println(selectionSort(nums))
}
