/**
* マージソート
* 分割統治法を用いて、配列を分割していき、最終的にはソートされた配列を得る。
* 分割した配列を統合することでソートを行う。
*/
package main

import "fmt"

// マージソートを行う関数
func mergeSort(nums []int) []int {
	// 配列の長さを取得
	var lenNums = len(nums)

	// 配列の長さが1の場合はそのまま返す
	if lenNums == 1 {
		return nums
	}

	// 配列を分割するための中央値を取得
	mid := lenNums / 2
	var (
		// 分割した配列を格納するための変数
		left = make([]int, mid)
		// 分割した配列を格納するための変数
		right = make([]int, lenNums-mid)
	)

	// 配列を分割
	for i := 0; i < lenNums; i++ {
		// 分割した配列を格納するための変数に値を格納
		if i < mid {
			left[i] = nums[i]
		} else {
			right[i-mid] = nums[i]
		}
	}

	// 分割した配列を再帰的にソート
	return merge(mergeSort(left), mergeSort(right))
}

// 分割した配列を統合するための関数
func merge(left, right []int) (result []int) {
	// 分割した配列を統合するための変数を作成
	result = make([]int, len(left)+len(right))

	// 分割した配列を統合するための変数を作成
	i := 0
	// 分割した配列を統合するための変数を作成
	for len(left) > 0 && len(right) > 0 {
		// 分割した配列を統合するための変数に値を格納
		if left[0] < right[0] {
			// 分割した配列を統合するための変数に値を格納
			result[i] = left[0]
			// 分割した配列を統合するための変数に値を格納
			left = left[1:]
		} else {
			// 分割した配列を統合するための変数に値を格納
			result[i] = right[0]
			// 分割した配列を統合するための変数に値を格納
			right = right[1:]
		}
		// 分割した配列を統合するための変数に値を格納
		i++
	}

	// 分割した配列を統合するための変数に値を格納
	for j := 0; j < len(left); j++ {
		// 分割した配列を統合するための変数に値を格納
		result[i] = left[j]
		i++
	}

	// 分割した配列を統合するための変数に値を格納
	for j := 0; j < len(right); j++ {
		// 分割した配列を統合するための変数に値を格納
		result[i] = right[j]
		i++
	}

	return
}

// メイン関数
func main() {
	// 配列を作成
	nums := []int{3, 5, 8, 10, 17, 11, 13, 19, 2, 4, 7}
	// マージソートを行う
	fmt.Println(mergeSort(nums))
}
