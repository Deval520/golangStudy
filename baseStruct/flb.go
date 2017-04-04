package main
import (
  "fmt"
)
func main(){
  var a int = 1
  var p *int= &a //申明一个指针
  fmt.Println(*p)
  fmt.Println(p)
}
