package main

import "fmt"

func main() {
	blaus := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	threads := 3
	threadsBlaus := make([][]int, threads)

	for i, v := range blaus {
		currentThread := i % threads
		threadsBlaus[currentThread] = append(threadsBlaus[currentThread], v)
	}

	fmt.Println(threadsBlaus)
}
