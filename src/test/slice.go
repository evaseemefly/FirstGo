package main

import "fmt"

//对切片进行初始化
func InitData(num []int){
	for i:=0;i<len(num);i++{
		num[i]=i
	}
}

func InitArr(num [10]int){
	for i:=0;i<len(num);i++{
		num[i]=i
	}
}

func main(){
	//s:=[]int{10,20,30,40,50}
	//var t = s[0:3:5]
	//slice:=s[0:3:5]

	// 注意此处的区别：
	//在go中数组作为参数传递是值传递
	//切片作为参数传递是引用传递
	s2:=make([]int,10)
	var arr [10]int
	InitArr(arr)
	fmt.Println(s2)
	InitData(s2)
	//fmt.Println(t)
	fmt.Println(s2)

	fmt.Println(arr)
}

//func main(){
//	fmt.Println("hello world")
//	//PlayGame()
//}
