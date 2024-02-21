package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type DBConnection struct {
	index int64
}

func main() {
	var count int64
	connectionPool := &sync.Pool{
		New: func() any {
			atomic.AddInt64(&count, 1)
			return &DBConnection{
				index: count,
			}
		},
	}

	for i := 0; i < 10; i++ {
		go func() {
			dbConn := connectionPool.Get().(*DBConnection)
			fmt.Printf("db connection: %v\n", dbConn)
			connectionPool.Put(dbConn)
		}()
	}

	fmt.Scanln()
}
