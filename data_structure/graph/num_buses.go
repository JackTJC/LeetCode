package graph

//timeout
func numBusesToDestination(routes [][]int, source int, target int) int {
	graph := make(map[int]map[int]struct{})
	open := make(map[int]struct{})
	close := make(map[int]struct{})
	for _, route := range routes {
		n := len(route)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				n1, n2 := route[i], route[j]
				open[n1] = struct{}{}
				open[n2] = struct{}{}
				_, ok := graph[n1]
				if !ok {
					graph[n1] = make(map[int]struct{})
				}
				graph[n1][n2] = struct{}{}
				_, ok = graph[n2]
				if !ok {
					graph[n2] = make(map[int]struct{})
				}
				graph[n2][n1] = struct{}{}
			}
		}
	}
	dist := make(map[int]int)
	delete(open, source)
	//初始化为无限大
	for n := range open {
		dist[n] = 1 << 31
	}
	//source到source为0，且将source加入close表
	close[source] = struct{}{}
	dist[source] = 0
	change := true
	for change {
		change = false
		for c := range close {
			for neibour := range graph[c] {
				//取较小的值
				if dist[c]+1 < dist[neibour] {
					if dist[neibour] == 1<<31 {
						delete(open, neibour)
						close[neibour] = struct{}{}
					}
					dist[neibour] = dist[c] + 1
					change = true
				}
			}
		}
	}
	if dist[target] == 1<<31 {
		return -1
	}
	return dist[target]
}
