package main

import "fmt"

// Product - конечный продукт, который мы строим
type Product struct {
	parts []string
}

func (p *Product) AddPart(part string) {
	p.parts = append(p.parts, part)
}

func (p *Product) Show() {
	fmt.Println("Product Parts:", p.parts)
}

// Builder - интерфейс для создания продукта
type Builder interface {
	BuildPart1()
	BuildPart2()
	BuildPart3()
	GetResult() *Product
}

// ConcreteBuilder - конкретная реализация интерфейса Builder
type ConcreteBuilder struct {
	product *Product
}

func NewConcreteBuilder() *ConcreteBuilder {
	return &ConcreteBuilder{product: &Product{}}
}

func (b *ConcreteBuilder) BuildPart1() {
	b.product.AddPart("Part 1")
}

func (b *ConcreteBuilder) BuildPart2() {
	b.product.AddPart("Part 2")
}

func (b *ConcreteBuilder) BuildPart3() {
	b.product.AddPart("Part 3")
}

func (b *ConcreteBuilder) GetResult() *Product {
	return b.product
}

// Director - директор, который управляет процессом строительства
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
	d.builder.BuildPart3()
}

func main() {
	builder := NewConcreteBuilder()
	director := NewDirector(builder)
	director.Construct()
	product := builder.GetResult()

	product.Show()
}
