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

	log.Println("app is starting")

	var err error

	// Create tables repository
	tr := memory.NewTableRepository()
	log.Println("repositories are created")

	// Create tables parsers
	tp := parser.NewTableParser(tr)
	log.Println("parsers are created")

	// Creaet page fetcher
	pf := parser.NewPageFetcher(os.Getenv("TABLES_URL"), os.Getenv("TABLES_PSUB"))
	log.Println("page fetcher is created")

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
		application.WithTableRepository(tr),
		application.WithTableParser(tp),
		application.WithPageFetcher(pf),
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
