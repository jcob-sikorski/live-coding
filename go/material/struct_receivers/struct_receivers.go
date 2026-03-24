// The Pointer Modifier
// Create a Player struct with a Health field (int). Write a method TakeDamage(amount int).

// Challenge: Try writing it first with a value receiver, then a pointer receiver. Observe
// why the health only actually drops when using the pointer.

package main

type Player struct {
	Health int
}

func (p Player) TakeDamageValue(amount int) {
	p.Health -= amount
}

func (p *Player) TakeDamagePointer(amount int) {
	p.Health -= amount
}
