package solutions

import "github.com/WindomZ/interesting-solutions/solutions/entity"

// Solution3 the problem 3 solution
func Solution3(dic []string, matrixStr string, matrixX, matrixY int) entity.Result {
	res := make(entity.Result, len(dic))
	if m := entity.NewLetterNodeMap(dic, matrixStr, matrixX, matrixY); m != nil {
		for s, nodes := range m {
			for _, node := range nodes {
				if letters := node.ToLetters(len(s)); letters != nil {
					res.Add(s, letters)
				}
			}
		}
	}
	return res
}
