package memory

import (
	"context"
	"time"

	"github.com/escalopa/itis-tables/core"
)

type TableRepository struct {
	db map[time.Weekday][]core.Subject
}

func NewTableRepository() *TableRepository {
	return &TableRepository{db: make(map[time.Weekday][]core.Subject)}
}

func (tr *TableRepository) GetSchedule(ctx context.Context, day time.Weekday) ([]core.Subject, error) {
	subjects, ok := tr.db[day]
	if !ok {
		return nil, core.ErrNotFound("could not find subjects for day: ", day.String())
	}
	return subjects, nil
}

func (tr *TableRepository) SetShedule(ctx context.Context, day time.Weekday, subjects []core.Subject) error {
	tr.db[day] = subjects
	return nil
}
