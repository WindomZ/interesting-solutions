package entity

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestLetter_XY(t *testing.T) {
	assert.Equal(t, 12, Letter{X: 1, Y: 2}.XY())
}

func TestLetter_IsNeighbor(t *testing.T) {
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 1, Y: 1}))
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 1, Y: 3}))
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 0, Y: 2}))
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 2, Y: 2}))
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 0, Y: 1}))
	assert.True(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 2, Y: 3}))

	assert.False(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 1, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 0, Y: 0}))
	assert.False(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 1, Y: 4}))
	assert.False(t, Letter{X: 1, Y: 2}.IsNeighbor(Letter{X: 3, Y: 2}))
}

func TestLetter_IsHorizontal(t *testing.T) {
	assert.True(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 0, Y: 2}))
	assert.True(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 2, Y: 2}))

	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 1, Y: 1}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 1, Y: 3}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 0, Y: 1}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 2, Y: 3}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 1, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 0, Y: 0}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 1, Y: 4}))
	assert.False(t, Letter{X: 1, Y: 2}.IsHorizontal(Letter{X: 3, Y: 2}))
}

func TestLetter_IsVertical(t *testing.T) {
	assert.True(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 1, Y: 1}))
	assert.True(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 1, Y: 3}))

	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 0, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 2, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 0, Y: 1}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 2, Y: 3}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 1, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 0, Y: 0}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 1, Y: 4}))
	assert.False(t, Letter{X: 1, Y: 2}.IsVertical(Letter{X: 3, Y: 2}))
}

func TestLetter_IsDiagonal(t *testing.T) {
	assert.True(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 0, Y: 1}))
	assert.True(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 2, Y: 3}))

	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 1, Y: 1}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 1, Y: 3}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 0, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 2, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 1, Y: 2}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 0, Y: 0}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 1, Y: 4}))
	assert.False(t, Letter{X: 1, Y: 2}.IsDiagonal(Letter{X: 3, Y: 2}))
}

func TestLetters_Pass(t *testing.T) {
	l := Letters{
		{X: 0, Y: 0},
		{X: 1, Y: 1},
		{X: 2, Y: 2},
	}
	assert.False(t, l.Pass(0, func(l1 Letter, l2 Letter) bool {
		return l1.IsHorizontal(l2)
	}))
	assert.False(t, l.Pass(0, func(l1 Letter, l2 Letter) bool {
		return l1.IsVertical(l2)
	}))
	assert.True(t, l.Pass(0, func(l1 Letter, l2 Letter) bool {
		return l1.IsDiagonal(l2)
	}))
	assert.True(t, l.Pass(-1, func(l1 Letter, l2 Letter) bool {
		return l1.IsDiagonal(l2)
	}))
}
