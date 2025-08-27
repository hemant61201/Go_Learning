package loading

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var waitingGroup sync.WaitGroup

var LogChannel = make(chan string)

type LoadLogs struct {
	FileName string
	logs     []string
}

func (loadLogs *LoadLogs) Load() error {

	fmt.Println("Loading logs...")

	file, err := os.Open(loadLogs.FileName)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	waitingGroup.Add(1)

	go loadLogs.readLogs(scanner)

	waitingGroup.Wait()

	fmt.Println("Loaded logs")

	return nil
}

func (loadLogs *LoadLogs) readLogs(scanner *bufio.Scanner) {

	defer waitingGroup.Done()

	for scanner.Scan() {
		line := scanner.Text()
		loadLogs.logs = append(loadLogs.logs, line)
		LogChannel <- line
	}

	close(LogChannel)
}
