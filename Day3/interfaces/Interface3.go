package main

import "fmt"

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() {
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {
	var d1 Describer
	p1 := Person{"Hemant", 23}
	d1 = p1
	d1.Describe()
	p2 := Person{"Aakash", 26}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Pune", "India"}

	//d2 = a

	d2 = &a
	d2.Describe()

}
