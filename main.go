package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		go producer(fmt.Sprintf("Channel: %d", i+1))
	}
	consumer()
}
