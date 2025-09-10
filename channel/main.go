package channel

import (
	"fmt"
	"time"
)

func Main() {
	ch := make(chan int)

	go func() {
		// defer close(ch)
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(value)
		}

	}()
	time.Sleep(1 * time.Second)
}
