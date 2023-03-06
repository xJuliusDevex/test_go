package main

import (
	"fmt"
	"strconv"
)

type Employe struct {
	id   int
	name string
}

// constructor
func NewEmploye(id int, name string) *Employe {
	return &Employe{id: id, name: name}
}

// set
func (e *Employe) SetId(id int) {
	e.id = id
}
func (e *Employe) SetName(name string) {
	e.name = name
}

// get
func (e *Employe) GetName() string {
	return e.name
}
func (e *Employe) GetID() int {
	return e.id
}

type People struct {
	name string
	age  int
}
type Empleado struct {
	id int
}
type FulltimeEmpleado struct {
	People
	Empleado
}
type PrintInfo interface {
	GetMessage() string
}

func (a Empleado) GetMessage() string {
	data := strconv.FormatInt(int64(a.id), 10)
	return "My id:" + data
}
func (a FulltimeEmpleado) GetMessage() string {
	return "Name:" + a.name + "\nid:" + strconv.FormatInt(int64(a.id), 10)
}
func GetMessage(p PrintInfo) {
	fmt.Println(p.GetMessage())
}
func main() {
	// e := Employe{}
	// fmt.Printf("%v", e)
	// e.id = 1
	// e.name = "Julio Cesar Garcia Quispe"
	// fmt.Printf("%v", e)
	// e.SetId(2)
	// e.SetName("Julius")
	// fmt.Printf("%v", e)
	// fmt.Println(e.GetID(), e.GetName())
	// e4 := *NewEmploye(5, "Yaneth")
	// fmt.Printf("%v", e4)
	a := Empleado{}
	a.id = 1
	ftEmploye := FulltimeEmpleado{}
	ftEmploye.age = 15
	ftEmploye.id = 1
	ftEmploye.name = "Camila"
	fmt.Printf("%v", ftEmploye)
	GetMessage(a)
	GetMessage(ftEmploye)
}
