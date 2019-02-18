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
	s.Person=new(Person)
	s.id=103
	s.name="王五"
	s.age=19
	s.score=89


	//var p Police
	//p=Police{Person{123,"222",20},"sdfsdf"}

	var p Police
	p=Police{Person{101,"警察A",25},"初级警员","85758"}
	fmt.Println(p)

	s.id=123
	s.name="123"
	s.age=20
	//s.code="89183912"
	//fmt.Println(s.id,s.name,s.age,s.score)
	fmt.Println(p.id,p.name,p.age,p.code)
}