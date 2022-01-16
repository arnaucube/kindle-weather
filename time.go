package main

import (
	"fmt"
	"strconv"
	"time"
)

func unmarshalTime(s []byte) (time.Time, error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return time.Now(), err
	}
	return time.Unix(q, 0), nil
}

type CurrentDate time.Time

func (t *CurrentDate) UnmarshalJSON(s []byte) (err error) {
	*(*time.Time)(t), err = unmarshalTime(s)
	return err
}
func (t CurrentDate) String() string {
	return time.Time(t).Format("2006-01-02, 15h")
}

type HourOnly time.Time

func (t *HourOnly) UnmarshalJSON(s []byte) (err error) {
	*(*time.Time)(t), err = unmarshalTime(s)
	return err
}
func (t HourOnly) String() string {
	return time.Time(t).Format("15h")
}

type Day time.Time

func (t *Day) UnmarshalJSON(s []byte) (err error) {
	*(*time.Time)(t), err = unmarshalTime(s)
	return err
}
func (t Day) String() string {
	return time.Time(t).Format("Mon")
}

type Hour time.Time

func (t *Hour) UnmarshalJSON(s []byte) (err error) {
	*(*time.Time)(t), err = unmarshalTime(s)
	return err
}
func (t Hour) String() string {
	return time.Time(t).Format("15:04:05")
}

type Temperature float64

func (t Temperature) String() string {
	return fmt.Sprintf("%.0f", t)
}
