package main

import "fmt"

// SubsystemA - первая часть подсистемы
type SubsystemA struct {
}

func (s *SubsystemA) OperationA() {
	fmt.Println("Subsystem A: Operation A")
}

// SubsystemB - вторая часть подсистемы
type SubsystemB struct {
}

func (s *SubsystemB) OperationB() {
	fmt.Println("Subsystem B: Operation B")
}

// Facade - фасад, предоставляющий унифицированный интерфейс к подсистеме
type Facade struct {
	subsystemA *SubsystemA
	subsystemB *SubsystemB
}

func NewFacade() *Facade {
	return &Facade{
		subsystemA: &SubsystemA{},
		subsystemB: &SubsystemB{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade: Operation")
	f.subsystemA.OperationA()
	f.subsystemB.OperationB()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
