package entity

type (
	// LetterNode chain structure with Letter
	LetterNode struct {
		Letter
		Next *LetterNode
	}
	// LetterNodes the slice of LetterNode
	LetterNodes []*LetterNode
	// LetterNodeMap the map of LetterNodes
	LetterNodeMap map[string]LetterNodes
)

// ToLetters get Letters from LetterNode
func (l *LetterNode) ToLetters(size int) Letters {
	res := make(Letters, size)
	i := size - 1
	for n := l; i >= 0 && n != nil; n = n.Next {
		res[i] = n.Letter
		i--
	}
	if i != -1 {
		return nil
	}
	return res
}

// NewLetterNodeMap creates a LetterNodeMap instance
func NewLetterNodeMap(dic []string, matrixStr string, matrixX, matrixY int) LetterNodeMap {
	if len(dic) == 0 {
		return nil
	}

	letterMap := NewLetterMap(matrixStr, matrixX, matrixY)
	if letterMap == nil {
		return nil
	}

	res := make(LetterNodeMap, len(dic))
	for _, d := range dic {
		if nodes := res.findLetterNodes(d, letterMap); len(nodes) != 0 {
			res[d] = nodes
		}
	}
	return res
}

func (l LetterNodeMap) findLetterNodes(s string, m LetterMap) (r LetterNodes) {
	l.nextLetterNodes(&r, nil, m, s, 0)
	return
}

func (l LetterNodeMap) nextLetterNodes(r *LetterNodes, n *LetterNode, m LetterMap, s string, i int) {
	if i == len(s) {
		if n != nil {
			*r = append(*r, n)
		}
		return
	}
	for _, letter := range m.Get(s[i]) {
		if n == nil || n.IsNeighbor(letter) {
			l.nextLetterNodes(r, &LetterNode{Letter: letter, Next: n}, m, s, i+1)
		}
	}
}
