/**
 * バイナリサーチ
 *
 * ソート済みのリストに対して、
 * 中央の値を取り出し、
 * その値と探している値を比較する。
 * その値が探している値より大きい場合、
 * 中央の値より前のリストを、
 * 小さい場合は中央の値より後ろのリストを
 * それぞれに対して再帰的にバイナリサーチを行う。
 *
 * 計算量はO(log n)
 */
package main

import "fmt"

// バイナリサーチ
func binarySearch(nums []int, value int) int {
        // 初期値
        low := 0
        // リストの最後のインデックス
	high := len(nums) - 1

        // リストの中央のインデックスを計算
	for low <= high {
		mid := (low + high) / 2

		// 中央の値が探している値と一致する場合
		if nums[mid] == value {
			return mid
		}
		// 中央の値が探している値より大きい場合
		if nums[mid] < value {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	// 見つからなかった場合
	return - 1
}

// 再帰を使用したバイナリサーチ
func binarySearchRecursive(nums []int, value int, low, high int) int {
	if low > high {
		return - 1
	}

	// リストの中央のインデックスを計算
	mid := (low + high) / 2

	// 中央の値が探している値と一致する場合
	if nums[mid] == value {
		return mid
	} else if nums[mid] < value {
		// 中央の値より後ろのリストを再帰的にバイナリサーチ
		return binarySearchRecursive(nums, value, mid + 1, high)
	} else {
		// 中央の値より前のリストを再帰的にバイナリサーチ
		return binarySearchRecursive(nums, value, low, mid - 1)
	}
}

func main() {
        // ソート済みのリスト
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
        // 探している値
	value := 5
        // バイナリサーチ
	fmt.Println(binarySearch(nums, value))
        // 再帰を使用したバイナリサーチ
        fmt.Println(binarySearchRecursive(nums, value, 0, len(nums) - 1))
}
