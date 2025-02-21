/*
215. Kth Largest Element in an Array
Solved
Medium
Topics
Companies
Given an integer array nums and an integer k, return the kth largest element in the array.

Note that it is the kth largest element in the sorted order, not the kth distinct element.

Can you solve it without sorting?

 

Example 1:

Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
Example 2:

Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
 

Constraints:

1 <= k <= nums.length <= 105
-104 <= nums[i] <= 104
*/
func findKthLargest(nums []int, k int) int {
    p, q := 0, len(nums)-1
    for q - p > 1 {
        pivot := int(math.Floor(float64(nums[p]+nums[q])/2))
        i, j := p, q
        for i < j {
            if nums[i] <= pivot {
                i++
            } else if nums[j] > pivot {
                j--
            } else {
                nums[i], nums[j] = nums[j], nums[i]
                i++
                j--
            }
        }
        if nums[i] > pivot {
            i--
        }



        np, nq := p, q
        if q-i >= k {
            np = i+1
        } else {
            nq, k = i, k+i-q
        }

        if np == p && nq == q {
            minid, maxid := p, q
            for i := p; i <= q; i++ {
                if nums[i] > nums[maxid] {
                    maxid = i
                }
                if nums[i] < nums[minid] {
                    minid = i
                }
            }
            
            if nums[minid] == nums[maxid] {
                return nums[minid]
            } else if k == 1 {
                return nums[maxid]
            } else if k == q-p+1 {
                return nums[minid]
            } else {
                nums[q], nums[maxid] = nums[maxid], nums[q]
                q--
                k--
            }
        } else {
            p, q = np, nq
        }
    }

    if q == p {
        return nums[p]
    } else {
        if k == 2 {
            return min(nums[p], nums[q])
        } else {
            return max(nums[p], nums[q])
        }
    }
}

func help(nums []int, p,q,k int) int {
    if p == q {
        return nums[p]
    } else if p+1 == q {
        if k == 2 {
            return min(nums[p], nums[q])
        } else {
            return max(nums[p], nums[q])
        }
    }

    mmin := math.MaxInt
    mmax := math.MinInt
    for i := p; i <= q; i++ {
        if nums[i] > mmax {
            mmax = nums[i]
        }
        if nums[i] < mmin {
            mmin = nums[i]
        }
    }
    if mmin == mmax {
        return mmin
    }

    pivot := (mmin+mmax)/2
    i, j := p, q
    for i < j {
        if nums[i] <= pivot {
            i++
        } else if nums[j] > pivot {
            j--
        } else {
            nums[i], nums[j] = nums[j], nums[i]
            i++
            j--
        }
    }
    if nums[i] > pivot {
        i--
    }

    //fmt.Println("k=%v, nums=%v, i=%v, j=%v", k, nums, i, j)
    //fmt.Println("k=%v, left=%v", k, left)
    //fmt.Println("k=%v, right=%v", k, right)
    if q-i >= k {
        return help(nums, i+1, q, k)
    } else {
        return help(nums, p, i, k+i-q)
    }
}
