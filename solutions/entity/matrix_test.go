package entity

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestMatrixLetterMap(t *testing.T) {
	m := NewLetterMap("HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)

	hs := m.Get('H')
	assert.Equal(t, 3, len(hs))
	assert.Equal(t, "H[0 0]", hs[0].ReadableString())
	assert.Equal(t, "H[3 1]", hs[1].ReadableString())
	assert.Equal(t, "H[0 4]", hs[2].ReadableString())

	es := m.Get('E')
	assert.Equal(t, 1, len(es))
	assert.Equal(t, "E[1 1]", es[0].ReadableString())

	ls := m.Get('L')
	assert.Equal(t, 2, len(ls))
	assert.Equal(t, "L[2 2]", ls[0].ReadableString())
	assert.Equal(t, "L[3 3]", ls[1].ReadableString())

	os := m.Get('O')
	assert.Equal(t, 4, len(os))
	assert.Equal(t, "O[2 1]", os[0].ReadableString())
	assert.Equal(t, "O[0 2]", os[1].ReadableString())
	assert.Equal(t, "O[2 3]", os[2].ReadableString())
	assert.Equal(t, "O[4 4]", os[3].ReadableString())

	assert.Empty(t, m.Get('?'))
}
