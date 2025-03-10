# two pointers

From: https://www.reddit.com/r/leetcode/comments/18g9383/twopointer_technique_an_indepth_guide_concepts/

Two-Pointer Technique, an In-Depth Guide: Concepts Explained | Questions to Try | Visuals and Animations
Intervew Prep
The two-pointer technique I’m referring to here involves using two pointers that start at opposite ends of an array and gradually move towards each other before meeting in the middle.

left                            right
 ↓                                ↓
 --- --- --- --- --- --- --- --- ---
| 2 | 1 | 2 | 0 | 1 | 0 | 1 | 0 | 1 |
 --- --- --- --- --- --- --- --- ---
This technique should be your go-to when you see a question that involves searching for a pair (or more!) of elements in an array that meet a certain criteria.

In this guide, we'll start by understanding how the technique produces the efficient O(n) time-complexity solutions that those questions require (answer: by eliminating pairs). I'll then provide follow-up questions for you to try, with lots of visuals and interactive animations to help along the way.

Sample Problem: Two Sum (easy)
Starting with a sorted array of integers, find a pair of numbers that sum to the given target.

Let’s walkthrough how the two-pointer technique eliminates unnecessary pairs from the search when the input array is [1, 3, 4, 6, 8, 10, 13] and target = 13


Step 1: Initialize pointers at opposite ends of array and sum the two elements together. This represents the first pair we are considering in our search.


left                     right
 ↓                         ↓
 --- --- --- --- --- ---- ----   
| 1 | 3 | 4 | 6 | 8 | 10 | 13 |     current_sum = array[left] + array[right] = 14
 --- --- --- --- --- ---- ----   

Step 2: Compare current_sum with target. Since current_sum > target, move the right pointer backwards.

To see why, notice that all other pairs that use 13 are also greater than our current_sum.So by moving the right pointer to 10, we can eliminate those unnecessary pairs from our search.


left                 right
 ↓                     ↓
 --- --- --- --- --- ---- ----
| 1 | 3 | 4 | 6 | 8 | 10 | 13 |          by moving right pointer back...
 --- --- --- --- --- ---- ----



		   21--------
                  |         | 
           17-----|-------- | 
          |       |       | | 
 --- --- --- --- --- ---- ----
| 1 | 3 | 4 | 6 | 8 | 10 | 13 |     we eliminated these pairs from our search  
 --- --- --- --- --- ---- ---- 
      |       |       |   | | |
       16-----|-------|---- | |
              |       |     | |
               19-----|------ |
                      |       |
                       23------

Step 3: Compare current_sum with target. Since current_sum < target, move the left pointer forwards. This follows similar reasoning to the step above: all other pairs that use element 1 are less than our target, so we should move our left pointer forward to eliminate those pairs.


left                 right
 ↓                     ↓
 --- --- --- --- --- ---- ----
| 1 | 3 | 4 | 6 | 8 | 10 | 13 |                move left pointer forwards
 --- --- --- --- --- ---- ----

============= TO =============

     left            right
      ↓                ↓
 --- --- --- --- --- ---- ----
| 1 | 3 | 4 | 6 | 8 | 10 | 13 |                   current_sum = 13
 --- --- --- --- --- ---- ----
Termination: Repeat process until current_sum == target, like it does here. Or, if the pointers meet at the same index, then a pair was not found.

Time Complexity: O(n). This is done in a single pass. By using the two-pointer technique, we avoid the nested for-loop required by the brute force solution.

Click here for a more in-depth breakdown of this question, including an interactive animation of the Python solution.

Try It Yourself: Follow-up Questions
Container With Most Water (medium)

Hint: Instead of summing the elements at each pointer, compare their values instead. Which containers can you eliminate?

Stuck? This link helps you visualize each step alongside the Python implementation.


3Sum (medium)


                     i, left, right represent current triplet      

          i  left              right 
          ↓   ↓                  ↓              
         ---- ---- ---- --- --- ---                   
        | -4 | -1 | -1 | 0 | 1 | 2 |             can you use two sum?
         ---- ---- ---- --- --- ---              

                   ...

               i   left         right
               ↓    ↓            ↓              
         ---- ---- ---- --- --- ---             
        | -4 | -1 | -1 | 0 | 1 | 2 |         can you use two sum again?    
         ---- ---- ---- --- --- ---              
Hint: sort the array, iterate over each item, repeatedly apply two sum.

This link helps you visualize each step of the implementation (without showing the Python implementation)


Valid Triangle Number (medium)

Hint: sort the array, then use the triangle inequality, which states that if a triangle has sides of lengths a, b, and c, then all three of (1) a + b > c (2) a + c > b (3) c + b > a must be true.


                   i, left, right represent current triplet
        

        left              right  i
         ↓                  ↓    ↓              
         --- --- --- ---- ---- ----                   
        | 4 | 6 | 9 | 11 | 15 | 18 |        which triplets can you eliminate?
         --- --- --- ---- ---- ----         
This link helps you visualize each step of the implementation (without showing the Python implementation)


3Sum Closest

A variation of 3Sum.


