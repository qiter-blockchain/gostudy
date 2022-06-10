package main

import "fmt"

func main() {
	var array = [...]int{24, 69, 80, 57, 13, 34, 1, 9, 56, 90, 44, 11, 16, 78, 99, 4}
	fmt.Println("排序前 array=", array)
	BubbleSort(array[:])
	fmt.Println("冒泡排序后 array=", array)
	in := 5
	findResult := BinaryFind(in, array[:], 0, len(array))
	fmt.Println("二分查找 in=", in, " 数组array=", array)
	fmt.Println("二分查找结果下标为： index=", findResult)
}

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for j := 0; j < len(arr)-1; j++ {
		for i := 0; i < len(arr)-1-j; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
}

// BinaryFind 对有序数组进行二分查找
func BinaryFind(in int, arr []int, leftIndex int, rightIndex int) int {

	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		return -1
	}

	if arr[middle] > in {
		//查找范围  leftIndex  -  middle-1
		return BinaryFind(in, arr, leftIndex, middle-1)
	} else if arr[middle] < in {
		//查找范围  middle+1   - rightIndex
		return BinaryFind(in, arr, middle+1, rightIndex)
	} else {
		//找到了
		return middle
	}

	return -1
}
