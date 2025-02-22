package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func asLegacyDate(t time.Time) string {
	seconds := t.Unix()
	nanos := t.Nanosecond()
	if nanos == 0 {
		return fmt.Sprintf("%d.0", seconds)
	}
	fraction := strings.TrimRight(fmt.Sprintf("%.9f", float64(nanos)/1e9)[2:], "0")
	return fmt.Sprintf("%d.%s", seconds, fraction)
}

func parseLegacyDate(d string) (time.Time, error) {
	parts := strings.Split(d, ".")
	if len(parts) != 2 {
		return time.Time{}, errors.New("invalid time format")
	}
	seconds, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return time.Time{}, errors.New("invalid seconds value")
	}
	if _, err := strconv.Atoi(parts[1]); err != nil {
		return time.Time{}, errors.New("invalid fraction value")
	}
	fraction := parts[1] + strings.Repeat("0", 9-len(parts[1]))
	nanos, err := strconv.ParseInt(fraction, 10, 64)
	if err != nil {
		return time.Time{}, errors.New("invalid nanos value")
	}
	return time.Unix(seconds, nanos).UTC(), nil
}

func main() {
	samples := map[time.Time]string{
		time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC): "3600.123456789",
		time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC):         "3600.0",
		time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC):         "0.0",
	}
	for src, want := range samples {
		got := asLegacyDate(src)
		if got != want {
			fmt.Printf("%v: got %v, want %v\n", src, got, want)
		}
	}

	samples2 := map[string]time.Time{
		"3600.123456789": time.Date(1970, 1, 1, 1, 0, 0, 123456789, time.UTC),
		"3600.0":         time.Date(1970, 1, 1, 1, 0, 0, 0, time.UTC),
		"0.0":            time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC),
		"1.123456789":    time.Date(1970, 1, 1, 0, 0, 1, 123456789, time.UTC),
	}
	for src, want := range samples2 {
		got, err := parseLegacyDate(src)
		if err != nil {
			fmt.Printf("%v: unexpected error", src)
		}
		if !got.Equal(want) {
			fmt.Printf("%v: got %v, want %v", src, got, want)
		}
	}
}
