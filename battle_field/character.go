package battle_field

import "math/rand"

type CharacterType string

const (
	KnightType CharacterType = "Knight"
	WitchType  CharacterType = "Witch"
)

func rollDice() int {
	return rand.Intn(6) + 1
}

type Base struct {
	Health int
	Type   CharacterType
}

func (m *Base) ReceiveDamage(damage int) {
	m.Health = m.Health - damage
}

func (m *Base) DamagePoints(_ CharacterType) int {
	return rollDice()
}

func (m *Base) IsDead() bool {
	return m.Health < 0
}

func (m *Base) CharacterType() CharacterType {
	return m.Type
}

type Knight struct {
	Base
}

func NewKnight() *Knight {
	return &Knight{
		Base{
			Health: 100,
			Type:   KnightType,
		},
	}
}

type Witch struct {
	Base
}

func NewWitch() *Witch {
	return &Witch{
		Base{
			Health: 100,
			Type:   WitchType,
		},
	}
}

func (w *Witch) DamagePoints(enemyType CharacterType) int {
	switch enemyType {
	case KnightType:
		return 2 * rollDice()
	default:
		return w.Base.DamagePoints(enemyType)
	}
}
