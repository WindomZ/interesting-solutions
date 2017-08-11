package entity

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestResult_Add(t *testing.T) {
	r := make(Result, 3)
	r.Add("", Letters{{X: 1, Y: 2}})
	r.Add("Z", Letters{})
	r.Add("Z", Letters{{X: 1, Y: 2}})
	r.Add("Z", Letters{{X: 2, Y: 3}})

	assert.Equal(t, 2, len(r["Z"]))
}
