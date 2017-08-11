package entity

// Result collection the results(Letters)
type Result map[string][]Letters

// Add append l with s
func (r Result) Add(s string, l Letters) {
	if len(s) == 0 || len(l) == 0 {
		return
	}
	if _, ok := r[s]; !ok {
		r[s] = make([]Letters, 0, 2)
	}
	r[s] = append(r[s], l)
}
