package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/escalopa/itis-tables/internal/adapters/db/memory"
	"github.com/escalopa/itis-tables/internal/adapters/evenodd"
	"github.com/escalopa/itis-tables/internal/adapters/parser"
	"github.com/escalopa/itis-tables/internal/application"
	"github.com/escalopa/itis-tables/internal/server"
)

func main() {
	appCtx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer func() {
		log.Println("app is shutting down")
	}()

	cr := memory.NewCourseRepository()
	tr := memory.NewTableRepository()

	tp := parser.NewTableParser(tr)
	cp := parser.NewCourseParser(cr)

	var err error

	// Parse tables
	err = tp.PraseTable(appCtx, os.Getenv("TABLE_PATH"))
	if err != nil {
		log.Println("failed to parse tables", err)
		return
	}
	log.Println("parsing table is done")

	// Parse courses
	err = cp.ParseCourses(appCtx, os.Getenv("COURSES_PATH"))
	if err != nil {
		log.Println("failed to parse tables", err)
		return
	}
	log.Println("parsing courses is done")

	// Create even odd date time
	startTime, err := time.Parse("2006-01-02", os.Getenv("START_TIME"))
	if err != nil {
		log.Println("failed to parse start time", err)
		return
	}
	log.Println("evenodd start time is", startTime)

	eod := evenodd.NewEvenOddDateTime(startTime)
	log.Println("even odd date time is created")

	// Create Usecases with repositories
	uc := application.New(appCtx,
		application.WithCourseRepository(cr),
		application.WithTableRepository(tr),
		application.WithEvenOddDate(eod),
	)

	// Create & run server
	s := server.New(appCtx, uc)

	log.Println("server is running")
	err = s.Run(os.Getenv("PORT"))
	if err != nil && err != http.ErrServerClosed {
		log.Println("failed to run server", err)
		return
	}
}
