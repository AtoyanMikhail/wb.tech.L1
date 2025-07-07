package main

import (
	"fmt"
	"math"
)

// Когда нужно использовать сторонний класс, но его интерфейс 
// не соответствует остальному коду приложения. адаптер позволяет создать 
// объект-прокладку, который будет превращать вызовы приложения в формат,
// понятный стороннему классу.

// Преимущество: отделяет и скрывает от клиента подробности преобразования различных интерфейсов.

// Недостаток: усложняет код программы из-за введения дополнительных классов.

// Пример: пусть есть класс, который обробатывает данные, поступаемые в формате json.
// Появляется новый источник данных, отправляющий xml. Вместо переписывания исходного класса
// можно использовать адаптер

type RoundPeg interface {
	getRadius() float64
}

type ConcreteRoundPeg struct {
	radius float64
}

func (p ConcreteRoundPeg) getRadius() float64 {
	return p.radius
}

type SquarePeg interface {
	getWidth() float64
}

type ConcreteSquarePeg struct {
	width float64
}

func (p ConcreteSquarePeg) getWidth() float64 {
	return p.width
}

type RoundHole struct {
	radius float64
}

func (h RoundHole) GetRadius() float64 {
	return h.radius
}

func (h RoundHole) fits(peg RoundPeg) bool {
	return h.radius >= peg.getRadius()
}

type PegAdapter struct { // implements RoundPeg
	peg SquarePeg
}

func (a *PegAdapter) getRadius() float64 {
	return a.peg.getWidth() * math.Sqrt(2) / 2
}

func main() {
	hole := RoundHole{5}

	var rp RoundPeg
	var sp SquarePeg

	rp = ConcreteRoundPeg{4}
	sp = ConcreteSquarePeg{7}

	fmt.Printf("Does round peg with radius %.1f suit? %v\n", rp.getRadius(), hole.fits(rp))

	adapter := &PegAdapter{peg: sp}
	fmt.Printf("Does square peg with readius %.1f suit? %v\n", sp.getWidth(),  hole.fits(adapter))
}
