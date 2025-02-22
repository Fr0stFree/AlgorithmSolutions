package main

import (
	"errors"
	"fmt"
	"time"
)

type TimeOfDay struct {
	date time.Time
}

func (t TimeOfDay) Hour() int {
	return t.date.Hour()
}

func (t TimeOfDay) Minute() int {
	return t.date.Minute()
}

func (t TimeOfDay) Second() int {
	return t.date.Second()
}

func (t TimeOfDay) String() string {
	return t.date.Format("15:04:05 MST")
}

func (t TimeOfDay) Equal(other TimeOfDay) bool {
	if t.date.Location().String() != other.date.Location().String() {
		return false
	}
	return t.date.Equal(other.date)
}

func (t TimeOfDay) Before(other TimeOfDay) (bool, error) {
	if t.date.Location().String() != other.date.Location().String() {
		return false, errors.New("location are not equal")
	}
	return t.date.Before(other.date), nil
}

func (t TimeOfDay) After(other TimeOfDay) (bool, error) {
	if t.date.Location().String() != other.date.Location().String() {
		return false, errors.New("location are not equal")
	}
	return t.date.After(other.date), nil
}

func MakeTimeOfDay(hour, min, sec int, loc *time.Location) TimeOfDay {
	date := time.Date(1970, 1, 1, hour, min, sec, 0, loc)
	return TimeOfDay{date}
}

func main() {
	t1 := MakeTimeOfDay(17, 45, 22, time.UTC)
	t2 := MakeTimeOfDay(20, 3, 4, time.UTC)

	if t1.Equal(t2) {
		fmt.Printf("%v should not be equal to %v", t1, t2)
	}

	before, _ := t1.Before(t2)
	if !before {
		fmt.Printf("%v should be before %v", t1, t2)
	}

	after, _ := t1.After(t2)
	if after {
		fmt.Printf("%v should NOT be after %v", t1, t2)
	}
	t1 = MakeTimeOfDay(12, 34, 56, time.UTC)
	t2 = MakeTimeOfDay(12, 34, 56, time.FixedZone("UTC", 0))
	fmt.Println(t1.Equal(t2))
}
