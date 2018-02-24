package main

import "fmt"

func main() {
	ptrSon := NewSon()
	ptrSon.CallPrint()

	ptrDaughter := NewDaughter()
	ptrDaughter.Print()

	ptrMan := NewMan()
	ptrMan.CallPrint()

}

type BaseInterface interface {
	CallPrint()
	Print()
}
type Parent struct {
	familyName string
}

func (this *Parent) Init() {
	this.familyName = "July"
}

func (this *Parent) Print() {
	this.Change()
	fmt.Println("parent + ", this.familyName)
}
func (this *Parent) Change() {
	this.familyName = this.familyName + "parent"
}
type Son struct {
	Parent
	base BaseInterface
}

func NewSon() *Son {
	instance := &Son{}
	instance.base = instance
	instance.Init()
	return instance
}
func (this *Son) Print() {
	fmt.Println("Son + ", this.familyName)
}

func (this *Son) CallPrint() {
	this.base.Print()
}
type Daughter struct {
	Parent
}

func NewDaughter() *Daughter {
	instance := &Daughter{}
	instance.Init()
	return instance
}

//func (this *Daughter) Print() {
//	fmt.Println("Daughter + ", this.familyName)
//}
func (this *Daughter) Change()  {
	this.familyName = this.familyName + "daughter"
}

type Human struct {
	base BaseInterface
	age uint32
}

func NewHuman() *Human {
	instance := &Human{}
	instance.base = instance
	return instance
}
func (this *Human) CallPrint()  {
	this.base.Print()
}
func (this *Human) Print()  {
	fmt.Println("human")
}

type Man struct {
	*Human
	area uint32
}

func NewMan() *Man {
	instance := &Man{
		Human:NewHuman(),
	}
	instance.base = instance
	return instance
}
//func (this *Man) CallPrint()  {
//	this.base.Print()
//}
func (this *Man) Print()  {
	fmt.Println("man")
}

