package main

import "fmt"

func main() {
	c := make(chan int, 3) // 용량이 3인 채널 생성, 용량을 넘어서면 deadlock
	arr := []int{1, 2, 3}

	for _, data := range arr {
		c <- data // ? 즉시 데이터를 채널에 넣음

		// select { // ? select문을 활용해서 채널에 넣음
		// case c <- data:
		// }
	}

	close(c) // 채널을 닫음

	// 채널을 닫았을 때 용량이 정해져 있는 채널이면 버퍼에 남은 데이터를 꺼낼 수 있음
	for data := range c {
		fmt.Println(data)
	}
}
