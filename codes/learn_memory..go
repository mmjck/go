package main

func learnMemory() (p, q *int) {
	p = new(int)

	s := make([]int, 20)

	s[3] = -7
	r := -2

	return &s[3], &r
}
