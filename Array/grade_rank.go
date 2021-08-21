package array

import "fmt"

//bytedance 20210820面试原题
//给定成绩列表，输出分数到rank的排名
//要求O(n)时间,O(1)空间
//0<=分数<=100
type MyMap struct {
	numPerGrade [100]int
	rank        [100]int
}

func (m *MyMap) build(grades []int) {
	for _, grade := range grades {
		m.numPerGrade[grade-1] += 1
	}
	for i := 99; i >= 0; i-- {
		if i == 99 {
			m.rank[i] = op(m.numPerGrade[i] > 0, 1, 0)
		} else {
			m.rank[i] = op(m.numPerGrade[i] > 0, 1, 0) + m.rank[i+1]
		}
	}
}

func (m *MyMap) transverse() {
	for grade, num := range m.numPerGrade {
		if num > 0 {
			fmt.Printf("grade=%d,rank=%d\n", grade+1, m.rank[grade])
		}
	}
}

func op(con bool, v1, v2 int) int {
	if con {
		return v1
	}
	return v2
}
