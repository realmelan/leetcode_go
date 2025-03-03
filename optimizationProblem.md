# How to tackle optimization problem?

## determine utility value/function
* this is the target the problem tries to optimize, like minimal cost/operations or maxmized value, etc

## determine solution space
* once the utility function determined, you can decide the solution space, like a range of values, or all subsets of given elements.
* Normally you need to go through all members of the solution space to find the optimal solution
* However, if target value is monotonic in the solution space, you may use __binary search__

## determine how utility value changes in the solution space

### binary search or not
* if utility value is monotonic in the solution space, consider __[binary search](../binarySearch)__
* else, you have to loop through all members of the solution space.

### determine how many utility values
* if the utility values(for each element or for all elements) in the solution space are range defined (stay the same for a range, and then change to another value for another (connected) range):
  * then capture the change at the end points of the ranges in a __difference array__ or __delta array__
  * break elements into smallest groups (maybe individual element) until they are independent when computing utility value
* if the utility value is jumpy (either no range for a utility value to stay the same, or there are too many ranges), then probably should go with __[dynamic programming](../dynamicProgramming)__.
