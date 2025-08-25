package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type LogEntry struct {
	Timestamp string
	Level     string
	Message   string
}

type Repository interface {
	List() ([]LogEntry, error)
	Query(level string) ([]LogEntry, error)
	Add(entry LogEntry) error
	Delete(index int) error
	Sort(by string) ([]LogEntry, error)
}

type CSVRepo struct {
	file string
	data []LogEntry
}

func (r *CSVRepo) load() error {
	f, err := os.Open(r.file)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var entries []LogEntry

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, " ", 4) // split into max 4 parts
		if len(parts) < 4 {
			continue // skip malformed line
		}
		entries = append(entries, LogEntry{
			Timestamp: parts[0] + " " + parts[1],
			Level:     parts[2],
			Message:   parts[3],
		})
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner error: %w", err)
	}

	r.data = entries
	return nil
}

func (r *CSVRepo) save() error {
	f, err := os.Create(r.file)
	if err != nil {
		return fmt.Errorf("failed to save file: %w", err)
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	writer.Comma = ' '

	for _, e := range r.data {
		record := []string{e.Timestamp, e.Level, e.Message}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	writer.Flush()
	return writer.Error()
}

func (r *CSVRepo) List() ([]LogEntry, error) {
	if err := r.load(); err != nil {
		return nil, err
	}
	return r.data, nil
}

func (r *CSVRepo) Query(level string) ([]LogEntry, error) {
	if err := r.load(); err != nil {
		return nil, err
	}
	var results []LogEntry
	for _, e := range r.data {
		if strings.EqualFold(e.Level, level) {
			results = append(results, e)
		}
	}
	if len(results) == 0 {
		return nil, errors.New("no entries found for given level")
	}
	return results, nil
}

func (r *CSVRepo) Add(entry LogEntry) error {
	if err := r.load(); err != nil {
		return err
	}
	r.data = append(r.data, entry)
	return r.save()
}

func (r *CSVRepo) Delete(index int) error {
	if err := r.load(); err != nil {
		return err
	}
	if index < 0 || index >= len(r.data) {
		return errors.New("index out of range")
	}
	r.data = append(r.data[:index], r.data[index+1:]...)
	return r.save()
}

func (r *CSVRepo) Sort(by string) ([]LogEntry, error) {
	if err := r.load(); err != nil {
		return nil, err
	}
	switch by {
	case "level":
		sort.Slice(r.data, func(i, j int) bool { return r.data[i].Level < r.data[j].Level })
	case "time":
		sort.Slice(r.data, func(i, j int) bool { return r.data[i].Timestamp < r.data[j].Timestamp })
	default:
		return nil, errors.New("unknown sort field")
	}
	return r.data, nil
}

func main() {
	action := flag.String("action", "list", "action: list|query|add|delete|sort")

	level := flag.String("level", "", "log level for query")

	msg := flag.String("msg", "", "message for add")

	idx := flag.Int("index", -1, "index for delete")

	sortBy := flag.String("by", "level", "sort by: level|time")

	flag.Parse()

	repo := &CSVRepo{file: "Day3/command_line_example/log.txt"}

	switch *action {
	case "list":
		entries, err := repo.List()
		check(err)
		printEntries(entries)

	case "query":
		if *level == "" {
			check(errors.New("query requires -level"))
		}
		entries, err := repo.Query(*level)
		check(err)
		printEntries(entries)

	case "add":
		if *level == "" || *msg == "" {
			check(errors.New("add requires -level and -msg"))
		}
		entry := LogEntry{Timestamp: "1970-01-01 00:00:00", Level: *level, Message: *msg}
		check(repo.Add(entry))
		fmt.Println("entry added")

	case "delete":
		if *idx < 0 {
			check(errors.New("delete requires -index"))
		}
		check(repo.Delete(*idx))
		fmt.Println("entry deleted")

	case "sort":
		entries, err := repo.Sort(*sortBy)
		check(err)
		printEntries(entries)

	default:
		check(errors.New("unknown action"))
	}
}

func printEntries(entries []LogEntry) {
	for i, e := range entries {
		fmt.Printf("%d: %s [%s] %s\n", i, e.Timestamp, e.Level, e.Message)
	}
}

func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
