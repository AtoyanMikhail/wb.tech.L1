package main

import "time"

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	print("Привет, ")
	sleep(3*time.Second)
	println("мир!")
}