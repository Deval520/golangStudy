/*
	包的注解 对于多文件包主需要在任意文件中注解一次即可
*/
package one

func Even(i int) bool {
	return i%2 == 0
}
func even(i int) bool {
	return i%2 == 0
}
