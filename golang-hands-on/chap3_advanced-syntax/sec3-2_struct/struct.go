package main

import "fmt"

type person struct { // 公開する場合は Person
	Name string
	Age  int
}

func (p person) printName() {
	fmt.Println(p.Name)
}

func (p *person) incrementAge() int {
	p.Age++
	return p.Age
}

func main() {
	var me struct {
		Name string
		Age  int
	}
	me.Name = "jun"
	me.Age = 25
	fmt.Println(me)

	someone := person{"Taro", 23}
	fmt.Println(someone)
	someone.printName()
	someone.incrementAge()
	fmt.Println("Happy birthday, "+someone.Name+" is now ", someone.Age)

	somebody := person{
		Name: "Hanako",
		Age:  17,
	}
	fmt.Println(somebody)

	pp := new(person)
	fmt.Println(pp)

	a := make([]int, 3, 5)
	fmt.Println(a)
	fmt.Println(len(a), cap(a))
}
