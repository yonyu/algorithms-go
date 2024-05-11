package arrays

import (
	"testing"
)

func TestKadaneBruteForce(t *testing.T) {
	nums1 := []int {4, -1, 2, -7, 3, 4}
  result := KadaneBruteForce(nums1)

	if result != 7 {
		t.Fatal("Result is incorrect:", result)
	}
}

func TestKodane(t *testing.T) {
	nums1 := []int {4, -1, 2, -7, 3, 4}
	result := Kadane(nums1)

	if result != 7 {
		t.Fatal("Result is incorrect:", result)
	}
}

func TestLargestSubarraySlidingWindow1(t *testing.T) {
	nums1 := []int {4, -1, 2, -7, 3, 4}
	_, _, result := LargestSubarraySlidingWindow(nums1)

	if result != 7 {
		t.Fatal("Result is incorrect:", result)
	}
}

func TestLargestSubarraySlidingWindow2(t *testing.T) {
	nums1 := []int {-2, -3, 4, -1, -2, 1, 5, -3}
	L, R, result := LargestSubarraySlidingWindow(nums1)

	if L != 2 && R!= 6 && result != 7 {
		t.Fatal("Result is incorrect:", result)
	}
}