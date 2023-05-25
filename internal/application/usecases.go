package application

import (
	"context"

	"github.com/escalopa/itis-tables/core"
)

type UseCase struct {
	tr  TableRepository
	cr  CourseRepository
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

func WithCourseRepository(cr CourseRepository) func(*UseCase) {
	return func(uc *UseCase) {
		uc.cr = cr
	}
}

func WithEvenOddDate(eod EvenOddDate) func(*UseCase) {
	return func(uc *UseCase) {
		uc.eod = eod
	}
}

func (uc *UseCase) GetSchedule() ([]core.Subject, error) {
	return nil, nil
}
