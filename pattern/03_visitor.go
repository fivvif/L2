package main

import "fmt"

// Element - интерфейс элемента, который может быть посещен
type Element interface {
	Accept(visitor Visitor)
}

// ConcreteElementA - конкретная реализация элемента A
type ConcreteElementA struct {
}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(e)
}

// ConcreteElementB - конкретная реализация элемента B
type ConcreteElementB struct {
}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(e)
}

// Visitor - интерфейс посетителя
type Visitor interface {
	VisitConcreteElementA(element *ConcreteElementA)
	VisitConcreteElementB(element *ConcreteElementB)
}

// ConcreteVisitor - конкретная реализация посетителя
type ConcreteVisitor struct {
}

func (v *ConcreteVisitor) VisitConcreteElementA(element *ConcreteElementA) {
	fmt.Println("Visitor is processing ConcreteElementA")
}

func (v *ConcreteVisitor) VisitConcreteElementB(element *ConcreteElementB) {
	fmt.Println("Visitor is processing ConcreteElementB")
}

// ObjectStructure - структура объектов, над которыми выполняется операция
type ObjectStructure struct {
	elements []Element
}

func (os *ObjectStructure) Attach(element Element) {
	os.elements = append(os.elements, element)
}

func (os *ObjectStructure) Accept(visitor Visitor) {
	for _, element := range os.elements {
		element.Accept(visitor)
	}
}

func main() {
	objectStructure := ObjectStructure{}

	elementA := ConcreteElementA{}
	elementB := ConcreteElementB{}

	objectStructure.Attach(&elementA)
	objectStructure.Attach(&elementB)

	visitor := ConcreteVisitor{}

	objectStructure.Accept(&visitor)
}
