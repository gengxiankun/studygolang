package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type Person struct {
	Name Name
	Email []Email
}

type Name struct {
	Family string
	Personal string
}

type Email struct {
	Kind string
	Address string
}

func main() {
	person := Person{
		Name: Name{Family: "Newmarch", Personal: "Jan"},
		Email: []Email{
			Email{Kind: "work", Address: "jan@newmarch.name"},
			Email{Kind: "work", Address: "j.newmarch@boxhill.edu.au"}}}
	pj, err := json.Marshal(person)
	checkError(err)
	fmt.Println("Marshal:"pj)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatel Error: ", err.Error())
		os.Exit(1)
	}
}