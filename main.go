//EXAMPLE 1.

// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Worker %d: Context canceled with error: %s, Exiting...\n", id, context.Cause(ctx))
// 			return
// 		default:
// 			fmt.Printf("Worker %d: Working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	// root context.
// 	// background context.
// 	rootCtx := context.Background()

// 	// new context, with cancel
// 	// pass a parent context
// 	// returns context & cancel
// 	ctx, cancel := context.WithCancel(rootCtx)

// 	// passing new context to goroutines.
// 	go worker(ctx, 1)
// 	go worker(ctx, 2)

// 	go func(canc context.CancelFunc) {
// 		time.Sleep(4 * time.Second)
// 		// call cancel func
// 		// sends message to worker
// 		canc()
// 	}(cancel)

// 	fmt.Scanln()
// }

// EXAMPLE 2
// package main

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			// should return error that canceled context
//      // helpful for none running application
// 			fmt.Printf("Worker %d: Context canceled with error: %s, Exiting...\n", id, context.Cause(ctx))
// 			return
// 		default:
// 			fmt.Printf("Worker %d: Working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	rootCtx := context.Background()

// 	// with cancel cause
// 	ctx, cancel := context.WithCancelCause(rootCtx)

// 	go worker(ctx, 1)
// 	go worker(ctx, 2)

// 	go func(canc context.CancelCauseFunc) {
// 		time.Sleep(4 * time.Second)
// 		canc(errors.New("error X"))
// 	}(cancel)

// 	fmt.Scanln()
// }

// EXAMPLE 3
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: Context canceled with error: %s, Exiting...\n", id, context.Cause(ctx))
			return
		default:
			fmt.Printf("Worker %d: Working...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	rootCtx := context.Background()
	// with timeout concept.
	// second parameter: time duration.
	timeoutCtx, _ := context.WithTimeout(rootCtx, 5*time.Second)

	go worker(timeoutCtx, 1)
	go worker(timeoutCtx, 2)

	fmt.Scanln()
}

// EXAMPLE 4
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Worker %d: Context canceled with error: %s, Exiting...\n", id, context.Cause(ctx))
// 			return
// 		default:
// 			fmt.Printf("Worker %d: Working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	rootCtx := context.Background()

// 	timeoutCtx, _ := context.WithTimeout(rootCtx, 10*time.Second)
// 	deadlineCtx, _ := context.WithDeadline(timeoutCtx, time.Now().Add(3*time.Second))

// 	go worker(timeoutCtx, 1)
// 	go worker(deadlineCtx, 2)

// 	fmt.Scanln()
// }

// EXAMPLE 5
// package main

// import (
// 	"context"
// 	"fmt"
// 	"time"
// )

// func worker(ctx context.Context, id int) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Printf("Worker %d: Context canceled with error: %s, Exiting...\n", id, context.Cause(ctx))
// 			return
// 		default:
// 			fmt.Printf("Worker %d: Working...\n", id)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// func main() {
// 	rootCtx := context.Background()

// 	timeoutCtx, _ := context.WithTimeout(rootCtx, 10*time.Second)
// 	deadlineCtx, _ := context.WithDeadline(timeoutCtx, time.Now().Add(3*time.Second))
// 	valueCtx := context.WithValue(timeoutCtx, "key", "the Value")

// 	go worker(timeoutCtx, 1)
// 	go worker(rootCtx, 2)

// 	go func(ctx context.Context) {
// 		<-ctx.Done()
// 		fmt.Printf("anonymous function canceled with value: %s... \n", ctx.Value("key"))
// 	}(valueCtx)

// 	fmt.Scanln()
// }
