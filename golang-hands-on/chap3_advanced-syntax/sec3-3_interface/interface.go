package main

import "fmt"

type General interface{}

type Data interface {
	Initialize(name string, age int)
	PrintData()
	SetValue(vals map[string]General)
}

type MyData struct {
	Name string
	Age  int
}

func (d *MyData) Initialize(name string, age int) {
	d.Name = name
	d.Age = age
}

func (d MyData) PrintData() {
	fmt.Println("Name: ", d.Name)
	fmt.Println("Age:  ", d.Age)
}

func (d *MyData) SetValue(vals map[string]General) {
	d.Name = vals["name"].(string)
	d.Age = vals["age"].(int)
}

type YourData struct {
	Name string
	Age  int
	Mail string
}

func (d *YourData) Initialize(name string, age int) {
	d.Name = name
	d.Age = age
}

func (d *YourData) PrintData() {
	fmt.Println("Name: ", d.Name)
	fmt.Println("Age:  ", d.Age)
	fmt.Println("Mail: ", d.Mail)
}

func (d *YourData) SetValue(vals map[string]General) {
	d.Name = vals["name"].(string)
	d.Age = vals["age"].(int)
	d.Mail = vals["mail"].(string)
}

func main() {
	var me MyData = MyData{}
	me.Initialize("Taro", 18)

	var you Data = new(YourData)
	you.SetValue(map[string]General{
		"name": "Suzuki",
		"age":  32,
		"mail": "suzuki@example.com",
	})

	a := []Data{&me, you}
	for _, d := range a {
		d.PrintData()
	}
}
