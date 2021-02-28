package main

import (
	"fmt"
)

// Pet struct contains a field name
type Pet struct {
	name string
}

// Play method is the memder of Pet and
// it sperified Play action of pet
func (p *Pet) Play() {
	fmt.Println(p.Speak())
}

// Speak method is the memder of Pet and
// it sperified Speak action of Pet
func (p *Pet) Speak() string {
	return fmt.Sprintf("my name is %v", p.name)
}

// Name , returns the Pet name
func (p *Pet) Name() string {
	return p.name
}

// Dog struct inheritance Pet
type Dog struct {
	Pet   // Embedded Pet
	Breed string
}

// Speak , Dog overrides the Speak() method
func (d *Dog) Speak() string {
	return fmt.Sprintf("%v and I am a %v", d.Pet.Speak(), d.Breed)
}

func main() {
	d := Dog{Pet: Pet{name: "spot"}, Breed: "pointer"}
	fmt.Println(d.Name())
	fmt.Println(d.Speak())
	d.Play()
}
