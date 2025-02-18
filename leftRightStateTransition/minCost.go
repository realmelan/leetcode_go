/*
2448. Minimum Cost to Make Array Equal
Solved
Hard
Topics
Companies
Hint
You are given two 0-indexed arrays nums and cost consisting each of n positive integers.

You can do the following operation any number of times:

Increase or decrease any element of the array nums by 1.
The cost of doing one operation on the ith element is cost[i].

Return the minimum total cost such that all the elements of the array nums become equal.

 

Example 1:

Input: nums = [1,3,5,2], cost = [2,3,1,14]
Output: 8
Explanation: We can make all the elements equal to 2 in the following way:
- Increase the 0th element one time. The cost is 2.
- Decrease the 1st element one time. The cost is 3.
- Decrease the 2nd element three times. The cost is 1 + 1 + 1 = 3.
The total cost is 2 + 3 + 3 = 8.
It can be shown that we cannot make the array equal with a smaller cost.
Example 2:

Input: nums = [2,2,2,2,2], cost = [4,2,8,1,3]
Output: 0
Explanation: All the elements are already equal, so no operations are needed.
 

Constraints:

n == nums.length == cost.length
1 <= n <= 105
1 <= nums[i], cost[i] <= 106
Test cases are generated in a way that the output doesn't exceed 253-1
*/
func minCost(nums []int, cost []int) int64 {
    n := len(nums)
    type pair struct {
        num int
        cost int64
    }
    totalCost := int64(0)
    costs := make([]pair, n)
    for i := 0; i < n; i++ {
        costs[i] = pair{nums[i], int64(cost[i])}
        totalCost += int64(cost[i])
    }
    sort.Slice(costs, func(i,j int)bool{
        return costs[i].num < costs[j].num
    })

    // now find the index i such that sum(costs[0..i].cost) is closed to half of sum(costs[0..n-1].cost)
    sum := int64(costs[0].cost)
    k := 1
    for ; k < n; k++ {
        nextSum := sum + costs[k].cost
        if  sum >= totalCost / 2 || (sum <= totalCost / 2 && nextSum > totalCost / 2) {
            break
        }
        sum = nextSum
    }

    // compare costs (->costs[k-1].num) and costs(->costs[k].num)
    cost1 := int64(0)
    target1 := costs[k-1].num
    for i := 0; i < k; i++ {
        cost1 += costs[i].cost * int64(target1 - costs[i].num)
    }
    for i := k; i < n; i++ {
        cost1 += costs[i].cost * int64(costs[i].num - target1)
    }

    cost2 := int64(0)
    target2 := costs[k].num
    for i := 0; i < k; i++ {
        cost2 += costs[i].cost * int64(target2 - costs[i].num)
    }
    for i := k; i < n; i++ {
        cost2 += costs[i].cost * int64(costs[i].num - target2)
    }
    
    return min(cost1, cost2)
}
