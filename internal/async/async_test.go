package async

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// go test ./... -race

// ! race detected during execution of test !
// func Test_Race(t *testing.T) {
// 	var c = 0

// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			c += i
// 		}()
// 	}

// 	time.Sleep(time.Second)

// 	assert.Equal(t, 45, c)
// }

func Test_Atomic(t *testing.T) {
	var c int32 = 0

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			atomic.AddInt32(&c, int32(i))
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(45), c)
}

func Test_Mutex(t *testing.T) {
	var c int32 = 0

	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mx.Lock()
			c += int32(i)
			mx.Unlock()
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(45), c)
}

func Test_RWMutex(t *testing.T) {
	var c int32 = 0

	wg := sync.WaitGroup{}
	mx := sync.RWMutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mx.Lock()
			c += int32(i)
			mx.Unlock()
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(45), c)
}

func Test_Once(t *testing.T) {
	var c int32 = 0

	wg := sync.WaitGroup{}
	once := sync.Once{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			once.Do(func() {
				c += 100 // executed just once
				for i := 0; i < 10; i++ {
					wg.Done()
				}
			})
		}()
	}

	wg.Wait()

	assert.Equal(t, int32(100), c)
}

func Test_Cond(t *testing.T) {
	var (
		mu    sync.Mutex
		cond  = sync.NewCond(&mu)
		queue []func()
	)

	// Workers
	for i := 0; i < 3; i++ {
		go func(id int) {
			for {
				mu.Lock()
				for len(queue) == 0 {
					cond.Wait()
				}
				job := queue[0]
				queue = queue[1:]
				mu.Unlock()

				fmt.Printf("Worker %d processing job\n", id)
				job()
			}
		}(i)
	}

	// Producer
	for i := 0; i < 10; i++ {
		id := i
		mu.Lock()
		queue = append(queue, func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Job %d done\n", id)
		})
		cond.Signal() // Wake one worker
		mu.Unlock()
	}
	time.Sleep(2 * time.Second)
}
