package main

import "time"

var (
	data = "mydata"
	N    = 5
)

func main() {
	ch := make(chan string, 1)

	go func() {
		ch <- data
	}()

	go func() {
		<-ch
	}()
	
	// В канал, возвращаемый time.After, придёт сообщение через N секунд
	// До тех пор main горутина заблокирована 
	<-time.After(time.Duration(N) * time.Second)
}
