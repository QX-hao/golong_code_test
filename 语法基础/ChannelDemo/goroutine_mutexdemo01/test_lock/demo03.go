package test_lock

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

// 声明

// 加写锁、解写锁
// func (rw *RWMutex) Lock()
// func (rw *RWMutex) Unlock()

// 加读锁、解读锁
// func (rw *RWMutex) RLock()
// func (rw *RWMutex) RUnlock()

var wg3 sync.WaitGroup

var rw sync.RWMutex


// 实现可以多人读  但是只能一人写
func write(count int)  {
	defer wg3.Done()
	rw.Lock()
	defer rw.Unlock()
	fmt.Println("goroutine",count,"写操作>>>>>>>")
	// time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	time.Sleep(time.Second * 2)
}

func read(count int)  {
	defer wg3.Done()
	fmt.Println("goroutine",count,"<<<<<<<<<读操作")
	// time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	time.Sleep(time.Second *2 )

}

func Demo03()  {
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go write(i)
	}

	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go read(i)
	}

	wg3.Wait()
}
