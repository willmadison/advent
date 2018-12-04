package advent2018

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Timestamp time.Time
	Content   string
}

func ParseEventLog(r io.Reader) []Entry {
	var entries []Entry

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		entries = append(entries, parseEntry(scanner.Text()))
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Timestamp.Before(entries[j].Timestamp)
	})

	return entries
}

var logEntryRegEx = regexp.MustCompile(`^\[(.+)\] (.*)$`)

func parseEntry(rawEntry string) Entry {
	matches := logEntryRegEx.FindStringSubmatch(rawEntry)

	if len(matches) < 3 {
		fmt.Println("Hmmmm something's weird.... Here's the rawEntry:", rawEntry, "Here's what matched:", matches)
	}

	rawTimestamp := matches[1]
	t, _ := time.Parse("2006-01-02 15:04", rawTimestamp)

	content := matches[2]

	return Entry{Timestamp: t, Content: content}
}

type Nap struct {
	Start, End time.Time
	Duration   time.Duration
}

var nilTime time.Time

func FindSleepiestGuard(entries []Entry) (int, []Nap) {
	napsByGuard := getNapsByGuard(entries)

	var maxNapDuation time.Duration

	var sleepiestGuard int

	for guard, naps := range napsByGuard {
		var totalNapTime time.Duration

		for _, nap := range naps {
			totalNapTime += nap.Duration
		}

		if totalNapTime > maxNapDuation {
			maxNapDuation = totalNapTime
			sleepiestGuard = guard
		}
	}

	return sleepiestGuard, napsByGuard[sleepiestGuard]
}

func getNapsByGuard(entries []Entry) map[int][]Nap {
	napsByGuard := map[int][]Nap{}

	var guardOnDuty int

	var napStart time.Time

	for _, entry := range entries {
		switch {
		case strings.HasSuffix(entry.Content, "begins shift"):
			rawGuardNumber := strings.TrimSuffix(strings.TrimPrefix(entry.Content, "Guard #"), " begins shift")
			guardOnDuty, _ = strconv.Atoi(rawGuardNumber)
		case strings.HasSuffix(entry.Content, "falls asleep"):
			napStart = entry.Timestamp
		case strings.HasSuffix(entry.Content, "wakes up"):
			nap := Nap{napStart, entry.Timestamp, entry.Timestamp.Sub(napStart)}
			napsByGuard[guardOnDuty] = append(napsByGuard[guardOnDuty], nap)
		}
	}

	return napsByGuard
}

func FindSleepiestMinute(naps []Nap) (int, int) {
	minutes := make([]int, 60)

	for _, nap := range naps {
		for i := nap.Start.Minute(); i < nap.End.Minute(); i++ {
			minutes[i]++
		}
	}

	minute := indexOfMax(minutes)

	return minute, minutes[minute]
}

func indexOfMax(numbers []int) int {
	var maxAt int

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[maxAt] {
			maxAt = i
		}
	}

	return maxAt
}

func FindSleepiestGuardMinute(entries []Entry) (int, int) {
	napsByGuard := getNapsByGuard(entries)

	var sleepiestGuard int
	var sleepiestMinute int
	maxTimeInGivenMinute := -1

	for guard, naps := range napsByGuard {
		minute, timeSpentInMinute := FindSleepiestMinute(naps)

		if timeSpentInMinute > maxTimeInGivenMinute {
			maxTimeInGivenMinute = timeSpentInMinute
			sleepiestGuard = guard
			sleepiestMinute = minute
		}
	}

	return sleepiestGuard, sleepiestMinute
}
