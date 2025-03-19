/*

1815. Maximum Number of Groups Getting Fresh Donuts
Solved
Hard
Topics
Companies
Hint
There is a donuts shop that bakes donuts in batches of batchSize. They have a rule where they must serve all of the donuts of a batch before serving any donuts of the next batch. You are given an integer batchSize and an integer array groups, where groups[i] denotes that there is a group of groups[i] customers that will visit the shop. Each customer will get exactly one donut.

When a group visits the shop, all customers of the group must be served before serving any of the following groups. A group will be happy if they all get fresh donuts. That is, the first customer of the group does not receive a donut that was left over from the previous group.

You can freely rearrange the ordering of the groups. Return the maximum possible number of happy groups after rearranging the groups.

 

Example 1:

Input: batchSize = 3, groups = [1,2,3,4,5,6]
Output: 4
Explanation: You can arrange the groups as [6,2,4,5,1,3]. Then the 1st, 2nd, 4th, and 6th groups will be happy.
Example 2:

Input: batchSize = 4, groups = [1,3,2,5,2,2,1,6]
Output: 4
 

Constraints:

1 <= batchSize <= 9
1 <= groups.length <= 30
1 <= groups[i] <= 109

*/

func maxHappyGroups_memo(batchSize int, groups []int) int {
    // dfs + memo
    count := make([]int, batchSize)
    res, total := 0, 0
    for _, g := range groups {
        mod := g % batchSize 
        if mod == 0 {
            res ++
        } else {
            count[mod]++
            total++
        }
    }

    i, j := 1, batchSize - 1
    for i < j {
        t := min(count[i], count[j])
        res += t
        count[i] -= t
        count[j] -= t
        i++
        j--
        total -= t*2
    }

    memo := make([]map[string]int, batchSize)
    for i := range memo {
        memo[i] = make(map[string]int)
    }

    return res + dfs(batchSize, 0, total, count, memo)
}

func dfs(b, lastBatch, left int, count []int, memo []map[string]int) int {
    if left == 0 {
        return 0
    }
    s := key(count)
    v, ok := memo[lastBatch][s]
    if ok {
        return v
    }

    res := 0
    for c, v := range count {
        if v == 0 {
            continue
        }

        fresh := 0
        if lastBatch == 0 {
            fresh = 1
        }
        count[c]--
        if c <= lastBatch {
            res = max(res, fresh + dfs(b, lastBatch - c, left-1, count, memo))
        } else {
            res = max(res, fresh + dfs(b, b+lastBatch - c, left-1, count, memo))
        }
        count[c]++
    }
    memo[lastBatch][s] = res
    return res
}

func key(count []int) string {
    var buf strings.Builder
    for _, v := range count {
        buf.WriteByte(byte('0')+byte(v))
    }
    return buf.String()
}
