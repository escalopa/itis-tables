package memory

import (
	"context"
	"time"

	"github.com/escalopa/itis-tables/core"
)

type TableRepository struct {
	db map[string]map[time.Weekday][]core.Subject
}

func NewTableRepository() *TableRepository {
	return &TableRepository{db: make(map[string]map[time.Weekday][]core.Subject)}
}

func (tr *TableRepository) GetSchedule(ctx context.Context, group string, day time.Weekday) ([]core.Subject, error) {
	if _, ok := tr.db[group]; !ok {
		return nil, core.ErrNotFound("could not find group: ", group)
	}
	subjects, ok := tr.db[group][day]
	if !ok || subjects == nil || len(subjects) == 0 {
		return nil, core.ErrNotFound("could not find subjects for day: ", day.String())
	}
	return subjects, nil
}

func (tr *TableRepository) SetShedule(ctx context.Context, group string, day time.Weekday, subjects []core.Subject) error {
	if _, ok := tr.db[group]; !ok {
		tr.db[group] = make(map[time.Weekday][]core.Subject)
	}
	tr.db[group][day] = subjects
	return nil
}
