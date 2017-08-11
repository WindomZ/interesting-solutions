package entity

// NewLetterMap creates a LetterMap instance
func NewLetterMap(matrixStr string, matrixX, matrixY int) LetterMap {
	if matrixX <= 0 || matrixY <= 0 || len(matrixStr) != matrixX*matrixY {
		return nil
	}
	m := make(LetterMap, matrixX*matrixY)
	for i, j := 0, 0; i < matrixY; j++ {
		b := matrixStr[i*matrixX+j]
		if _, ok := m[b]; !ok {
			m[b] = make([]Letter, 0, 2)
		}
		m[b] = append(m[b], Letter{X: j, Y: i, V: b})
		if j+1 >= matrixX {
			i++
			j = -1
		}
	}
	return m
}
