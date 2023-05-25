package evenodd

import (
	"context"
	"fmt"
	"time"

	"github.com/escalopa/itis-tables/core"
)

type EvenOddDateTime struct {
	startDate time.Time
}

func NewEvenOddDateTime(startDate time.Time) *EvenOddDateTime {
	return &EvenOddDateTime{
		startDate: startDate,
	}
}

func (eod *EvenOddDateTime) GetWeek(ctx context.Context, now time.Time) core.WeekType {
	_, weekOdd := eod.startDate.ISOWeek()
	_, weekCurr := now.ISOWeek()

	fmt.Println(weekOdd, weekCurr, (weekCurr-weekOdd)%2)
	if (weekCurr-weekOdd)%2 == 0 {
		return core.WeekOdd
	} else {
		return core.WeekEven
	}
}
