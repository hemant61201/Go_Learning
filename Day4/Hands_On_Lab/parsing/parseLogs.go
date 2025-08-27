package parsing

import (
	"Simple_Programs/Day4/Hands_On_Lab/loading"
	"fmt"
	"regexp"
	"sync"
	"time"
)

var waitingGroup sync.WaitGroup

var waitingLock sync.Mutex

const numberOfWorkers = 20

var logChannel = loading.LogChannel

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}
type ParsedLogs struct {
	logEntries []LogEntry
}

func Parse() {

	fmt.Println("Parsing logs...")

	fmt.Println(time.Now())

	parsedLogs := &ParsedLogs{}

	for worker := 0; worker < numberOfWorkers; worker++ {
		waitingGroup.Add(1)
		go parsedLogs.writeLogs()
	}

	waitingGroup.Wait()

	fmt.Println("Parsing logs...done")

	fmt.Println(time.Now())
}

func (parsedLogs *ParsedLogs) writeLogs() {

	defer waitingGroup.Done()

	for item := range logChannel {

		regex := regexp.MustCompile("(\\d+-\\d+-\\d+\\s+\\d+:\\d+:\\d+)\\s+(\\S+)\\s+(.*)")

		matches := regex.FindStringSubmatch(item)

		if len(matches) == 4 {

			logEntry := LogEntry{Timestamp: matches[1], Level: matches[2], Message: matches[3]}

			waitingLock.Lock()

			parsedLogs.logEntries = append(parsedLogs.logEntries, logEntry)

			waitingLock.Unlock()
		}
	}
}
