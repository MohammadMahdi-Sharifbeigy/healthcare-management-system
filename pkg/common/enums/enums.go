package enums

// Gender enum
type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

func (g Gender) String() string {
	return string(g)
}

func (g Gender) IsValid() bool {
	return g == GenderMale || g == GenderFemale || g == GenderOther
}

// BloodType enum
type BloodType string

const (
	BloodTypeOPositive  BloodType = "O+"
	BloodTypeONegative  BloodType = "O-"
	BloodTypeAPositive  BloodType = "A+"
	BloodTypeANegative  BloodType = "A-"
	BloodTypeBPositive  BloodType = "B+"
	BloodTypeBNegative  BloodType = "B-"
	BloodTypeABPositive BloodType = "AB+"
	BloodTypeABNegative BloodType = "AB-"
)

func (bt BloodType) String() string {
	return string(bt)
}

func (bt BloodType) IsValid() bool {
	validTypes := map[BloodType]bool{
		BloodTypeOPositive:  true,
		BloodTypeONegative:  true,
		BloodTypeAPositive:  true,
		BloodTypeANegative:  true,
		BloodTypeBPositive:  true,
		BloodTypeBNegative:  true,
		BloodTypeABPositive: true,
		BloodTypeABNegative: true,
	}
	return validTypes[bt]
}

// Severity enum
type Severity string

const (
	SeverityMild     Severity = "mild"
	SeverityModerate Severity = "moderate"
	SeveritySevere   Severity = "severe"
)

func (s Severity) String() string {
	return string(s)
}

func (s Severity) IsValid() bool {
	return s == SeverityMild || s == SeverityModerate || s == SeveritySevere
}

// AppointmentStatus enum
type AppointmentStatus string

const (
	StatusConfirmed AppointmentStatus = "confirmed"
	StatusCompleted AppointmentStatus = "completed"
	StatusCancelled AppointmentStatus = "cancelled"
)

func (as AppointmentStatus) String() string {
	return string(as)
}

func (as AppointmentStatus) IsValid() bool {
	return as == StatusConfirmed || as == StatusCompleted || as == StatusCancelled
}

// TestResultStatus enum
type TestResultStatus string

const (
	ResultNormal   TestResultStatus = "normal"
	ResultAbnormal TestResultStatus = "abnormal"
	ResultCritical TestResultStatus = "critical"
)

func (trs TestResultStatus) String() string {
	return string(trs)
}

func (trs TestResultStatus) IsValid() bool {
	return trs == ResultNormal || trs == ResultAbnormal || trs == ResultCritical
}

// TreatmentPlanStatus enum
type TreatmentPlanStatus string

const (
	PlanActive    TreatmentPlanStatus = "active"
	PlanCompleted TreatmentPlanStatus = "completed"
	PlanCancelled TreatmentPlanStatus = "cancelled"
)

func (tps TreatmentPlanStatus) String() string {
	return string(tps)
}

func (tps TreatmentPlanStatus) IsValid() bool {
	return tps == PlanActive || tps == PlanCompleted || tps == PlanCancelled
}

// MedicationForm enum
type MedicationForm string

const (
	FormTablet    MedicationForm = "tablet"
	FormCapsule   MedicationForm = "capsule"
	FormInjection MedicationForm = "injection"
	FormLiquid    MedicationForm = "liquid"
)

func (mf MedicationForm) String() string {
	return string(mf)
}

func (mf MedicationForm) IsValid() bool {
	return mf == FormTablet || mf == FormCapsule || mf == FormInjection || mf == FormLiquid
}

// TreatmentType enum
type TreatmentType string

const (
	TreatmentPharmacotherapy TreatmentType = "pharmacotherapy"
	TreatmentPhysiotherapy   TreatmentType = "physiotherapy"
	TreatmentPsychotherapy   TreatmentType = "psychotherapy"
	TreatmentOccupational    TreatmentType = "occupational"
	TreatmentSpeech          TreatmentType = "speech"
	TreatmentRehabilitation  TreatmentType = "rehabilitation"
)

func (tt TreatmentType) String() string {
	return string(tt)
}

func (tt TreatmentType) IsValid() bool {
	validTypes := map[TreatmentType]bool{
		TreatmentPharmacotherapy: true,
		TreatmentPhysiotherapy:   true,
		TreatmentPsychotherapy:   true,
		TreatmentOccupational:    true,
		TreatmentSpeech:          true,
		TreatmentRehabilitation:  true,
	}
	return validTypes[tt]
}

// Specialization enum
type Specialization string

const (
	SpecNeurology    Specialization = "Neurology"
	SpecPsychiatry   Specialization = "Psychiatry"
	SpecRadiology    Specialization = "Radiology"
	SpecPhysiology   Specialization = "Physiology"
	SpecPsychology   Specialization = "Psychology"
	SpecLaboratory   Specialization = "Laboratory"
)

func (s Specialization) String() string {
	return string(s)
}

func (s Specialization) IsValid() bool {
	validSpecs := map[Specialization]bool{
		SpecNeurology:  true,
		SpecPsychiatry: true,
		SpecRadiology:  true,
		SpecPhysiology: true,
		SpecPsychology: true,
		SpecLaboratory: true,
	}
	return validSpecs[s]
}

// UserRole enum
type UserRole string

const (
	RoleAdmin  UserRole = "admin"
	RoleDoctor UserRole = "doctor"
	RoleNurse  UserRole = "nurse"
	RolePatient UserRole = "patient"
)

func (ur UserRole) String() string {
	return string(ur)
}

func (ur UserRole) IsValid() bool {
	return ur == RoleAdmin || ur == RoleDoctor || ur == RoleNurse || ur == RolePatient
}

// TestCategory enum
type TestCategory string

const (
	TestImaging   TestCategory = "imaging"
	TestLab       TestCategory = "laboratory"
	TestCognitive TestCategory = "cognitive"
)

func (tc TestCategory) String() string {
	return string(tc)
}

func (tc TestCategory) IsValid() bool {
	return tc == TestImaging || tc == TestLab || tc == TestCognitive
}
