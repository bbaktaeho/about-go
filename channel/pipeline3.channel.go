package main

import "log"

func sliceToChanne(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()

	return out
}

func main() {
	nums := []int{2, 3, 4, 7, 1}
	dataChan := sliceToChanne(nums)
	resultChan := sq(dataChan)
	for result := range resultChan {
		log.Println(result)
	}
}
