package application

import (
	"context"
	"time"

	"github.com/escalopa/itis-tables/core"
)

type UseCase struct {
	tr  TableRepository
	tp  TableParser
	pf  PageFetcher
	eod EvenOddDate
}

func New(ctx context.Context, opts ...func(*UseCase)) *UseCase {
	uc := &UseCase{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

func WithTableRepository(tr TableRepository) func(*UseCase) {
	return func(uc *UseCase) {
		uc.tr = tr
	}
}

func WithTableParser(tp TableParser) func(*UseCase) {
	return func(uc *UseCase) {
		uc.tp = tp
	}
}

func WithPageFetcher(pf PageFetcher) func(*UseCase) {
	return func(uc *UseCase) {
		uc.pf = pf
	}
}

func WithEvenOddDate(eod EvenOddDate) func(*UseCase) {
	return func(uc *UseCase) {
		uc.eod = eod
	}
}

func (uc *UseCase) GetSchedule(ctx context.Context, group string, date time.Time) ([][]core.Subject, error) {
	// Check if data already stored in cache
	if subjects, err := uc.tr.GetSchedule(ctx, group, date.Weekday()); err == nil {
		return subjects, nil
	}

	// Fetch the page and parse it
	d := date.Format("02.01.2006")
	body, err := uc.pf.FetchTablePage(ctx, group, d)
	if err != nil {
		return nil, err
	}

	// Parse the table
	tables, err := uc.tp.PraseTable(ctx, body)
	if err != nil {
		return nil, err
	}

	// Store in cache
	go func(tables map[time.Weekday][][]core.Subject) {
		for i := range core.DaysInOrder {
			uc.tr.SetShedule(ctx, group, core.DaysInOrder[i], tables[core.DaysInOrder[i]])
		}
	}(tables)

	return tables[date.Weekday()], nil
}
