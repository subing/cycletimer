# cycle-timer
a timer cycle queue to resolve application whitch need too much timer

## use
```bash
go get github.com/subing/cycletimer
```

## example
```go
import (
    timer "github.com/subing/cycletimer"
)
c := timer.NewTicker()
_, ok := <-c
if !ok {
	fmt.Println("close c")
	return
}
```