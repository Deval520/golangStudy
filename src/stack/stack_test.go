package stack

import "testing"

func TestPushPop(t *testing.T) {
	a := &Stack{}
	a.Push(8)
	if a.Pop() != 8 {
		t.Log("pop 不返回 8")
		t.Fail()
	}
}
