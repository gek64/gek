package gMath

import (
    "math/rand"
    "time"
)

// randSlice 生成长度为n,每个值最大为max的随机切片
func randSlice(n int, max int) []int {
    // 根据时间的随机数seed
    rand.Seed(time.Now().Unix())
    slice := make([]int, n)
    for i := 0; i < n; i++ {
        slice[i] = rand.Intn(max)
    }
    return slice
}

// SelectionSort 选择排序
func SelectionSort(slice []int) []int {
    length := len(slice)
    // i当前已完成排序的下标
    for i := 0; i < length; i++ {
        min := i
        for j := i + 1; j < length; j++ {
            if slice[min] > slice[j] {
                min = j
            }
        }
        slice[i], slice[min] = slice[min], slice[i]
    }
    return slice
}

// BubbleSort 冒泡排序
func BubbleSort(slice []int) []int {
    length := len(slice)
    for i := 0; i < length; i++ {
        // length-i 本次运行所需要的对比次数
        for j := 0; j < length-i-1; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
    return slice
}

// QuickSort 快速排序(递归)
func QuickSort(arr []int) []int {
    return _quickSort(arr, 0, len(arr)-1)
}

func _quickSort(arr []int, left, right int) []int {
    if left < right {
        partitionIndex := partition(arr, left, right)
        _quickSort(arr, left, partitionIndex-1)
        _quickSort(arr, partitionIndex+1, right)
    }
    return arr
}

func partition(arr []int, left, right int) int {
    pivot := left
    index := pivot + 1

    for i := index; i <= right; i++ {
        if arr[i] < arr[pivot] {
            swap(arr, i, index)
            index++
        }
    }
    swap(arr, pivot, index-1)
    return index - 1
}

func swap(arr []int, i, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}
