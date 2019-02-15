package main

import "fmt"

func main(){
	var i int=100
	var p *int
	// 将变量i的地址复制给p，获取变量i的地址使用&i
	// 注意，获取一个变量的地址使用 &i
	// 获取一个指针的对应的值使用*p
	p=&i
	fmt.Println(i)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println()
	// 修改p指针的所指向的值
	*p=120
	fmt.Println(p)
	fmt.Println(*p)
	//p*=80
}
