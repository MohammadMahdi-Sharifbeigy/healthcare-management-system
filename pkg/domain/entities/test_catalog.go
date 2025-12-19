package entities

import (
	"time"
)

type TestCatalog struct {
	TestCatalogID int
	TestName      string
	TestCategory  string // imaging, laboratory, cognitive
	Description   string
	NormalRange   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (tc *TestCatalog) Validate() error {
	if tc.TestName == "" {
		return ErrInvalidInput("test name required")
	}
	if tc.TestCategory == "" {
		return ErrInvalidInput("test category required")
	}
	return nil
}
