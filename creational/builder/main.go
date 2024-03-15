package main

import (
	"fmt"
)

type Car struct {
    color         string
    engineType    string
    hasSunroof    bool
    hasNavigation bool
}

type CarBuilder interface {
    SetColor(color string) CarBuilder
    SetEngineType(engineType string) CarBuilder
    SetSunroof(hasSunroof bool) CarBuilder
    SetNavigation(hasNavigation bool) CarBuilder
    Build() *Car
}

func NewCarBuilder() CarBuilder {
    return &carBuilder{
        car: &Car{}, 
    }
}

type carBuilder struct {
    car *Car
}

func (cb *carBuilder) SetColor(color string) CarBuilder {
    cb.car.color = color
    return cb
}

func (cb *carBuilder) SetEngineType(engineType string) CarBuilder {
    cb.car.engineType = engineType
    return cb
}

func (cb *carBuilder) SetSunroof(hasSunroof bool) CarBuilder {
    cb.car.hasSunroof = hasSunroof
    return cb
}

func (cb *carBuilder) SetNavigation(hasNavigation bool) CarBuilder {
    cb.car.hasNavigation = hasNavigation
    return cb
}

func (cb *carBuilder) Build() *Car {
    return cb.car
}

type Director struct {
    builder CarBuilder
}

func (d *Director) ConstructCar(color, engineType string, hasSunroof, hasNavigation bool) *Car {
    d.builder.SetColor(color).
        SetEngineType(engineType).
        SetSunroof(hasSunroof).
        SetNavigation(hasNavigation)

    return d.builder.Build()
}

func main() {
    builder := NewCarBuilder()

    director := &Director{builder: builder}
    myCar := director.ConstructCar("blue", "electric", true, true)
	fmt.Printf("%v", myCar)
}