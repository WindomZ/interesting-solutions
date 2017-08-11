package solutions

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestSolution2(t *testing.T) {
	r := Solution2([]string{"HELLO", "OMLMB", "QOLOD", "MOMOSO"},
		"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	for k, v := range r {
		switch k {
		case "HELLO":
			assert.Equal(t, 2, len(v))
			assert.Equal(t, "H[0 0] E[1 1] L[2 2] L[3 3] O[2 3]", v[0].ReadableString())
			assert.Equal(t, "H[0 0] E[1 1] L[2 2] L[3 3] O[4 4]", v[1].ReadableString())
		case "OMLMB":
			assert.Equal(t, 3, len(v))
			assert.Equal(t, "O[2 1] M[1 2] L[2 2] M[3 2] B[4 2]", v[0].ReadableString())
			assert.Equal(t, "O[0 2] M[1 2] L[2 2] M[3 2] B[4 2]", v[1].ReadableString())
			assert.Equal(t, "O[2 3] M[1 2] L[2 2] M[3 2] B[4 2]", v[2].ReadableString())
		case "QOLOD":
			assert.Equal(t, 2, len(v))
			assert.Equal(t, "Q[2 0] O[2 1] L[2 2] O[2 3] D[2 4]", v[0].ReadableString())
			assert.Equal(t, "Q[4 3] O[4 4] L[3 3] O[2 3] D[2 4]", v[1].ReadableString())
		case "MOMOSO":
			assert.Equal(t, 3, len(v))
			assert.Equal(t, "M[1 2] O[2 1] M[3 2] O[2 3] S[1 3] O[0 2]", v[0].ReadableString())
			assert.Equal(t, "M[3 2] O[2 1] M[1 2] O[0 2] S[1 3] O[2 3]", v[1].ReadableString())
			assert.Equal(t, "M[3 2] O[2 1] M[1 2] O[2 3] S[1 3] O[0 2]", v[2].ReadableString())
		default:
			assert.FailNowf(t, "'%v'does not exist", k)
		}
	}
}

func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution2([]string{"HELLO", "OMLMB", "QOLOD", "MOMOSO"},
			"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	}
}
