package cycletimer

import (
	"fmt"
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {
	go Start(10)
	for i := 0; i < 10; i++ {
		//time.Sleep(time.Duration(3) * time.Second)
		go tickerTest()
	}
	time.Sleep(time.Duration(120) * time.Second)
}

func tickerTest() {
	c := NewTicker()
	res, ok := <-c
	if !ok {
		fmt.Println("close c")
		return
	}
	fmt.Println("c ", res)
}
