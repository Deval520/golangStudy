/*
  生成一个Stack数据结构实现切片的存和取得记忆功能
*/
package stack

type Stack struct {
	//标记位
	i int
	//存储单元切片
	data [10]int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
	// return s.data
}

func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
