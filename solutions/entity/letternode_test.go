package entity

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestNewLetterNodeMap(t *testing.T) {
	m := NewLetterNodeMap([]string{"HELLO", "OMLMB", "QOLOD", "Z"},
		"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	for k, v := range m {
		switch k {
		case "HELLO":
			for _, n := range v {
				switch s := n.ToLetters(len(k)).ReadableString(); s {
				case "H[0 0] E[1 1] L[2 2] L[3 3] O[2 3]":
				case "H[0 0] E[1 1] L[2 2] L[3 3] O[4 4]":
				default:
					assert.FailNow(t, s)
				}
			}
		case "OMLMB":
			for _, n := range v {
				switch s := n.ToLetters(len(k)).ReadableString(); s {
				case "O[2 1] M[1 2] L[2 2] M[3 2] B[4 2]":
				case "O[2 1] M[3 2] L[2 2] M[3 2] B[4 2]":
				case "O[2 1] M[3 2] L[3 3] M[3 2] B[4 2]":
				case "O[0 2] M[1 2] L[2 2] M[3 2] B[4 2]":
				case "O[2 3] M[1 2] L[2 2] M[3 2] B[4 2]":
				case "O[2 3] M[3 2] L[2 2] M[3 2] B[4 2]":
				case "O[2 3] M[3 2] L[3 3] M[3 2] B[4 2]":
				default:
					assert.FailNow(t, s)
				}
			}
		case "QOLOD":
			for _, n := range v {
				switch s := n.ToLetters(len(k)).ReadableString(); s {
				case "Q[2 0] O[2 1] L[2 2] O[2 3] D[2 4]":
				case "Q[4 3] O[4 4] L[3 3] O[2 3] D[2 4]":
				default:
					assert.FailNow(t, s)
				}
			}
		case "Z":
			for _, n := range v {
				switch s := n.ToLetters(len(k)).ReadableString(); s {
				case "Z[3 0]":
				default:
					assert.FailNow(t, s)
				}
				assert.Empty(t, n.ToLetters(2))
			}
		default:
			assert.FailNowf(t, "'%v'does not exist", k)
		}
	}

	assert.Empty(t, NewLetterNodeMap([]string{},
		"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5))
	assert.Empty(t, NewLetterNodeMap([]string{"Z"},
		"", 5, 5))
}
