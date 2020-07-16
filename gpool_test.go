package pool

import (
	"log"
	"runtime"
	"testing"
	"time"
)

func Test_Example(t *testing.T) {
	pool := New(3)
	log.Println(runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		pool.Add(1)
		go func(num int) {
			time.Sleep(time.Second)
			log.Println(runtime.NumGoroutine(), num)
			pool.Done()
		}(i)
	}
	pool.Wait()
	log.Println(runtime.NumGoroutine())
}
