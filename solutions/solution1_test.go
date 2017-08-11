package solutions

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

func TestSolution1(t *testing.T) {
	r := Solution1([]string{"Z", "HELLO", "OMLMB", "QOLOD"},
		"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	for k, v := range r {
		switch k {
		case "Z":
			assert.Equal(t, 1, len(v))
			assert.Equal(t, "Z[3 0]", v[0].ReadableString())
		case "HELLO":
			assert.Equal(t, 1, len(v))
			assert.Equal(t, "H[0 0] E[1 1] L[2 2] L[3 3] O[4 4]", v[0].ReadableString())
		case "OMLMB":
			assert.Equal(t, 1, len(v))
			assert.Equal(t, "O[0 2] M[1 2] L[2 2] M[3 2] B[4 2]", v[0].ReadableString())
		case "QOLOD":
			assert.Equal(t, 1, len(v))
			assert.Equal(t, "Q[2 0] O[2 1] L[2 2] O[2 3] D[2 4]", v[0].ReadableString())
		default:
			assert.FailNowf(t, "'%v'does not exist", k)
		}
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution1([]string{"Z", "HELLO", "OMLMB", "QOLOD"},
			"HTQZAFEOHPOMLMBCSOLQHXDKO", 5, 5)
	}
}
