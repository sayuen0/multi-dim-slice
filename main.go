package main

import "fmt"

func main() {
	fmt.Println(to2D(makeSlice(12), 0))
	fmt.Println(to2D(makeSlice(30), 6))
	fmt.Println(to2D(makeSlice(40), 6))
	fmt.Println(to2D(makeSlice(37), 8))
}

func makeSlice(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return res
}

func to2D(s []int, n int) [][]int {
	res := make([][]int, 0)
	if len(s) == 0 || n == 0 {
		return res
	}
	c := len(s) / n
	loop := c
	if len(s)%n != 0 {
		loop += 1
	}

	for i := 0; i < loop; i++ {
		tmp := make([]int, 0)
		for j := 0; j < n; j++ {
			idx := i*n + j
			if idx == len(s) {
				break
			}
			tmp = append(tmp, s[idx])
		}
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
	}
	return res
}
