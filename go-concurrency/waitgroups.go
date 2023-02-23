package main

import (
	"fmt"
	"sync"
	"time"
)

func process(processNum int) {
	fmt.Printf("Now processing process #%d\n", processNum)
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("Completed processing of process #%d\n", processNum)
}

func ProcessManager() {
	var waitGroup sync.WaitGroup

	for i := 1; i <= 10; i++ {
		waitGroup.Add(1)

		// SEE NOTES AS TO WHY DO WE REASSIGN i TO ANOTHER VALUE j
		i := i

		go func() {
			defer waitGroup.Done()
			process(i)
		}()
	}
	waitGroup.Wait()
}
