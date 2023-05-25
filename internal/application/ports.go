package application

import (
	"context"
	"time"

	"github.com/escalopa/itis-tables/core"
)

type TableRepository interface {
	GetSchedule(ctx context.Context, day time.Weekday) ([]core.Subject, error)
	SetShedule(ctx context.Context, day time.Weekday, subjects []core.Subject) error
}

type CourseRepository interface {
	GetCourse(ctx context.Context, studentID string) ([]core.Course, error)
	SetCourse(ctx context.Context, studentID string, course core.Course) error
}

type EvenOddDate interface {
	GetWeek(ctx context.Context, now time.Time) core.WeekType
}
