package main

func FibTail(n int, a ...int) int {
	if len(a) == 0 {
		a = []int{1, 0}
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return a[0]
	}

	return FibTail(n-1, a[0]+a[1], a[0])
}
