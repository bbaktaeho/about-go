package main

import (
	"fmt"
	"sync"
	"time"
)

// 고루틴들이 대기하거나, 어떤 이벤트의 발생을 알리는 집결 지점
// 이벤트란, 두 개 이상의 고루틴 사이에서 어떤 사실이 발생했다는 사실 외에는 아
// 무런 정보를 전달하지 않는 임의의 신호를 말한다.

func main() {
	// Cond은 동시에 실행해도 안전한 방식으로 손쉽게 다른 고루틴들과의 조정이 가
	// 능하다.
	c := sync.NewCond(&sync.Mutex{})

	// 길이가 2로 고정된 큐라고 가정하자.
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		// 대기하고 있는 고루틴 하나를 깨운다. 이때 오래 기다린 고루틴을 깨운다.
		c.Signal()
	}

	// 여유가 생기면 즉시 항목을 큐에 넣기를 원하므로 큐에 여유가 있는 경우 즉시
	// 알림을 받고자 한다.
	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			// 조건에 대한 신호가 전송될 때 까지 main 고루틴은 일시 중단된다.
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		// 1초 후에 항목을 큐에서 꺼내는 새로운 고루틴을 생성한다.
		go removeFromQueue(time.Second)
		// 항목을 큐에 성공적으로 추가했으므로 조건의 임계 영역을 벗어난다.
		c.L.Unlock()
	}
}
