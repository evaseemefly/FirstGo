package main

import "fmt"

func main(){

	var a int=10
	var b int=20

	//匿名函数
	f:=func (a int,b int)int{
		return a+b
	}
	v:=f(a,b)
	fmt.Println(v)

	f1:=func(a int,b int)int{
		return a+b
	}(a,b)
	fmt.Println(f1)
}
