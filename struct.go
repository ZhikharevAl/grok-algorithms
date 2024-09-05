package main



type param struct {
	a int
	b int
}

func (p param) sum() int {
	return p.a + p.b
}