package main

import "fmt"

type Person struct {
	id int
	name string
	age int
}

type Student struct {
	//指定指针类型（匿名字段）
	*Person

	score float64
}

type Police struct {
	Person

	name string
	code string // 警员编号
}

func main(){
	var s Student
	s=Student{&Person{101,"小明",19},90}
	fmt.Println(s)
	fmt.Println(s.id,s.name,s.age,s.score)

	// 创建一个警员
	var police1 Police
	police1=Police{Person{101,"警察A",25},"初级警员","85758"}

	fmt.Println(police1)
}