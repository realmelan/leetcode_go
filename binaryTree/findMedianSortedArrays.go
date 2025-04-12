/*
4. Median of Two Sorted Arrays
Solved
Hard
Topics
Companies
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).

 

Example 1:

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
Example 2:

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
 

Constraints:

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// binary serach on one array, m1, then m2 can also be determined
	// check nums1[m1] > nums2[m2+1] or nums2[m2] < nums1[m1+1]
    // the partition holds (n1+n2+1)/2 elements from both
    // if n1+n2 is odd, then the median is the largest of the partition
    // else, median is average of largest of the left partition and smallest
    //  of the right partition
	n1, n2 := len(nums1), len(nums2)
    if n1 > n2 {
        return findMedianSortedArrays(nums2, nums1)
    }

    // invariant: select k elements from nums1, and 0 <= k <= n1
	// m might be -1 when no num in nums1 <= median of two arrays
	l, r := 0, n1
	m1, m2 := 0, 0
    var t1,t2,t3,t4 int
	for l <= r {
		m1 = (l + r) / 2
		m2 = (n1+n2+1)/2 - m1

		t1 = int(-1e7)
		if m1-1 >= 0 {
			t1 = nums1[m1-1]
		}
		t2 = int(1e7)
		if m1 < n1 {
			t2 = nums1[m1]
		}
		t3 = int(-1e7)
		if m2-1 >= 0  {
            t3 = nums2[m2-1]
		}
		t4 = int(1e7)
		if m2 < n2 {
			t4 = nums2[m2]
		}
		if t1 > t4 { // move l right
			r = m1 - 1
		} else if t3 > t2 {
			l = m1 + 1
		} else {
			break // found
		}
	}
    //fmt.Println("m1=%v, m2=%v", m1, m2)

	if (n1+n2)%2 == 1 {
        return float64(max(t1,t3))
    } else {
        return (float64(max(t1,t3)+min(t2,t4)))/2
    }

}
