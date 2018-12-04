package advent2018

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseEventLog(t *testing.T) {
	cases := []struct {
		given    string
		expected []Entry
	}{
		{
			`[1518-10-31 00:58] wakes up
[1518-02-27 00:57] wakes up
[1518-04-05 00:03] falls asleep`,
			[]Entry{
				{time.Date(1518, time.February, 27, 0, 57, 0, 0, time.UTC), `wakes up`},
				{time.Date(1518, time.April, 5, 0, 3, 0, 0, time.UTC), `falls asleep`},
				{time.Date(1518, time.October, 31, 0, 58, 0, 0, time.UTC), `wakes up`},
			},
		},
	}

	for _, tc := range cases {
		actual := ParseEventLog(strings.NewReader(tc.given))
		assert.Equal(t, tc.expected, actual)
	}
}

func TestFindSleepiestGuardAndMinute(t *testing.T) {
	cases := []struct {
		given    []Entry
		expected struct {
			GuardID, NumberOfNaps, SleepiestMinute int
		}
	}{
		{
			ParseEventLog(strings.NewReader(`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`)),
			struct {
				GuardID, NumberOfNaps, SleepiestMinute int
			}{
				10, 3, 24,
			},
		},
	}

	for _, tc := range cases {
		actualGuardID, naps := FindSleepiestGuard(tc.given)
		assert.Equal(t, tc.expected.GuardID, actualGuardID)
		assert.Equal(t, tc.expected.NumberOfNaps, len(naps))

		actualSleepiestMinute, _ := FindSleepiestMinute(naps)
		assert.Equal(t, tc.expected.SleepiestMinute, actualSleepiestMinute)
	}
}

func TestFindSleepiestGuardMinute(t *testing.T) {
	cases := []struct {
		given    []Entry
		expected struct {
			GuardID, SleepiestMinute int
		}
	}{
		{
			ParseEventLog(strings.NewReader(`[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
[1518-11-04 00:46] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`)),
			struct {
				GuardID, SleepiestMinute int
			}{
				99, 45,
			},
		},
	}

	for _, tc := range cases {
		actualGuardID, actualMinute := FindSleepiestGuardMinute(tc.given)
		assert.Equal(t, tc.expected.GuardID, actualGuardID)
		assert.Equal(t, tc.expected.SleepiestMinute, actualMinute)
	}
}
