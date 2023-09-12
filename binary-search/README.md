# 二分查找
关键词：
- array
- target
- sorted
- equal or close to target
- ...

## 递归法

## 迭代法
- 找确定值，返回任意一个即可
- 找确定值，返回第一个值
- 找确定值，返回最后一个值
- 找不确定值，比如大于等于的某一个值

万能模板：
```go
package main

func BinarySearch(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		//mid := (start + end) / 2
		mid := start + (end-start)/2 // 相比于上面的写法，当值特别大的时不会出现溢出
        // 不+1，不-1，递归条件也会终止
		if nums[mid] > target {
			end = mid
		} else if nums[mid] < target {
			start = mid
		} else {
			end = mid   // 需要修改的地方
		}
	}
	
	// 这里的顺序需要修改
	if nums[start] == target {
		return start
	}
	if nums[end] == target {
		return end
	}
	return -1
}
```