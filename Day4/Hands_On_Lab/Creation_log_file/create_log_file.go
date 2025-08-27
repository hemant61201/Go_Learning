package Hands_On_Lab

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Create("Day4/Hands_On_Lab/large_log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	levels := []string{"INFO", "DEBUG", "WARNING", "ERROR", "CRITICAL", "EMERGENCY"}
	words := []string{"user", "server", "connection", "failed", "success", "timeout", "database", "request", "response", "cache"}

	targetSize := int64(1 << 30) // 1GB
	var size int64

	rand.Seed(time.Now().UnixNano())

	for size < targetSize {
		msg := make([]string, rand.Intn(8)+5)
		for i := range msg {
			msg[i] = words[rand.Intn(len(words))]
		}
		line := fmt.Sprintf("2025-08-01 00:00:00 %s %s\n",
			levels[rand.Intn(len(levels))], strings.Join(msg, " "))
		n, _ := file.WriteString(line)
		size += int64(n)
	}
}
