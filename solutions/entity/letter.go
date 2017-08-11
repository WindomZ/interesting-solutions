package entity

import (
	"bytes"
	"fmt"
)

// Letter the letter of word in the matrix
type Letter struct {
	X int
	Y int
	V byte
}

// XY returns integer, X*10+Y
func (l Letter) XY() int {
	return l.X*10 + l.Y
}

// ReadableString returns human readable string.
func (l Letter) ReadableString() string {
	return fmt.Sprintf("%s[%d %d]", string(l.V), l.X, l.Y)
}

// IsNeighbor returns true if other in all 8 directions
func (l Letter) IsNeighbor(other Letter) bool {
	if l.X == other.X && l.Y == other.Y {
		return false
	}
	return abs(l.X-other.X) <= 1 && abs(l.Y-other.Y) <= 1
}

// IsHorizontal returns true if other in all horizontal directions
func (l Letter) IsHorizontal(other Letter) bool {
	if l.X == other.X && l.Y == other.Y {
		return false
	}
	return abs(l.X-other.X) <= 1 && abs(l.Y-other.Y) == 0
}

// IsVertical returns true if other in all vertical directions
func (l Letter) IsVertical(other Letter) bool {
	if l.X == other.X && l.Y == other.Y {
		return false
	}
	return abs(l.X-other.X) == 0 && abs(l.Y-other.Y) <= 1
}

// IsDiagonal returns true if other in all diagonal directions
func (l Letter) IsDiagonal(other Letter) bool {
	if l.X == other.X && l.Y == other.Y {
		return false
	}
	return abs(l.X-other.X) == 1 && abs(l.Y-other.Y) == 1
}

// Letters the slice of Letter
type Letters []Letter

// Pass returns true if pass all f in the traversal
func (l Letters) Pass(start int, f func(Letter, Letter) bool) bool {
	if f != nil {
		if start < 0 {
			start = 0
		}
		for i := len(l) - 1; i > start; i-- {
			if !f(l[i], l[i-1]) {
				return false
			}
		}
	}
	return true
}

// ReadableString returns human readable string.
func (l Letters) ReadableString() string {
	buf := bytes.Buffer{}
	for _, letter := range l {
		buf.WriteString(letter.ReadableString())
		buf.WriteString(" ")
	}
	return buf.String()[:buf.Len()-1]
}

// LetterMap the map of Letters
type LetterMap map[byte]Letters

// Get returns Letters by b
func (l LetterMap) Get(b byte) Letters {
	if v, ok := l[b]; ok {
		return v
	}
	return Letters{}
}