Summary
If a question involves searching for a pair (or more!) items in an array that meet a certain criteria, see if you can use the two-pointer technique to come up with an efficient solution.

The questions linked here use the two-pointer technique to eliminate unnecessary pairs from the search, producing O(n) solutions compared to the O(n2) brute-force solutions.

To use the technique: initialize the pointers (typically at opposite ends of the array, but not always). Look at the values at each pointer. From those values, think about how to move each pointer so that you can eliminate unecessary pairs from the search.


Bonus! Partitioning Arrays
The two-pointer technique can also be used to solve problems that involve partition arrays into different regions. For these questions, each pointer represents where the next element belonging to that region should go.


     (next 0 here!)  (next 2 here!)   
         left           right                      Sorting array of 0, 1, 2s
          ↓               ↓
 --- --- --- --- --- --- --- --- --- ---
| 0 | 0 | 1 | 1 | 2 | 1 | 1 | 2 | 2 | 2 |
 --- --- --- --- --- --- --- --- --- ---
|______||_______||__________||______|       
    0s     1s      unsorted     2s
Example: Sort Colors (medium)
Given an unsorted array nums with n integers that are either 0, 1, or 2.
Sort the array in-place in ascending order. 

Solve this problem in one-pass without any extra space.
We'll actually initialize 3 pointers:

left and right at opposite ends of the array. The left pointer represents the position of the next 0, and the right pointer represents the position of the next 2.

i at the beginning of the array. This pointer represents the current element we are trying to sort, as well as the boundary of the "ones" region.

These three pointers split our array into four regions, an unsorted region, and 0s, 1s, and 2s regions, which are all empty and not shown.

         i
        left                right
         ↓                    ↓
         --- --- --- --- --- --- 
        | 2 | 1 | 2 | 0 | 1 | 0 |
         --- --- --- --- --- --- 
        |_______unsorted________|
The idea here is that we iterate until i crosses right. At each iteration:

if nums[i] == 0, we swap i with the element at the left pointer, move left pointer forward and increment i

if nums[i] == 1, we increment i

if nums[i] == 2, we swap i with the element at the right pointer, move right pointer backward.


         i
        left                right
         ↓                    ↓
         --- --- --- --- --- --- 
        | 2 | 1 | 2 | 0 | 1 | 0 |        Start 
         --- --- --- --- --- ---         (empty regions not shown)
        |_______unsorted________|   
    
      
         i
        left            right
         ↓                ↓              Step 1: nums[i] == 2
         --- --- --- --- --- --- 
        | 0 | 1 | 2 | 0 | 1 | 2 |        swap i with right
         --- --- --- --- --- ---         move right pointer back
        |______unsorted_____||__|   
                              2s

              i
             left       right
              ↓           ↓             Step 2: nums[i] == 0
         --- --- --- --- --- ---
        | 0 | 1 | 2 | 0 | 1 | 2 |       swap i with left
         --- --- --- --- --- ---        move left pointer forward
        |__||___unsorted____||__|       increment i
         0s                   2s


             left i     right
              ↓   ↓       ↓             Step 3: nums[i] == 1
         --- --- --- --- --- ---
        | 0 | 1 | 2 | 0 | 1 | 2 |       increment i
         --- --- --- --- --- ---     
        |___||__||_unsorted_||__|       
          0s  1s              2s


             left i  right
              ↓   ↓   ↓                  Step 4: nums[i] == 2
         --- --- --- --- --- ---
        | 0 | 1 | 1 | 0 | 2 | 2 |       swap i with right
         --- --- --- --- --- ---        move right pointer back
        |___||__||______||______|       
          0s  1s unsorted   2s


                    ...


               left right i 
                  ↓   ↓   ↓               Termination (i > right)
         --- --- --- --- --- ---
        | 0 | 0 | 1 | 1 | 2 | 2 |         return sorted array
         --- --- --- --- --- ---
        |_______||______||______|
           0s       1s      2s
Time Complexity: Single pass, O(n).Space Complexity: O(1).

Click here a more in-depth breakdown of this question, including an interactive animation of the Python solution.

Try It Yourself
Move Zeroes (easy)

Not exactly the two-pointer technique described here, but good practice for using a pointer to represent a region of an array.

         what goes here? 
          nextNonZero i
              ↓       ↓  
         --- --- --- --- ---- 
        | 1 | 0 | 0 | 3 | 12 | 
         --- --- --- --- ---- 

Hint: use a pointer to represent the position of the next non-zero you find.

This link helps you visualize each step of the implementation (without showing the Python implementation)


Partition Array According to Given Pivot (medium)

Hint: follow a similar approach to sort colors, but copy items to a new output array to maintain relative ordering.



Summary
If a question calls for partitioning an array into different regions: initialize one pointer for each region you need to create.

Then iterate over the array and place each element in the correct position (as dictated by the pointer).

Move the pointer to indicate where the next element that belongs in that region should go.


I love breaking down the algorithm patterns that will help you land your next dream job in tech. There will be many more coming in the near future. If you found this guide helpful, or if there is anything you would like me to cover in the future, please leave a comment! It means a lot :)

