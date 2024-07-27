/**
* シェルソート
* 挿入ソートを改良したもの
* 最悪計算時間: O(n^2)
* 最良計算時間: O(n)
* 平均計算時間: O(n^2)
* 最悪空間計算時間: O(1)
*/
package main

import "fmt"

func shellSort(nums []int) []int {
	// リストの長さ
	lenNums := len(nums)
	// ギャップ
	gap := lenNums / 2

	// ギャップが0になるまでループ
	for gap > 0 {
		// ギャップ分だけループ
		for i := gap; i < lenNums; i++ {
			// 未ソートの要素を一時保存
			temp := nums[i]
			// ソート済みの部分の最後の要素のインデックス
			j := i
			// ソート済みの部分の最後の要素が未ソートの要素より大きい場合
			for j >= gap && nums[j-gap] > temp {
				// ソート済みの部分の最後の要素を一つ後ろに移動
				nums[j] = nums[j-gap]
				// ソート済みの部分の最後の要素のインデックスを一つ前に移動
				j -= gap
			}
			// 未ソートの要素をソート済みの部分に挿入
			nums[j] = temp
		}
		// ギャップを半分にする
		gap /= 2
	}
	return nums
}

func main() {
	// ソートするリスト
	nums := []int{5, 3, 8, 4, 2}
	// シェルソートを実行
	fmt.Println(shellSort(nums))
}
