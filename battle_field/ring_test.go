package battle_field_test

import (
	"knights/battle_field"
	mock_characters "knights/battle_field/testdata"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/golang/mock/gomock"
)

func TestRing_Bury(t *testing.T) {
	ctrl := gomock.NewController(t)
	char1 := mock_characters.NewMockCharacter(ctrl)
	char2 := mock_characters.NewMockCharacter(ctrl)
	char3 := mock_characters.NewMockCharacter(ctrl)
	char4 := mock_characters.NewMockCharacter(ctrl)

	r := battle_field.NewRing(char1, char2, char3, char4)
	require.Equal(t, 4, r.Head().Len())

	ch2 := r.Head().Move(1)
	require.Equal(t, 2, ch2.ID)

	r.Bury(ch2)
	require.Equal(t, 3, r.Head().Len())

	newCh2 := r.Head().Move(1)
	require.Equal(t, 3, newCh2.ID)
}

func TestRing_BuryHead(t *testing.T) {
	ctrl := gomock.NewController(t)
	char1 := mock_characters.NewMockCharacter(ctrl)
	char2 := mock_characters.NewMockCharacter(ctrl)
	char3 := mock_characters.NewMockCharacter(ctrl)

	r := battle_field.NewRing(char1, char2, char3)
	require.Equal(t, 3, r.Head().Len())

	r.Bury(r.Head())
	require.Equal(t, 2, r.Head().Len())

	require.Equal(t, 2, r.Head().ID)
}

func TestRing_BuryLast(t *testing.T) {
	ctrl := gomock.NewController(t)
	char1 := mock_characters.NewMockCharacter(ctrl)

	r := battle_field.NewRing(char1)
	require.Equal(t, 1, r.Head().Len())

	r.Bury(r.Head())
	require.Equal(t, 0, r.Head().Len())
}
