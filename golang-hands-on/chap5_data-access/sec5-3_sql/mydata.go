package main

import "fmt"

type Mydata struct {
	ID   int
	Name string
	Mail string
	Age  int
}

func (m Mydata) ToString() string {
	return fmt.Sprintf(`<"%d:%s" %s,%d>`, m.ID, m.Name, m.Mail, m.Age)
}
