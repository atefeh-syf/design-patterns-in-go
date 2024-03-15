package main

import (
	"fmt"
)

type Button interface {
	Render() string
}

type LightButton struct{}

func (b LightButton) Render() string {
	return "Rendering light theme button."
}

type DarkButton struct{}

func (b DarkButton) Render() string {
	return "Rendering dark theme button."
}

type UIFactory interface {
	CreateButton() Button
}

type LightThemeFactory struct{}

func (f LightThemeFactory) CreateButton() Button {
	return LightButton{}
}

type DarkThemeFactory struct{}

func (f DarkThemeFactory) CreateButton() Button {
	return DarkButton{}
}

func main() {
	var factory UIFactory
	var button Button

	userPrefersDarkMode := true 

	if userPrefersDarkMode {
		factory = DarkThemeFactory{}
	} else {
		factory = LightThemeFactory{}
	}

	button = factory.CreateButton()

	fmt.Println(button.Render())
}