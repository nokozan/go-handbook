package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name string
}

// Move implements Pet.
func (d Dog) Move() {
	panic("unimplemented")
}

// Speaker implements Speaker.

func (d Dog) Speak() {
	fmt.Println(d.Name)
}
func (d Dog) Speak2() {
	fmt.Println(d.Name + "????")
}

type Mover interface {
	Move()
}

type Cat struct {
	Name string
}

// Speaker implements Speaker.

func (d Cat) Move() {
	fmt.Println(d.Name)
}
func (d Cat) Move2() {
	fmt.Println(d.Name + "????")
}

type Pet interface {
	Mover
	Speaker
}

func main() {
	var s1 Speaker
	var d1 Dog
	d1.Name = "DOG"

	s1 = d1
	s1.Speak()

	var s2 Mover
	var d2 Cat
	d2.Name = "CAT"

	s2 = d2
	s2.Move()

	var s3 Pet = Dog{Name: "asd"}

	s3.Move()
	s3.Speak()

}
