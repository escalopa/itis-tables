package application

import (
	"context"
	"time"

	"github.com/escalopa/itis-tables/core"
	"golang.org/x/net/html"
)

type TableParser interface {
	PraseTable(ctx context.Context, doc *html.Node) (map[time.Weekday][][]core.Subject, error)
}

type TableRepository interface {
	GetSchedule(ctx context.Context, group string, day time.Weekday) ([][]core.Subject, error)
	SetShedule(ctx context.Context, group string, day time.Weekday, subjects [][]core.Subject) error
}

type EvenOddDate interface {
	GetWeek(ctx context.Context, now time.Time) core.WeekType
}

type PageFetcher interface {
	FetchTablePage(ctx context.Context, group string, date string) (*html.Node, error)
}
