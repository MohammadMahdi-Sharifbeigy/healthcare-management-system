package entities

import (
	"time"
)

type MedicalTest struct {
	TestID        int
	PatientID     int
	TestCatalogID int
	TestDate      time.Time
	ResultValue   string
	NormalRange   string
	Findings      string
	ResultStatus  string // normal, abnormal, critical
	ImageURL      string
	InterpretedBy int // DoctorID
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (mt *MedicalTest) Validate() error {
	if mt.PatientID == 0 || mt.TestCatalogID == 0 {
		return ErrInvalidInput("patient and test catalog required")
	}
	if mt.TestDate.IsZero() {
		return ErrInvalidInput("test date required")
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
