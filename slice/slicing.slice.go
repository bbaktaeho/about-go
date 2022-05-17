package main

import "log"

func main() {
	var arr [10]int
	log.Println(arr, len(arr), cap(arr))

	// 용량은 arr크기 - 1
	slice := arr[1:3]
	log.Println(slice, len(slice), cap(slice))

	// 용량은 6 - 1
	slice2 := arr[1:3:6]
	log.Println(slice2, len(slice2), cap(slice2))
}
