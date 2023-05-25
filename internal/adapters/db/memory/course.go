package memory

import (
	"context"

	"github.com/escalopa/itis-tables/core"
)

type CourseRepository struct {
	db map[string][]core.Course
}

func NewCourseRepository() *CourseRepository {
	return &CourseRepository{db: make(map[string][]core.Course)}
}

func (cr *CourseRepository) GetCourse(ctx context.Context, studentID string) ([]core.Course, error) {
	courses, ok := cr.db[studentID]
	if !ok {
		return nil, core.ErrNotFound("could not find courses for student: ", studentID)
	}
	return courses, nil
}

func (cr *CourseRepository) SetCourse(ctx context.Context, studentID string, course []core.Course) error {
	cr.db[studentID] = course
	return nil
}
