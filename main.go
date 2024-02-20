package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Person struct {
	Name string
	Age  int32
}

func main() {
	// option 1
	// var count int32

	// option 2
	// var count atomic.Int32
	var waitGroup sync.WaitGroup
	var person atomic.Value

	person.Store(&Person{Name: "Nobert", Age: 28})

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			// option 1
			// atomic.AddInt32(&count, 1)

			// option 2
			// count.Add(1)

			// using values, load person into variable
			// atomic operations? pointers!
			// .(*Person), cast to person.
			altPerson := person.Load().(*Person)
			atomic.AddInt32(&altPerson.Age, 1)
		}()
	}

	waitGroup.Wait()
	// option 1
	// fmt.Println(atomic.LoadInt32(&count))

	// option 2
	// fmt.Println(count.Load())

	// use atomic operations, load atomic values & cast
	fmt.Println(person.Load().(*Person))
}
