package one

import "testing"

//testing包
//
/**
 * go test 只会执行测试函数 每个测试函数都有相同的标识，名字以Test开头
 */

func TestEven(t *testing.T) {
	if !Even(2) {
		t.Log("2 should be even!")
		t.Fail()
	}
}
