package solutions

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestSolution3(t *testing.T) {
	r := Solution3([]string{"HELLO", "OMLMB", "QOLOD"},
		"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	for k, v := range r {
		switch k {
		case "HELLO":
			assert.Equal(t, 2, len(v))
			assert.Equal(t, "H[0 0] E[1 1] L[2 2] L[3 3] O[2 3]", v[0].ReadableString())
			assert.Equal(t, "H[0 0] E[1 1] L[2 2] L[3 3] O[4 4]", v[1].ReadableString())
		case "OMLMB":
			assert.Equal(t, 7, len(v))
			assert.Equal(t, "O[2 1] M[1 2] L[2 2] M[3 2] B[4 2]", v[0].ReadableString())
			assert.Equal(t, "O[2 1] M[3 2] L[2 2] M[3 2] B[4 2]", v[1].ReadableString())
			assert.Equal(t, "O[2 1] M[3 2] L[3 3] M[3 2] B[4 2]", v[2].ReadableString())
			assert.Equal(t, "O[0 2] M[1 2] L[2 2] M[3 2] B[4 2]", v[3].ReadableString())
			assert.Equal(t, "O[2 3] M[1 2] L[2 2] M[3 2] B[4 2]", v[4].ReadableString())
			assert.Equal(t, "O[2 3] M[3 2] L[2 2] M[3 2] B[4 2]", v[5].ReadableString())
			assert.Equal(t, "O[2 3] M[3 2] L[3 3] M[3 2] B[4 2]", v[6].ReadableString())
		case "QOLOD":
			assert.Equal(t, 2, len(v))
			assert.Equal(t, "Q[2 0] O[2 1] L[2 2] O[2 3] D[2 4]", v[0].ReadableString())
			assert.Equal(t, "Q[4 3] O[4 4] L[3 3] O[2 3] D[2 4]", v[1].ReadableString())
		default:
			assert.FailNowf(t, "'%v'does not exist", k)
		}
	}
}

func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution3([]string{"HELLO", "OMLMB", "QOLOD"},
			"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	}
}
