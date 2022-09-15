package main

import (
	"log"
	"time"
)

func main() {
	size := 10
	newArr := make([]int, 0, size)
	arrChannel := make(chan int)

	start := time.Now()
	for i := 0; i < size; i++ {
		go func(data int) {
			time.Sleep(time.Millisecond * 100)
			arrChannel <- data
		}(i)
	}

	for data := range arrChannel {
		newArr = append(newArr, data)
		if len(newArr) == size {
			close(arrChannel)
			break
		}
	}

	end := time.Since(start)

	log.Println(end)

	log.Println(len(newArr))
	log.Println(newArr)
}
