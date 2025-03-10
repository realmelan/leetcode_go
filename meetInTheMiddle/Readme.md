# Meet in the middle

"Meet in the middle" is a strategy to divide 1 large dataset into 2 smaller datasets, then apply algorithms on the two smaller datasets, and finally combine outputs from the smaller dataset to produce the solution to the large dataset.

The reason for splitting the dataset is, time complexity is normally O(k^n), for instance, enumerating all subsets or all subsets of a given size. By splitting the dataset into 2 smaller ones, the computation cost is k^(n/2), dramatically smaller than k^n.

# Leetcode problems
* [2035. Partition Array Into Two Arrays to Minimize Sum Difference](https://leetcode.com/problems/partition-array-into-two-arrays-to-minimize-sum-difference/)
* 

# Reference
* https://medium.com/@sherlock_ed/programming-meet-in-the-middle-technique-5025dbc1c6b6

