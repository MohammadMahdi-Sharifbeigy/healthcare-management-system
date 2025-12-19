package usecase

import (
	"context"
	"time"

	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/dto"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/application/services"
	"github.com/MohammadMahdi-Sharifbeigy/healthcare-management-system/pkg/domain/entities"
)

type MedicalTestUseCase struct {
	service *services.MedicalTestService
}

func NewMedicalTestUseCase(service *services.MedicalTestService) *MedicalTestUseCase {
	return &MedicalTestUseCase{service: service}
}

// OrderMedicalTest creates new medical test order
func (uc *MedicalTestUseCase) OrderMedicalTest(ctx context.Context, req *dto.CreateMedicalTestRequest) (*dto.MedicalTestResponse, error) {
	testDate, _ := time.Parse("2006-01-02", req.TestDate)

	var interpretedBy *int
	if req.InterpretedByID > 0 {
		interpretedBy = &req.InterpretedByID
	}

	test := &entities.MedicalTest{
		PatientID:     req.PatientID,
		TestCatalogID: req.TestCatalogID,
		TestDate:      testDate,
		ResultValue:   req.ResultValue,
		NormalRange:   req.NormalRange,
		Findings:      req.Findings,
		ResultStatus:  req.ResultStatus,
		ImageURL:      req.ImageURL,
		InterpretedBy: interpretedBy,
	}

	if err := uc.service.CreateTest(ctx, test); err != nil {
		return nil, err
	}

	return &dto.MedicalTestResponse{
		TestID:        test.TestID,
		PatientID:     test.PatientID,
		TestCatalogID: test.TestCatalogID,
		TestDate:      test.TestDate,
		ResultValue:   test.ResultValue,
		NormalRange:   test.NormalRange,
		Findings:      test.Findings,
		ResultStatus:  test.ResultStatus,
		ImageURL:      test.ImageURL,
		CreatedAt:     test.CreatedAt,
		UpdatedAt:     test.UpdatedAt,
	}, nil
}

// GetTestResults retrieves test details
func (uc *MedicalTestUseCase) GetTestResults(ctx context.Context, testID int) (*dto.MedicalTestResponse, error) {
	test, err := uc.service.GetTest(ctx, testID)
	if err != nil {
		return nil, err
	}

	return &dto.MedicalTestResponse{
		TestID:        test.TestID,
		PatientID:     test.PatientID,
		TestCatalogID: test.TestCatalogID,
		TestDate:      test.TestDate,
		ResultValue:   test.ResultValue,
		NormalRange:   test.NormalRange,
		Findings:      test.Findings,
		ResultStatus:  test.ResultStatus,
		ImageURL:      test.ImageURL,
		CreatedAt:     test.CreatedAt,
		UpdatedAt:     test.UpdatedAt,
	}, nil
}

// UpdateTestResults updates test results and findings
func (uc *MedicalTestUseCase) UpdateTestResults(ctx context.Context, testID int, req *dto.UpdateMedicalTestRequest) (*dto.MedicalTestResponse, error) {
	test, err := uc.service.GetTest(ctx, testID)
	if err != nil {
		return nil, err
	}

	if req.ResultValue != "" {
		test.ResultValue = req.ResultValue
	}
	if req.NormalRange != "" {
		test.NormalRange = req.NormalRange
	}
	if req.Findings != "" {
		test.Findings = req.Findings
	}
	if req.ResultStatus != "" {
		test.ResultStatus = req.ResultStatus
	}

	if err := uc.service.UpdateTestResults(ctx, test); err != nil {
		return nil, err
	}

	return &dto.MedicalTestResponse{
		TestID:        test.TestID,
		PatientID:     test.PatientID,
		TestCatalogID: test.TestCatalogID,
		TestDate:      test.TestDate,
		ResultValue:   test.ResultValue,
		NormalRange:   test.NormalRange,
		Findings:      test.Findings,
		ResultStatus:  test.ResultStatus,
		ImageURL:      test.ImageURL,
		CreatedAt:     test.CreatedAt,
		UpdatedAt:     test.UpdatedAt,
	}, nil
}

// AssignInterpreter assigns doctor to interpret test
func (uc *MedicalTestUseCase) AssignInterpreter(ctx context.Context, testID, doctorID int) error {
	return uc.service.InterpretTest(ctx, testID, doctorID)
}

// GetPatientTests retrieves all tests for patient
func (uc *MedicalTestUseCase) GetPatientTests(ctx context.Context, patientID int) ([]*dto.MedicalTestListResponse, error) {
	tests, err := uc.service.GetPatientTests(ctx, patientID)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicalTestListResponse
	for _, t := range tests {
		responses = append(responses, &dto.MedicalTestListResponse{
			TestID:       t.TestID,
			TestDate:     t.TestDate,
			ResultStatus: t.ResultStatus,
		})
	}

	return responses, nil
}

// GetAbnormalTestsAlert retrieves abnormal results alert
func (uc *MedicalTestUseCase) GetAbnormalTestsAlert(ctx context.Context) ([]*dto.MedicalTestListResponse, error) {
	tests, err := uc.service.GetAbnormalResults(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicalTestListResponse
	for _, t := range tests {
		responses = append(responses, &dto.MedicalTestListResponse{
			TestID:       t.TestID,
			TestDate:     t.TestDate,
			ResultStatus: t.ResultStatus,
		})
	}

	return responses, nil
}

// GetCriticalTestsAlert retrieves critical results alert
func (uc *MedicalTestUseCase) GetCriticalTestsAlert(ctx context.Context) ([]*dto.MedicalTestListResponse, error) {
	tests, err := uc.service.GetCriticalResults(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*dto.MedicalTestListResponse
	for _, t := range tests {
		responses = append(responses, &dto.MedicalTestListResponse{
			TestID:       t.TestID,
			TestDate:     t.TestDate,
			ResultStatus: t.ResultStatus,
		})
	}

	return responses, nil
}
