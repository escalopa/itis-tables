package parser

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/escalopa/itis-tables/internal/application"
)

type CourseParser struct {
	tr application.CourseRepository
}

func NewCourseParser(tr application.CourseRepository) *CourseParser {
	return &CourseParser{tr: tr}
}

func (*CourseParser) ParseCourses(ctx context.Context, path string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Printf("failed to close csv file, %s", err)
		}
	}()

	// Create csv reader
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}

		log.Println(record)
	}

	return nil
}
