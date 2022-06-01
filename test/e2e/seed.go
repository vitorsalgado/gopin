package e2e

import (
	"database/sql"
	"fmt"
	"github.com/vitorsalgado/gopin/internal/domain"
	"time"
)

const (
	u1       = "79561481-fc11-419c-a9e8-e5a079b853c1"
	u2       = "79561481-fc11-419c-a9e8-e5a079b853c2"
	session1 = "daff4b9f-24e2-478d-8b42-6d3f59a08b31"
	session2 = "daff4b9f-24e2-478d-8b42-6d3f59a08b32"
	session3 = "daff4b9f-24e2-478d-8b42-6d3f59a08b33"
)

type Seed struct {
	db *sql.DB
}

func (s *Seed) seed() {
	fmt.Println("seeding Database ...")

	repository := domain.NewLocationRepository(s.db)

	// User 1
	// --
	_ = repository.ReportNew(domain.Location{
		UserID:     u1,
		SessionID:  session1,
		Latitude:   -33.22325847832756,
		Longitude:  -70.21369951517998,
		Precision:  1000,
		ReportedAt: time.Now(),
	})
	_ = repository.ReportNew(domain.Location{
		UserID:     u1,
		SessionID:  session2,
		Latitude:   -30.22325847832756,
		Longitude:  -70.21369951517998,
		Precision:  1250,
		ReportedAt: time.Now().Add(-5 * time.Minute),
	})
	_ = repository.ReportNew(domain.Location{
		UserID:     u1,
		SessionID:  session2,
		Latitude:   -25.22325847832756,
		Longitude:  -65.21369951517998,
		Precision:  1500,
		ReportedAt: time.Now().Add(-30 * time.Minute),
	})

	// User 2
	// ..
	_ = repository.ReportNew(domain.Location{
		UserID:     u2,
		SessionID:  session3,
		Latitude:   -10.22325847832756,
		Longitude:  -50.21369951517998,
		Precision:  50,
		ReportedAt: time.Now().Add(-15 * time.Minute),
	})
	_ = repository.ReportNew(domain.Location{
		UserID:     u2,
		SessionID:  session3,
		Latitude:   -15.22325847832756,
		Longitude:  -40.21369951517998,
		Precision:  2000,
		ReportedAt: time.Now().Add(-30 * time.Minute),
	})

	fmt.Println("seeding complete")
}

func (s *Seed) cleanDb() {
	fmt.Println("cleaning database ...")

	stmt, err := s.db.Prepare("DELETE FROM locations")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}

	fmt.Println("cleaning complete ...")
}
