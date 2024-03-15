package main

import "fmt"

type Form interface {
	Clone() Form
}

type RegistrationForm struct {
	Name  string
	Email string
}

func (r *RegistrationForm) Clone() Form {
	return &RegistrationForm{
		Name:  r.Name,
		Email: r.Email,
	}
}

func main() {
	prototype := &RegistrationForm{
		Name:  "example Name",
		Email: "example@example.com",
	}

	clonedForm := prototype.Clone()
	regForm, ok := clonedForm.(*RegistrationForm)
	if !ok {
		fmt.Println("Type assertion to RegistrationForm failed")
		return
	}

	fmt.Printf("Name: %s, Email: %s\n", regForm.Name, regForm.Email)
}