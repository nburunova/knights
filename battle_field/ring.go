package battle_field

import (
	"fmt"
)

type Character interface {
	CharacterType() CharacterType
	IsDead() bool
	DamagePoints(enemy CharacterType) int
	ReceiveDamage(damage int)
}

//go:generate mockgen --source=./ring.go --destination=./testdata/ring.go

type Member struct {
	Character
	ID   int
	next *Member
	prev *Member
}

type Ring struct {
	head *Member
}

func NewRing(characters ...Character) *Ring {
	var first, prev *Member
	for i, character := range characters {
		if first == nil {
			first = &Member{
				Character: characters[0],
				ID:        i + 1,
			}
			prev = first
			continue
		}
		prev.next = &Member{
			Character: character,
			ID:        i + 1,
			prev:      prev,
		}
		prev = prev.next
	}
	prev.next = first
	first.prev = prev
	return &Ring{
		head: first,
	}
}

func (r *Ring) Head() *Member {
	return r.head
}

func (r *Ring) Bury(mem *Member) {
	fmt.Printf("Bury %v\n", mem.Name())
	if r.head == mem {
		if mem.Len() == 1 {
			r.head = nil
			return
		}
		r.head = r.head.next
	}
	for p := r.head.next; p != r.head; p = p.next {
		if p != mem {
			continue
		}
		p.prev.next = p.next
		p.next.prev = p.prev
		return
	}
}

func (r *Member) Len() int {
	n := 0
	if r != nil {
		n = 1
		for p := r.next; p != r; p = p.next {
			n++
		}
	}
	return n
}

func (r *Member) Move(steps int) *Member {
	cur := r
	for i := 0; i < steps; i++ {
		cur = cur.next
	}
	return cur
}

func (r *Member) Name() string {
	return fmt.Sprintf("%v %v", r.CharacterType(), r.ID)
}

func (r *Member) Attack(enemy *Member) {
	damage := r.DamagePoints(enemy.CharacterType())
	fmt.Printf("%v is active. %v attacks %v with damage %v \n", r.Name(), r.Name(), enemy.Name(), damage)
	enemy.ReceiveDamage(damage)
}
