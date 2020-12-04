package main

const a = 1
const b = a

func A(a int) int {
	return a
}

func B(b int64) int64 {
	return b
}
func main() {
	A(a)
	B(b)
}
