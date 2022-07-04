package util

import (
	"fmt"
	"testing"
	"time"
)

func TestSchedulerAll(t *testing.T) {
	done := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				return
			case <-ticker.C:
				fmt.Println("Hello !!")
			}
		}
	}()

	// wait for 10 seconds
	time.Sleep(10 * time.Second)
	done <- true
}

func TestSelectRecv(t *testing.T) {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)
	c3 := make(chan interface{})
	close(c3)

	var c1Count, c2Count, c3Count int
	for i := 10000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-c3:
			c3Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\nc3Count: %d\n", c1Count, c2Count, c3Count)
}
