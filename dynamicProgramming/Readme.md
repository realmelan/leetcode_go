# dynamic programming

All dynamic programming solutions are brute force solutions, because we need to compute target value for each state.
The size of the input won't be large. So, look at the boundary as scale hints.


## determine state
consider the state variables, like position, number of elements left and target value

## state enumeration
After identifying states, we need to iteratively enumerate the states.
* Normally it is quite straight forward like number of elements, etc. We can simply use integers to represent the state
* In case when we need to enumerate on subsets of a list of elements, then we can use bit mask starting from 0 to 1<<cont.

## state transition
This is the core part of dynamic programming. We compute larger states using smaller states(already computed).

Typical way is, when iterating/recursing to current state, we break it down into smaller states, and use target value
of smaller states to compute target value for current state and record it in a lookup table.

Another way, when iterating, current state is already computed/updated by iterating previous smaller states. Now that we have target value
for current statue, we can find ways to transform current state into larger states, for example, by adding one more element. This is
often used in bit mask enumeration.



# method 1: DFS + Memoization

```go
func dfs(state variable, memo ) int {
  // base case

  // lookup memo table and return value if already computed

  // using state variables, compute target value
  // set target value into memo and return
}
```

1. [1815. Maximum Number of Groups Getting Fresh Donuts (Hard)](https://leetcode.com/problems/maximum-number-of-groups-getting-fresh-donuts/)


# method 2: Iterative/Tabulation
This is the typical bottom up solution for dynamic programing. It computes all states iteratively, starting from the base case, like length=1,
and gradually increases the size of the states. 

When calculating target value for each state, binary search or monotonic queue may be used to speed up.


1. [1000. Minimum Cost to Merge Stones](https://leetcode.com/problems/minimum-cost-to-merge-stones/description/)
2. [300. Longest Increasing Subsequence](https://leetcode.com/problems/longest-increasing-subsequence/)
3. [1478. Allocate Mailboxes](https://leetcode.com/problems/allocate-mailboxes/description/)


# method 3: BFS + DP
This handles the case when there are multiple ways to arrive at a given state, but they can't be done in just one loop. For instance, traversing a grid from different
direction with different cost.

Then we can save the optimal value for a given state in a table, and every time we arrive at a given state, compare the value against the value in the table, and then update
 the table and continue BFS from current state if current state is optimal.

 1. [1368. Minimum Cost to Make at Least One Valid Path in a Grid (Hard)](https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/)


# Complexity analysis
The complexity is determined by number of states we need to compute(state enumerating), and cost for each state(state transition).


# Reference
1. https://cp-algorithms.com/dynamic_programming/intro-to-dp.html
