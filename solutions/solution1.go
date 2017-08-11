package solutions

import "github.com/WindomZ/interesting-solutions/solutions/entity"

// Solution1 the problem 1 solution
func Solution1(dic []string, matrixStr string, matrixX, matrixY int) entity.Result {
	res := make(entity.Result, len(dic))
	if m := entity.NewLetterNodeMap(dic, matrixStr, matrixX, matrixY); m != nil {
		for s, nodes := range m {
			for _, node := range nodes {
				letters := node.ToLetters(len(s))
				if len(letters) >= 2 {
					switch {
					case letters[0].IsHorizontal(letters[1]): // horizontal
						if letters.Pass(2, func(l1 entity.Letter, l2 entity.Letter) bool {
							return l1.IsHorizontal(l2)
						}) {
							res.Add(s, letters)
						}
					case letters[0].IsVertical(letters[1]): // vertical
						if letters.Pass(2, func(l1 entity.Letter, l2 entity.Letter) bool {
							return l1.IsVertical(l2)
						}) {
							res.Add(s, letters)
						}
					case letters[0].IsDiagonal(letters[1]): // diagonal
						if letters.Pass(2, func(l1 entity.Letter, l2 entity.Letter) bool {
							return l1.IsDiagonal(l2)
						}) {
							res.Add(s, letters)
						}
					}
				} else {
					res.Add(s, letters)
				}
			}
		}
	}
	return res
}
