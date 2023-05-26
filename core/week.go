package core

import "time"

type WeekType int

const (
	UNKOWN WeekType = iota
	WeekEven
	WeekOdd
	WeekAll
)

var (
	DaysInOrder = []time.Weekday{
		time.Monday,
		time.Tuesday,
		time.Wednesday,
		time.Thursday,
		time.Friday,
		time.Saturday,
	}
)
