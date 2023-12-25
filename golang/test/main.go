package main

func main() {
	x := make([]int, 3)

	x[0], x[1], x[2] = 1, 2, 3

	for i, val := range x {
		println(&x[i], "vs.", &val, val)
		val = 10
		x[i] = 10
	}

	for i, val := range x {
		println(&x[i], "vs.", &val, val)
	}
}
