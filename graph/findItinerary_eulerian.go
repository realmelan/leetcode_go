/*

332. Reconstruct Itinerary
Solved
Hard
Topics
Companies
You are given a list of airline tickets where tickets[i] = [fromi, toi] represent the departure and the arrival airports of one flight. Reconstruct the itinerary in order and return it.

All of the tickets belong to a man who departs from "JFK", thus, the itinerary must begin with "JFK". If there are multiple valid itineraries, you should return the itinerary that has the smallest lexical order when read as a single string.

For example, the itinerary ["JFK", "LGA"] has a smaller lexical order than ["JFK", "LGB"].
You may assume all tickets form at least one valid itinerary. You must use all the tickets once and only once.

 

Example 1:


Input: tickets = [["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
Output: ["JFK","MUC","LHR","SFO","SJC"]
Example 2:


Input: tickets = [["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
Output: ["JFK","ATL","JFK","SFO","ATL","SFO"]
Explanation: Another possible reconstruction is ["JFK","SFO","ATL","JFK","ATL","SFO"] but it is larger in lexical order.
 

Constraints:

1 <= tickets.length <= 300
tickets[i].length == 2
fromi.length == 3
toi.length == 3
fromi and toi consist of uppercase English letters.
fromi != toi

*/

func findItinerary(tickets [][]string) []string {
	cons := make(map[string]map[string]int)
	for _, t := range tickets {
		_, ok := cons[t[0]]
		if !ok {
			cons[t[0]] = make(map[string]int)
		}
		cons[t[0]][t[1]]++
	}

	start := "JFK"
	fp := []string{}
	euler(cons, start, &fp)
	slices.Reverse(fp)
	return fp
}

func euler(cons map[string]map[string]int, cur string, fp *[]string) {
	for len(cons[cur]) > 0 {
		next := "aaa"
		for k, v := range cons[cur] {
			if v > 0 {
				next = min(next, k)
			}
		}
		if next == "aaa" {
			break
		}
		cons[cur][next]--
		euler(cons, next, fp)

	}
	*fp = append(*fp, cur)
}
