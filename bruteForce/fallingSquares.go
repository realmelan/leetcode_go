/*

699. Falling Squares
Solved
Hard
Topics
Companies
Hint
There are several squares being dropped onto the X-axis of a 2D plane.

You are given a 2D integer array positions where positions[i] = [lefti, sideLengthi] represents the ith square with a side length of sideLengthi that is dropped with its left edge aligned with X-coordinate lefti.

Each square is dropped one at a time from a height above any landed squares. It then falls downward (negative Y direction) until it either lands on the top side of another square or on the X-axis. A square brushing the left/right side of another square does not count as landing on it. Once it lands, it freezes in place and cannot be moved.

After each square is dropped, you must record the height of the current tallest stack of squares.

Return an integer array ans where ans[i] represents the height described above after dropping the ith square.

 

Example 1:


Input: positions = [[1,2],[2,3],[6,1]]
Output: [2,5,5]
Explanation:
After the first drop, the tallest stack is square 1 with a height of 2.
After the second drop, the tallest stack is squares 1 and 2 with a height of 5.
After the third drop, the tallest stack is still squares 1 and 2 with a height of 5.
Thus, we return an answer of [2, 5, 5].
Example 2:

Input: positions = [[100,100],[200,100]]
Output: [100,100]
Explanation:
After the first drop, the tallest stack is square 1 with a height of 100.
After the second drop, the tallest stack is either square 1 or square 2, both with heights of 100.
Thus, we return an answer of [100, 100].
Note that square 2 only brushes the right side of square 1, which does not count as landing on it.
 

Constraints:

1 <= positions.length <= 1000
1 <= lefti <= 108
1 <= sideLengthi <= 106

*/
func fallingSquares(positions [][]int) []int {
    heights := make([][]int, 0)
    var res []int
    maxh := 0
    for _, square := range positions {
        left, right, height := square[0], square[0]+square[1], square[1]
        sqh := height
        for _, h := range heights {
            if right <= h[0] || left >= h[1] {
                continue
            }
            sqh = max(sqh, h[2] + height)
        }
        maxh = max(maxh, sqh)
        res = append(res, maxh)

        heights = append(heights, []int{left, right, sqh})
    }
    return res
}

func fallingSquares_wrong(positions [][]int) []int {
    // maintain a list of non-overlap ranges (x1, x2, h)
    // when dropping a new square, use range based change to calculate height of the
    // new square, and then loop over the ranges to update range list using the new height.

    // Verdict: it can't handle squares landing on its right corner.
    ranges := make([][]int, 0)
    var res []int
    for _, square := range positions {
        left, right := square[0], square[0] + square[1]
        changes := [][]int{{left, square[1]},{right, -square[1]}}
        for _, r := range ranges {
            changes = append(changes, []int{r[0], r[2]}, []int{r[1], -r[2]})
        }
        sort.Slice(changes, func(i, j int) bool{
            if changes[i][0] != changes[j][0] {
                return changes[i][0] < changes[j][0]
            } else {
                return changes[i][1] < changes[j][1]
            }
        })
        fmt.Println("changes=%v", changes)

        nh, curh := 0, 0
        for _, ch := range changes {
            curh += ch[1]
            if left <= ch[0] && ch[0] <= right {
                nh = max(nh, curh)
            }
        }



        var newRanges [][]int
        maxh, curh := 0, 0
        for _, ch := range changes {
            curh += ch[1]
            if len(newRanges) == 0 {
                newRanges = append(newRanges, []int{ch[0], math.MaxInt, curh})
            } else if newRanges[len(newRanges)-1][2] != curh {
                newRanges[len(newRanges)-1][1] = ch[0]
                newRanges = append(newRanges, []int{ch[0], math.MaxInt, curh})
            }
            
            maxh = max(maxh, curh)
        }
        res = append(res, maxh)
        ranges = newRanges
        fmt.Println("ranges=%v", ranges)
    }
    return res
}
func fallingSquares2(positions [][]int) []int {
    // maintain a list of non-overlap <range, height> pairs, each representing
    // the a segment of x-axis at some height. The ranges are sorted in ascending order
    // of x.
    var res []int
    ranges := make([][]int, 0) // {x1, x2, height}
    maxh := 0
    for _, square := range positions {
        // find the first heights position overlaps
        idx, found := slices.BinarySearchFunc(ranges, square, func(rangeHeight []int, pos []int) int {
            return rangeHeight[0] - pos[0]
        })
        if !found {
            // check whether position has overlap with previous range Height's 
            if idx - 1 >= 0 && ranges[idx-1][1] > square[0]{
                idx--
            }
        }
        // now starting from idx, calculate new range heights.
        // must first calculate height of the new square
        nh, i := square[1], idx
        for ; i < len(ranges) && square[0]+square[1] > ranges[i][0]; i++ {
            nh = max(nh, ranges[i][2] + square[1])
        }
        //fmt.Println("idx=%v, i=%v", idx, i)
        maxh = max(maxh, nh)
        res = append(res, maxh)

        newRanges := slices.Clone(ranges[0:idx])
        // add first half
        var prior []int
        if idx < len(ranges) && ranges[idx][0] < square[0] {
            prior = []int{
                ranges[idx][0], square[0], ranges[idx][2],
            }
            newRanges = append(newRanges, prior)
        }

        // add squ
        current := []int{square[0], square[0]+square[1], nh}
        newRanges = append(newRanges, current)

        // add last part
        var trail []int
        if i-1 >= 0 && ranges[i-1][1] > square[0]+square[1] {
            trail = []int{
                square[0]+square[1], ranges[i-1][1], ranges[i-1][2],
            }
            newRanges = append(newRanges, trail)
        }
        //fmt.Println("new ranges=%v", newRanges)
        //fmt.Println("ranges=%v", ranges)
        newRanges = append(newRanges, ranges[i:len(ranges)]...)
        //fmt.Println("new ranges=%v", newRanges)
        ranges = newRanges
        //fmt.Println("ranges=%v", ranges)
    }
    return res
}
