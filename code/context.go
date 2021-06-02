
package main

import (
	"context"
	"log"
	"time"
)

func f1(ctx context.Context) {

	done := make(chan bool)

	ctxTO, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	
	go a(ctxTO, done)
	go b(ctxTO, done)

	<-done
	<-done
	//time.Sleep(time.Second)
	close(done)

}

func a(ctxTO context.Context, done chan bool) {
	for {
		select {
		case <-ctxTO.Done():
			log.Println("Func A timeout")
			done <- true
			return
		default:
			time.Sleep(time.Millisecond)	
		}
	}
}

func b(ctxTO context.Context, done chan bool) {
	for {
		select {
		case <-ctxTO.Done():
			log.Println("Func B timeout")
			done <- true
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
