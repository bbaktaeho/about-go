package main

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 조건 변수 생성
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for range time.Tick(time.Millisecond) {
			// tick 마다 모든 조건 변수를 깨움
			cadence.Broadcast()
		}
	}()

	takeStep := func() {
		cadence.L.Lock()
		fmt.Println("takeStep lock.")
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName) // 로그에 방향 기록
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". Success!")
			return true
		}

		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name)

		// 최대 5번 이동할 수 있음
		// 왼쪽이나 오른쪽 중 하나만 통과하면 종료
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name)
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Taeho")
	go walk(&peopleInHallway, "Mark")
	peopleInHallway.Wait()
}
