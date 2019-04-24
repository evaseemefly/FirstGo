package main

import "fmt"

func main() {
	// 声明map的几种方式：

	// 声明方式1
	ages := map[string]int{
		"liu":  32,
		"wang": 32,
	}

	// 声明方式2
	ages1 := make(map[string]int)
	ages1["wangwu"] = 29

	// 声明方式3
	ages2:=map[string]int{}
	ages2["zhangsan"]=19
	fmt.Println(ages["liu"])
	fmt.Println(ages1["wangwu"])
	fmt.Println(ages2["zhangsan"])
	fmt.Println("---------------")

	// 遍历map
	for name,age:=range ages{
		fmt.Printf("%s\t%d\n",name,age)
	}
}
