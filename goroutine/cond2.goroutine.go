package main

import (
	"fmt"
	"sync"
	"time"
)

// 버튼이 있는 GUI 애플리케이션을 만들고 있다고 가정해보자.
// 버튼을 클릭했을 때 실행되는 임의의 수의 함수를 등록하고자 한다.
// Cond는 Broadcast 메서드를 사용해 등록된 모든 핸들러에 통지할 수 있기 때문에
// 이 상황에 완벽하게 부합한다.

func main() {
	type Button struct {
		// Clicked 조건 정의
		Clicked *sync.Cond
	}
	button := Button{sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, eventName string, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			fmt.Printf("%v 이벤트 등록 완료\n", eventName)
			// 해당 고루틴을 일시 중단시킨다.
			c.Wait()
			fmt.Printf("%v 이벤트 발생\n", eventName)
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)

	subscribe(button.Clicked, "window", func() {
		defer clickRegistered.Done()
		fmt.Println("Maximizing window.")
	})

	subscribe(button.Clicked, "dialog", func() {
		defer clickRegistered.Done()
		fmt.Println("Displaying annoying dialog box!")
	})

	subscribe(button.Clicked, "mouse", func() {
		defer clickRegistered.Done()
		fmt.Println("Mouse clicked.")
	})

	time.Sleep(time.Second * 5)
	button.Clicked.Broadcast() // 모두 깨우기
	// 하나씩 깨우기
	// button.Clicked.Signal()
	// button.Clicked.Signal()
	// button.Clicked.Signal()

	clickRegistered.Wait()
}
