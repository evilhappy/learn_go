package main

import "fmt"

type Personer interface {
	Eater
	Sayer
}

type Eater interface {
	eat()
}

type Sayer interface {
	say()
}

type Man struct {
	name string
}

func (m Man) say()  {
	fmt.Println(m.name + " say something")
}

func (m Man) eat()  {
	fmt.Println(m.name + " eat something")
}

func doSomething(p Personer)  {
	p.say()
	p.eat()
}

func doAnything(e Eater) {
	fmt.Println("doAnything")
}

func main()  {
	var p Personer
	doAnything(p)
}