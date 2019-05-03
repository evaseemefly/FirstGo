package main

import "fmt"

func main() {
	//参考文章
	//https://www.cnblogs.com/ghj1976/archive/2013/02/28/2936595.html
	var i int
	i = 1
	//我的理解就是 *代表的取这个地址的值；而&代表的是 是这个变量是一个地址

	// 表示地址用&
	// 修改地址对应的值用*
	// 若将一个地址(&x)赋值给一个变量，eg：p=&i，那么p本身就是指向i的地址，修改p地址对应的值使用*i

	// p就是一个值，但是p对应的是一个int类型的指针
	var p *int // p 的类型是[int型的指针]
	// p等于i变量的地址
	p = &i     // p 的值为 [i的地址]
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)

	// p指针的值修改为2，由于p指向了i的地址，所以修改了p也就是修改了i
	// *p 的值为 [[i的地址]的指针] (其实就是i嘛),这行代码也就等价于 i = 2
	*p = 2
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)


	i = 3 // 验证想法
	fmt.Printf("i=%d;p=%d;*p=%d\n", i, p, *p)
}
