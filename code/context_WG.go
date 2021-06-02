
package main

import (
	"context"
	"log"
	"sync"
	"time"
	
)

func f1(ctx context.Context) {

	//done := make(chan bool)

	ctxTO, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go a(ctxTO, &wg)
	go b(ctxTO, &wg)

	//<-done
	//<-done
	//time.Sleep(time.Second)
	//close(done)

	wg.Wait()

}

func a(ctxTO context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctxTO.Done():
			log.Println("Func A timeout")
			//done <- true
			return
		default:
			time.Sleep(time.Millisecond)	
		}
	}
}

func b(ctxTO context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctxTO.Done():
			log.Println("Func B timeout")
			//done <- true
			return
		default:
			time.Sleep(time.Millisecond)	
		}
	}
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	f1(ctx)

	time.Sleep(time.Second*2)
}
