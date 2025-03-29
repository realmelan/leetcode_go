/*

1353. Maximum Number of Events That Can Be Attended
Solved
Medium
Topics
Companies
Hint
You are given an array of events where events[i] = [startDayi, endDayi]. Every event i starts at startDayi and ends at endDayi.

You can attend an event i at any day d where startTimei <= d <= endTimei. You can only attend one event at any time d.

Return the maximum number of events you can attend.

 

Example 1:


Input: events = [[1,2],[2,3],[3,4]]
Output: 3
Explanation: You can attend all the three events.
One way to attend them all is as shown.
Attend the first event on day 1.
Attend the second event on day 2.
Attend the third event on day 3.
Example 2:

Input: events= [[1,2],[2,3],[3,4],[1,2]]
Output: 4
 

Constraints:

1 <= events.length <= 105
events[i].length == 2
1 <= startDayi <= endDayi <= 105

*/

func maxEvents(events [][]int) int {
    sort.Slice(events, func(i, j int) bool {
        if events[i][0] != events[j][0] {
            return events[i][0] < events[j][0]
        } else {
            return events[i][1] < events[j][1]
        }
    })
    q := ReadyQ(make([][]int, 0))
    res, current := 0, 0
    i, n := 0, len(events)
    for q.Len() > 0 || i < n {
        if q.Len() == 0 {
            current = events[i][0]
        }
        for ; i < n && events[i][0] <= current; i++ {
            if events[i][1] < current {
                continue
            }
            heap.Push(&q, events[i])
        }

        //fmt.Println("cur=%v, q=%v", current, q)
        for q.Len() > 0 {
            e := heap.Pop(&q).([]int)
            if e[1] < current {
                continue
            }
            current++
            res++
            break
        }
    }
    return res
}

type ReadyQ [][]int
func (q ReadyQ) Len() int {
    return len(q)
}
func (q ReadyQ) Less(i, j int) bool {
    return q[i][1] < q[j][1]
}
func (q ReadyQ) Swap(i, j int) {
    q[i], q[j] = q[j], q[i]
}
func (q *ReadyQ) Push(o any) {
    *q = append(*q, o.([]int))
}
func (q *ReadyQ) Pop() any {
    res := (*q)[len(*q)-1]
    (*q) = (*q)[:len(*q)-1]
    return res
}
