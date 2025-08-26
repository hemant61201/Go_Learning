package main

func countReports(numSentCh chan int) int {
	count := 0
	for {
		ch, ok := <-numSentCh
		if !ok {
			break
		}
		count += ch
	}

	return count
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
