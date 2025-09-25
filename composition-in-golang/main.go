package main

import "fmt"

// type Position

///
/// COWBOY
///

type Cowboy struct {
	Health
	Armour
	Hat bool
}

func NewCowboy(health Health, armour Armour, hat bool) Cowboy {
	return Cowboy{
		Health: health,
		Armour: armour,
		Hat:    hat,
	}
}

func (c Cowboy) Fight(opponent Fightable) {
	if c.Guns {
		c.Damage += 20
		opponent.ApplyDamage(c.Damage)

	}
}

func (c Cowboy) ApplyDamage(amount int) {
	c.ApplyDamage(amount)
}

///
/// BANDIT
///

type Bandit struct {
	Health
	Armour
	Bandana bool
}

func newBandit(maxHp int) *Bandit {
	return &Bandit{
		Health: Health{
			MaxHP:     maxHp,
			CurrentHP: maxHp,
		},
		Armour: Armour{},
	}
}

func (b Bandit) Fight(opponent Fightable) {
	if b.Bandana {
		b.Damage += 40
		opponent.ApplyDamage(b.Damage)

	}
}

func (b Bandit) ApplyDamage(amount int) {
	b.ApplyDamage(amount)
}

func (b *Bandit) TakeOffBandana() {
	b.Bandana = false
}

func (b *Bandit) Move(move bool) {

}

///
/// HEALTH
///

type Health struct {
	CurrentHP int
	MaxHP     int
	SnakeOil  bool
}

func (h *Health) TakeDamage(amount int) {
	h.CurrentHP -= amount
	if h.CurrentHP < 0 {
		h.CurrentHP = 0
	}
}

///
/// ARMOUR
///

type Armour struct {
	Guns    bool
	Defense int
	Potions bool
	Damage  int
}

///
/// FIGHTING
///

// Interface is like a behaviour group

type Fightable interface {
	Fight(opponent Fightable)
	ApplyDamage(amount int)
}

type Movable interface {
	Move(move bool)
}

// Homework (21/09)
// 1. Create construction function for Bandit. Single argument of maxHp: NewBandit(maxHp)
// 2. Create two new methods called Equip and Unequip for the Armour struct (either guns, or change the defense value)
// 3. For the Cowboy, create a method called EquipHat, and this changes the Hat value AND it also adds a defensive bonus


func (a *Armour) Equip() {
	if a.Guns {
		fmt.Println("You already have guns equipped")
		return
	} else {
		a.Guns = true
	}

}

func (a *Armour) Unequip() {
	if a.Guns == false {
		fmt.Println("You already have guns unequipped")

	} else if a.Guns == true {
		a.Guns = false
	}

}

func (c *Cowboy) EquipHat() {
	if c.Hat == true {
		fmt.Println("You already have hat equipped")

	} else if c.Hat == false {
		c.Hat = true
		c.Defense += 15
	}

}

// func newCowboy(maxHP int) *Cowboy{
// 	return &Cowboy{

// 	}
// }

func main() {
	cowboy1 := Cowboy{
		Health: Health{
			CurrentHP: 90,
			MaxHP:     90,
			SnakeOil:  true,
		},
		Armour: Armour{
			Guns:    false,
			Defense: 60,
			Potions: false,
		},
		Hat: false,
	}
	cowboy1.TakeDamage(10)
	fmt.Println(cowboy1.CurrentHP)
	bandit1 := Bandit{
		Health: Health{
			MaxHP:    95,
			SnakeOil: true,
		},
		Armour: Armour{
			Guns:    true,
			Defense: 50,
			Potions: false,
		}, Bandana: false,
	}
	fmt.Println(bandit1.MaxHP)

	var people []Fightable
	people = append(people, cowboy1, bandit1)
	var moving []Movable
	moving = append(moving, &bandit1)

	fmt.Println("Before equip:", cowboy1.Guns)
	cowboy1.Equip()
	fmt.Println("After equip:", cowboy1.Guns)
	fmt.Println("Before unequip:", cowboy1.Guns)
	cowboy1.Unequip()
	fmt.Println("After unequip:", cowboy1.Guns)
	fmt.Println("Before equipHat:", cowboy1.Hat)
	fmt.Println("Defense before equipHat:", cowboy1.Defense)
	cowboy1.EquipHat()
	fmt.Println("After equiphat:", cowboy1.Hat)
	fmt.Println("Defense after equipHat:", cowboy1.Defense)
}
