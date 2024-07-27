/**
* バブルソート
* 隣り合う要素を比較して、大小関係が逆であれば入れ替える
* 1回のスキャンで、最後の要素がソートされる
* 最悪計算時間: O(n^2)
* 最良計算時間: O(n)
* 平均計算時間: O(n^2)
* 最悪空間計算時間: O(1)
*/
package main

import "fmt"

func bubbleSort(nums []int) []int {
	// リストの長さ
	lenNums := len(nums)

	// リストの長さ-1回ループ
	for i := 0; i < lenNums-1; i++ {
		// リストの長さ-1-i回ループ
		for j := 0; j < lenNums-1-i; j++ {
			// 隣り合う要素を比較して、大小関係が逆であれば入れ替える
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func main() {
	// リスト
	nums := []int{5, 3, 8, 4, 2}
	// バブルソート出力
	fmt.Println(bubbleSort(nums))
}
