package cycletimer

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Timer(t *testing.T) {
	go Start(10)
	time.Sleep(time.Duration(1) * time.Second)
	var tcList []CycleTicker
	for i := 0; i < 10; i++ {
		tc := NewTicker()
		tcList = append(tcList, tc)
		go tickerTest(tc, i)
		if i == 4 {
			Close(tcList[1])
			Close(tcList[2])
		}
	}
	time.Sleep(time.Duration(120) * time.Second)
}

func tickerTest(tc CycleTicker, index int) {
	res, ok := <-tc.C
	if !ok {
		return
	}
	fmt.Println("c ", res, " "+strconv.Itoa(index))
}
