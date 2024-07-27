/**
* インサーションソート
* 未ソートの部分から要素を取り出して、ソート済みの部分に適切な位置に挿入する
* 最悪計算時間: O(n^2)
* 最良計算時間: O(n)
* 平均計算時間: O(n^2)
* 最悪空間計算時間: O(1)
*/
package main

import "fmt"

func insersionSort(nums []int) []int {
	// リストの長さ
	lenNums := len(nums)

	// リストの長さ-1回ループ
	for i := 1; i < lenNums; i++ {
		// 未ソートの要素を一時保存
		temp := nums[i]
		// ソート済みの部分の最後の要素のインデックス
		j := i - 1
		// ソート済みの部分の最後の要素が未ソートの要素より大きい場合
		for j >= 0 && nums[j] > temp {
			// ソート済みの部分の最後の要素を一つ後ろに移動
			nums[j+1] = nums[j]
			// ソート済みの部分の最後の要素のインデックスを一つ前に移動
			j -= 1
		}
		// 未ソートの要素をソート済みの部分に挿入
		nums[j+1] = temp
	}
	return nums
}

func main() {
	// リスト
	nums := []int{5, 3, 8, 4, 2}
	// インサーションソート
	fmt.Println(insersionSort(nums))
}
