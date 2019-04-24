package main

import "time"

func main(){
	// 定义结构体
	type Employee struct {
		id int
		name string
		address string
		dob time.Time
		position string
		salary int
		managerid int
	}

	var dilbert Employee

	dilbert.address="北京"

	position:=&dilbert.position

	*position="senior"+*position
}