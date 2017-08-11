package solutions

import "github.com/WindomZ/interesting-solutions/solutions/entity"

// Solution2 the problem 2 solution
func Solution2(dic []string, matrixStr string, matrixX, matrixY int) entity.Result {
	res := make(entity.Result, len(dic))
	if m := entity.NewLetterNodeMap(dic, matrixStr, matrixX, matrixY); m != nil {
		for s, nodes := range m {
			for _, node := range nodes {
				if letters := node.ToLetters(len(s)); letters != nil {
					once := make(map[int]int, len(letters))
					if letters.Pass(0, func(l1 entity.Letter, l2 entity.Letter) bool {
						xy := l1.XY() + l2.XY()
						for _, xyOther := range once {
							if xy == xyOther {
								return false
							}
						}
						once[l1.XY()] = xy
						_, ok := once[l2.XY()]
						return !ok
					}) {
						res.Add(s, letters)
					}
				}
			}
		}
	}
	return res
}
