/*

630. Course Schedule III
Solved
Hard
Topics
Companies
Hint
There are n different online courses numbered from 1 to n. You are given an array courses where courses[i] = [durationi, lastDayi] indicate that the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.

You will start on the 1st day and you cannot take two or more courses simultaneously.

Return the maximum number of courses that you can take.

 

Example 1:

Input: courses = [[100,200],[200,1300],[1000,1250],[2000,3200]]
Output: 3
Explanation: 
There are totally 4 courses, but you can take 3 courses at most:
First, take the 1st course, it costs 100 days so you will finish it on the 100th day, and ready to take the next course on the 101st day.
Second, take the 3rd course, it costs 1000 days so you will finish it on the 1100th day, and ready to take the next course on the 1101st day. 
Third, take the 2nd course, it costs 200 days so you will finish it on the 1300th day. 
The 4th course cannot be taken now, since you will finish it on the 3300th day, which exceeds the closed date.
Example 2:

Input: courses = [[1,2]]
Output: 1
Example 3:

Input: courses = [[3,2],[4,3]]
Output: 0
 

Constraints:

1 <= courses.length <= 104
1 <= durationi, lastDayi <= 104

*/

func scheduleCourse(courses [][]int) int {
    // sort courses by last day in ascending order and duration if a tie
    // at a given day, always favor the course ending the earliest.
    sort.Slice(courses, func(i, j int) bool {
        if courses[i][1] != courses[j][1] {
            return courses[i][1] < courses[j][1]
        } else {
            return courses[i][0] < courses[j][0]
        }
    })
    var q PQ
    current := 0
    for _, co := range courses {
        current += co[0]
        heap.Push(&q, co)
        if current > co[1] {
            c := heap.Pop(&q).([]int)
            current -= c[0]
        }
    }
    return q.Len()
}

type PQ [][]int
func (p PQ) Len() int {
    return len(p)
}
func (p PQ) Less(i, j int) bool {
    return p[i][0] > p[j][0] // pop longest one
}
func (p PQ) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}
func (p *PQ) Push(o any) {
    *p = append(*p, o.([]int))
}
func (p *PQ) Pop() any {
    res := (*p)[len(*p)-1]
    *p = (*p)[:len(*p)-1]
    return res
}
