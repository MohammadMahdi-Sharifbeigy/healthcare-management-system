package entities

import (
	"time"
)

type MedicalTest struct {
	TestID        int
	PatientID     int
	TestCatalogID int
	TestDate      time.Time
	ResultValue   string // For quantitative results
	NormalRange   string // Reference range
	Findings      string // Description of findings
	ResultStatus  string // normal, abnormal, critical
	ImageURL      string // URL to imaging files if any
	InterpretedBy *int   // DoctorID (nullable - may not be interpreted yet)
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (mt *MedicalTest) Validate() error {
	if mt.PatientID == 0 {
		return ErrInvalidInput("patient ID required")
	}
	if mt.TestCatalogID == 0 {
		return ErrInvalidInput("test catalog ID required")
	}
	if mt.TestDate.IsZero() {
		return ErrInvalidInput("test date required")
	}
	if mt.ResultStatus == "" {
		return ErrInvalidInput("result status required")
	}
	if mt.ResultStatus != "normal" && mt.ResultStatus != "abnormal" && mt.ResultStatus != "critical" {
		return ErrInvalidInput("invalid result status: must be normal, abnormal, or critical")
	}
	return nil
}

func (mt *MedicalTest) IsAbnormal() bool {
	return mt.ResultStatus == "abnormal" || mt.ResultStatus == "critical"
}

func (mt *MedicalTest) IsCritical() bool {
	return mt.ResultStatus == "critical"
}

func (mt *MedicalTest) IsNormal() bool {
	return mt.ResultStatus == "normal"
}

func (mt *MedicalTest) NeedsInterpretation() bool {
	return mt.InterpretedBy == nil
}
