package main

func main(){
	
}
/*
 * 执行用时：16 ms, 在所有 Go 提交中击败了50.17%的用户
 * 内存消耗：6.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func findOrder(numCourses int, prerequisites [][]int) []int {
	queue := make([]int, 0)
	s := make(map[int]int)
	m := make(map[int][]int)
	for _, v := range prerequisites {
		s[v[0]] += 1
		if _, ok := m[v[1]]; !ok {
			m[v[1]] = make([]int, 0)
		}
		m[v[1]] = append(m[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if s[i] <= 0 {
			queue = append(queue, i)
		}
	}
	i := 0
	for i < len(queue) {
		if s[queue[i]] > 0{
			break
		}
		for _, v := range m[queue[i]] {
			s[v]--
			if s[v] == 0 {
				delete(s, v)
				queue = append(queue, v)
			}
		}
		if len(s) == 0 {
			return queue
		}
		i++
	}
	return []int{}
}
