package math

type Math interface {
	Add(int, int) int
}

type MathImpl struct{}

func (MathImpl) Add(a, b int) int {
	return a + b
}
