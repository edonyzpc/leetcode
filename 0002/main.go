package main

import "sync"

func main() {
	wg := sync.WaitGroup{}

	si := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	/* Data Race
	 * to find the data race, run the command line `CGO_ENABLED=1 go run -race main.go`
	 */
	for i := range si {
		wg.Add(1)
		go func() {
			println(i)
			wg.Done()
		}()
	}

	// to solve the data race, pass the real value to the closure
	for i := range si {
		wg.Add(1)
		go func(i int) {
			println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()
}
