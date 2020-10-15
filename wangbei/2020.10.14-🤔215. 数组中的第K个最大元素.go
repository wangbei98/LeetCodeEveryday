package main

func findKthLargest(nums []int, k int) int {
	var partition func(nums[]int,left int,right int)int
	partition = func(nums[]int,left int,right int)int{
		temp := nums[left]
		for left < right{
			for left < right && nums[right] <= temp{right--}
			nums[left] = nums[right]
			for left < right && nums[left] >= temp{left++}
			nums[right] = nums[left]
		}
		nums[left] = temp
		return left
	}
	var findK func(nums []int, targetIndex int, left int, right int)int
	findK = func(nums []int, targetIndex int, left int, right int)int{
		index := partition(nums,left,right)
		if index == targetIndex{
			return nums[index]
		}else if index < targetIndex{
			return findK(nums,targetIndex,index+1,right)
		}else{
			return findK(nums,targetIndex,left,index-1)
		}
	}
	return findK(nums,k-1,0,len(nums)-1)
}