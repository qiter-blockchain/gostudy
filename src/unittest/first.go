package main

func addUpper(n int) int {
	ret := 0
	for i := 0; i < n; i++ {
		ret += i
	}
	return ret
}
