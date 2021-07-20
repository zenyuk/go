package main

import (
	"fmt"
	"time"
)

/* Both selects will be working concurrently, using the same channels. 
   Each message will be processed only once by one of the two selects. 
   Which message will be processed by which select is "random".
   
   Example output:
    
    another skipping
    skipping
    another skipping
    skipping
    another: 1
    skipping
    another skipping
    skipping
    another skipping
    another: 2
    skipping
    another skipping
    skipping
    3

   */



func main() {
	ch := make(chan int)
	go func(){
		for i := 1; i <= 10; i++ {
			time.Sleep(2 * time.Second)
			ch <- i
		}
	}()
	go anotherSelect(ch)
	for {
		select {
		case i := <-ch:
			println(i)
		default:
			time.Sleep(time.Second)
			println("skipping")
		}
	}
}

func anotherSelect(ch <-chan int) {
	for {
		select {
		case i := <-ch:
			fmt.Printf("another: %d\n", i)
		default:
			time.Sleep(time.Second)
			println("another skipping")
		}
	}
}
