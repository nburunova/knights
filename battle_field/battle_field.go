package battle_field

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type BattleField struct {
	ring       *Ring
	selectNext func(current *Member) *Member
}

func (bf *BattleField) Fight() {
	rand.Seed(time.Now().UnixNano())
	var attacker, defender, dead *Member
	attacker = bf.ring.head
	for {
		if attacker.IsDead() {
			dead = attacker
			attacker = bf.selectNext(dead)
			bf.ring.Bury(dead)
		}
		defender = bf.selectNext(attacker)
		if defender == nil {
			fmt.Printf("Winner is %v", attacker.Name())
			return
		}
		attacker.Attack(defender)
		attacker = defender
	}
}

func selectNext(member *Member) *Member {
	if member == member.next {
		return nil
	}
	return member.next
}

func selectPrev(member *Member) *Member {
	if member == member.prev {
		return nil
	}
	return member.prev
}

func selectRandom(member *Member) *Member {
	l := member.Len()
	if l == 1 {
		return nil
	}
	for {
		randMove := rand.Intn(l)
		next := member.Move(randMove)
		if next != member {
			return next
		}
	}
}

type BattleFieldBuilder struct {
	characters          []Character
	isBackwardDirection bool
	isRandom            bool
}

func NewBuilder() *BattleFieldBuilder {
	return &BattleFieldBuilder{
		characters: make([]Character, 0),
	}
}

func (b *BattleFieldBuilder) WithBackwardPolicy() *BattleFieldBuilder {
	b.isBackwardDirection = true
	return b
}

func (b *BattleFieldBuilder) WithRandomPolicy() *BattleFieldBuilder {
	b.isRandom = true
	return b
}

func (b *BattleFieldBuilder) WithKnights(number int) *BattleFieldBuilder {
	for i := 0; i < number; i++ {
		b.characters = append(b.characters, NewKnight())
	}
	return b
}

func (b *BattleFieldBuilder) WithWitches(number int) *BattleFieldBuilder {
	for i := 0; i < number; i++ {
		b.characters = append(b.characters, NewWitch())
	}
	return b
}

func (b *BattleFieldBuilder) WithCharacters(characters ...Character) *BattleFieldBuilder {
	b.characters = characters
	return b
}

func (b *BattleFieldBuilder) Build() (*BattleField, error) {
	if len(b.characters) == 0 {
		return nil, errors.New("character required")
	}
	for _, c := range b.characters {
		if c.IsDead() {
			return nil, errors.New("found dead character")
		}
	}
	if b.isRandom && b.isBackwardDirection {
		return nil, errors.New("select only one policy for choosing next character")
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(
		len(b.characters),
		func(i, j int) { b.characters[i], b.characters[j] = b.characters[j], b.characters[i] },
	)
	bf := &BattleField{
		ring:       NewRing(b.characters...),
		selectNext: selectNext,
	}
	if b.isBackwardDirection {
		bf.selectNext = selectPrev
	}
	if b.isRandom {
		bf.selectNext = selectRandom
	}
	return bf, nil
}
