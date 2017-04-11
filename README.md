# golangStudy #

主要用于对go感兴趣的新手相互交流

目前由于本人边工作边学习，所以该版本非常紊乱，所以请随意看下即可，后期我会推出一个系统化适合新手学习的版本啦~

# go env #

## -、各个参数的意义

<table class="table table-bordered table-striped table-condensed">
   <tr>
      <th>环境变量</th>
      <th>意义</th>
   </tr>
   <tr>
      <td>GOROOT</td>
      <td>go的安装路径,各系统默认安装路径会有不同遂不再赘述,若要执行go命令和go工具,需要设置go的可执行文件路径 $PATH:$GOOT/bin</td>
   </tr>
   <tr>
      <td>GOPATH</td>
      <td>go install和go get的工具等会用到GOPATH环境变量.GOPATH是作为编译后的二进制的存放目的地和import包时的搜索路径(其实也是你的工作目录,你可以在src下创建你自己的go源文件)</td>
</table>

### GOPATH注解

1. GOPATH之下主要有三个目录:bin、pkg、src
2. bin目录主要存放可执行文件;pkg目录存放编译好的库文件主要是*.a文件;src目录主要存放go的源文件

ps:切记不要把GOPATH设置成go的安装路径,如可以自己创建一个目录devalDev

```export $GOPATH=/Users/newband/Documents/GOWORKSPACE/golangStudy```

## go新手踩坑持续更新

  ###1.同级目录下不可以定义不同的package(这条坑的我好惨~)

## 后续未完~
