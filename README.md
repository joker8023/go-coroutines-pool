# go-coroutines-pool

pool for go coroutines

## Table of Contents

- [Requirements](#requirements)
- [Usage](#usage)
- [define](#define)

## Requirements

requires the following to run:

- go ^1.12.6

## Usage

```
   package main

import (
	"log"
	"runtime"
	"time"

	pool "github.com/joker8023/go-coroutines-pool"
)

func main() {
	pool := pool.New(3)
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

```

Build

---

```
   go build
```
