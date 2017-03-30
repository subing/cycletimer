//Package cycletimer 批量定时器
package cycletimer

import (
	"fmt"
	"time"

	set "github.com/deckarep/golang-set"
)

var _interval int
var _cycleSlice []set.Set
var _current int
var timer *time.Ticker

//Start 启动定时器
func Start(interval int) {
	_interval = interval
	if _interval == 0 {
		_interval = 10
	}
	setInit(_interval)
	timer = time.NewTicker(time.Duration(1) * time.Second)
	for {
		_, ok := <-timer.C
		if !ok {
			fmt.Println("timer chan is closed")
			timer.Stop()
			break
		}
		go checkTimeout(_current)
		_current++
		if _current == _interval+1 {
			_current = 0
		}
	}
}

//NewTicker 创建定时器
func NewTicker() chan string {
	c := make(chan string)
	putIndex := 0
	if _current == 0 {
		putIndex = _interval
	} else {
		putIndex = _current - 1
	}
	if _cycleSlice[putIndex] == nil {
		_cycleSlice[putIndex] = set.NewSet()
	}
	_cycleSlice[putIndex].Add(c)
	return c
}

//Stop 结束定时器
func Stop() {
	timer.Stop()
}

func checkTimeout(index int) {
	if _cycleSlice[index] == nil {
		_cycleSlice[index] = set.NewThreadUnsafeSet()
	}
	if _cycleSlice[index].Cardinality() > 0 {
		for c := range _cycleSlice[index].Iter() {
			tmp := c.(chan string)
			tmp <- "time_out"
			close(tmp)
			_cycleSlice[index].Remove(tmp)
		}
	}
}
func setInit(interval int) {
	_cycleSlice = make([]set.Set, interval+1)
}
