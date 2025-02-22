package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"time"
)

type Task struct {
	Date  time.Time
	Dur   time.Duration
	Title string
}

func ParsePage(src string) ([]Task, error) {
	lines := strings.Split(strings.TrimSpace(src), "\n")
	if len(lines) < 2 {
		return nil, errors.New("invalid journal format")
	}

	date, err := parseDate(lines[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %w", err)
	}

	tasks, err := parseTasks(date, lines[1:])
	if err != nil {
		return nil, fmt.Errorf("failed to parse tasks: %w", err)
	}

	sortTasks(tasks)
	return tasks, nil
}

func parseDate(src string) (time.Time, error) {
	layout := "02.01.2006"
	parsed, err := time.ParseInLocation(layout, src, time.UTC)
	if err != nil {
		return time.Time{}, err
	}
	return parsed, nil
}

func parseTasks(date time.Time, lines []string) ([]Task, error) {
	var tasksMap = make(map[string]time.Duration)
	re := regexp.MustCompile(`(\d{1,2}:\d{2}) - (\d{1,2}:\d{2}) (.+)`) 
	
	for _, line := range lines {
		matches := re.FindStringSubmatch(line)
		if len(matches) != 4 {
			return nil, fmt.Errorf("invalid task format: %s", line)
		}

		start, err := time.Parse("15:04", matches[1])
		if err != nil {
			return nil, fmt.Errorf("invalid start time in: %s", line)
		}
		end, err := time.Parse("15:04", matches[2])
		if err != nil {
			return nil, fmt.Errorf("invalid end time in: %s", line)
		}
		if end.Before(start) || end == start {
			return nil, fmt.Errorf("end time before start time in: %s", line)
		}

		dur := end.Sub(start)
		title := matches[3]

		tasksMap[title] += dur
	}

	var tasks []Task
	for title, dur := range tasksMap {
		tasks = append(tasks, Task{Date: date, Dur: dur, Title: title})
	}

	return tasks, nil
}

func sortTasks(tasks []Task) {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Dur > tasks[j].Dur
	})
}

func main() {
	page := `15.04.2022
8:00 - 8:30 Завтрак
8:30 - 9:30 Оглаживание кота
9:30 - 10:00 Интернеты
10:00 - 14:00 Напряженная работа
14:00 - 14:45 Обед
14:45 - 15:00 Оглаживание кота
15:00 - 19:00 Напряженная работа
19:00 - 19:30 Интернеты
19:30 - 22:30 Безудержное веселье
22:30 - 23:00 Оглаживание кота`

	entries, err := ParsePage(page)
	if err != nil {
		panic(err)
	}

	fmt.Println("Мои достижения за", entries[0].Date.Format("2006-01-02"))
	for _, entry := range entries {
		fmt.Printf("- %v: %v\n", entry.Title, entry.Dur)
	}
}
