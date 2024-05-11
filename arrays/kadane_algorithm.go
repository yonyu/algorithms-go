package arrays

import (
	"log"
)

// Given an array arr[] of size N. The task is to find the sum of the contiguous 
// subarray within a arr[] with the largest sum.
// non-empty subarray
// subarray: consecutive items
// if all are negative, return the maximum element

// brutal force: O(n^2)
func KadaneBruteForce(nums []int) int {
	if nums == nil || len(nums) == 0 {
		log.Fatal("input is nil or empty")
		return 0
	}
		
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		curSum := 0
		for j := i; j < len(nums); j++ {
			curSum += nums[j]
			maxSum = max(maxSum, curSum)
		}
	}
	return maxSum
}

// Kadane's Algorithm: O(n)
func Kadane(nums []int) int {
	if nums == nil || len(nums) == 0 {
		log.Fatal("input is nil or empty")
		return 0
	}
		
	maxSum := nums[0]
	curSum := 0
	for i := 0; i < len(nums); i++ {
		curSum = max(curSum, 0) // discard the current sum if it is negative
		curSum += nums[i]
		maxSum = max(maxSum, curSum)
	}
	return maxSum
}

// Return the left and right index of the max subarray sum,
// assuming there's exactly one result (no ties).
// Sliding window variation of Kadane's: O(n)
// two pointers
func LargestSubarraySlidingWindow(nums []int) (int, int, int) {
	if nums == nil || len(nums) == 0 {
		log.Fatal("input is nil or empty")
		return 0, 0, 0
	}

	maxL := 0
	maxR := 0
	L := 0
	maxSum := nums[0]
	curSum := 0
	for r := 0; r < len(nums); r++ {
		// discard the current sum if it is negative
		if curSum < 0 {
			curSum = 0
			L = r
		}
	
		curSum += nums[r]
		if curSum > maxSum {
			maxSum = curSum
			maxL, maxR = L, r
		}
	}
	return maxL, maxR, maxSum
}

func max(a, b int) int {
	if b < a {
		return a
	}
	return b
}