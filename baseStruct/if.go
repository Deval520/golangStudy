package main

import (
  "fmt"
)
func main()  {
  //   fmt.Printf("a is: %d, b is %d\n", a,b)
  // }

  // var sum int
  // for i := 0; i < 1000000000; i++{
  //   sum += i
  // }
  // fmt.Printf("%d\n",sum)

  // a := 10
  // if a,b := 1, 2; a < b{
  //   fmt.Println("ok")
  // }
  // fmt.Println(a)

  // for i := 0; i < 5; i++{
  //   var p *int
  //   p = &i
  //   fmt.Println(p)
  //   fmt.Println(*p)
  //   }

  // a := 10
  // var p *int = &a
  // fmt.Println(&a)
  // fmt.Printf("%d\n",p)

  //do while
  // i := 0
  // for{
  //   i ++
  //   if i > 10{
  //     break
  //   }
  // }
  // fmt.Println(i)

  //while do
  // i := 0
  // for i<10 {
  //   i++
  // }
  // fmt.Println(i)

  //for
  //switch
  // a := 1
  // switch a {
  // case 0:
  //   fmt.Println("a等于0")
  // case 1:
  //   fmt.Println("a等于1")
  // default:
  //   fmt.Println("no")
  //
  // }

  // switch a:=1;{ //局部变量a
  // case a>=0:
  //   fmt.Println("0")
  //   fallthrough
  // case a>= 1:
  //   fmt.Println("1")
  //   // fallthrough
  // default:
  //   fmt.Println("no")
  // }

  //break

    for{
      for i :=0; i< 10;i++{
        if i>3{
          goto LABEL1 //会直接跳出一级循环 goto continue
        }
      }
    }
    LABEL2:
    for i:=0; i<5; i++{

      for{
          fmt.Println(i)
        goto LABEL2
      }

    }

    LABEL1:
    fmt.Println("ok")

}
