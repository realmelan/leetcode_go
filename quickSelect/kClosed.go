/*
973. K Closest Points to Origin
Solved
Medium
Topics
Companies
Given an array of points where points[i] = [xi, yi] represents a point on the X-Y plane and an integer k, return the k closest points to the origin (0, 0).

The distance between two points on the X-Y plane is the Euclidean distance (i.e., √(x1 - x2)2 + (y1 - y2)2).

You may return the answer in any order. The answer is guaranteed to be unique (except for the order that it is in).

 

Example 1:


Input: points = [[1,3],[-2,2]], k = 1
Output: [[-2,2]]
Explanation:
The distance between (1, 3) and the origin is sqrt(10).
The distance between (-2, 2) and the origin is sqrt(8).
Since sqrt(8) < sqrt(10), (-2, 2) is closer to the origin.
We only want the closest k = 1 points from the origin, so the answer is just [[-2,2]].
Example 2:

Input: points = [[3,3],[5,-1],[-2,4]], k = 2
Output: [[3,3],[-2,4]]
Explanation: The answer [[-2,4],[3,3]] would also be accepted.
 

Constraints:

1 <= k <= points.length <= 104
-104 <= xi, yi <= 104
*/

func kClosest(points [][]int, k int) [][]int {
    // maintain a list of k closed points in sorted order
    // for a new point, insert it in the list using binary search
    // or a priority queue can be used
    ps := make([]*pair, len(points))
    for i, p := range points {
        ps[i] = &pair{p, distance(p)}
    }

    i, j := 0, len(ps)-1
    for {
        x := partition(ps, i, j)
        //fmt.Println("x=%v, i=%v, j=%v, ps=%v", x, i, j, ps)
        if x == k-1 {
            break
        } else if x > k-1 {
            j = x-1
        } else {
            i = x+1
        }
    }

    res := make([][]int, k)
    for i = 0; i < k; i++ {
        res[i] = ps[i].point
    }
    return res
}

type pair struct {
        point []int
        dist int
    }
func partition(points []*pair, lo, hi int) int {
    i := lo+1
    j := hi
    pivot := points[lo]
    for i <= j {
        if points[i].dist < pivot.dist {
            i++
        } else if points[j].dist > pivot.dist {
            j--
        } else if i >= j {
            break
        } else {
            points[i], points[j] = points[j], points[i]
            i++
            j--
        }
    }
    points[lo], points[j] = points[j], points[lo]
    return j
}

func kClosest2(points [][]int, k int) [][]int {
    // maintain a list of k closed points in sorted order
    // for a new point, insert it in the list using binary search
    // or a priority queue can be used
    res := make([]*pair, 0)
    for _, p := range points {
        pd := &pair{p, distance(p)}
        i, _ := slices.BinarySearchFunc(res, pd, func(a, t *pair) int{
            return a.dist - t.dist
        })
        if i == k {
            continue
        }
        if len(res) < k {
            res = append(res, pd)
        }
        //fmt.Println("pd=%v, %p, i=%v, res=%v", pd, pd, i, res)
        for j := len(res)-1; j > i; j-- {
            res[j] = res[j-1]
        }
        res[i] = pd
        //fmt.Println("pd=%v, i=%v, res=%v", pd, i, res)
    }
    
    r := make([][]int, 0)
    for _, p := range res {
        r = append(r, p.point)
    }
    return r
}

func distance(p []int) int {
    return p[0]*p[0] + p[1]*p[1]
}
