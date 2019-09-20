package redis

import (
	"fmt"
	"testing"
	"time"
)

func TestTryLock(t *testing.T) {
	for i := 0; i < 50; i++ {
		go func() {
			//TryLock("mc", 2)
			sum()
		}()
	}

	time.Sleep(time.Second * 4)
	fmt.Println(su)
}

var su = 0

func sum() {
	su += 1
}
