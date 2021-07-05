package main

func swap(a, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

func swapAns(a, b int) (a1, a2 int) {
	a1 = b
	a2 = a
	return
}

func main() {
	a, b := 10, 20
	swap(&a, &b)
	println(a, b)
	println(swapAns(a, b))
}
