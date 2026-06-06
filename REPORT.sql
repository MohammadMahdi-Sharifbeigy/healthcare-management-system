-- ============================================
-- HEALTHCARE DATABASE - UPDATED REPORT PROCEDURES
-- Optimized for SQL Server (T-SQL)
-- ============================================

-- ============================================
-- REPORT 1: Complete Medical History
-- Updated joins for Diagnosis and Test_Catalog
-- ============================================
CREATE OR ALTER PROCEDURE sp_PatientMedicalHistory
    @patient_id INT
AS
BEGIN
    SELECT 
        p.patient_id,
        CONCAT(p.first_name, ' ', p.last_name) AS full_name,
        p.date_of_birth,
        p.blood_type,
        p.medical_history,
        
        -- Appointments
        a.appointment_date,
        a.reason AS appointment_reason,
        CONCAT(d1.first_name, ' ', d1.last_name) AS doctor_name,
        
        -- Diagnosis (Linked correctly via ID)
        dis.disease_name,
        diag.diagnosis_date,
        diag.severity,
        
        -- Medications
        med.medication_name,
        pm.dosage,
        pm.frequency,
        pm.end_date AS med_end_date,
        
        -- Medical Tests (Linked via Test_Catalog)
        tc.test_name,
        mt.test_date,
        mt.result_value,
        mt.result_status,
        
        -- Treatment Plans
        tp.treatment_type,
        tp.status AS plan_status,
        tp.goals
        
    FROM Patient p
    LEFT JOIN Appointment a ON p.patient_id = a.patient_id
    LEFT JOIN Doctor d1 ON a.doctor_id = d1.doctor_id
    LEFT JOIN Diagnosis diag ON p.patient_id = diag.patient_id
    LEFT JOIN Disease dis ON diag.disease_id = dis.disease_id
    LEFT JOIN Prescription pr ON p.patient_id = pr.patient_id
    LEFT JOIN Prescription_Medication pm ON pr.prescription_id = pm.prescription_id
    LEFT JOIN Medication med ON pm.medication_id = med.medication_id
    LEFT JOIN Medical_Test mt ON p.patient_id = mt.patient_id
    LEFT JOIN Test_Catalog tc ON mt.test_catalog_id = tc.test_catalog_id
    LEFT JOIN Treatment_Plan tp ON p.patient_id = tp.patient_id
    
    WHERE p.patient_id = @patient_id
    ORDER BY a.appointment_date DESC;
END;
GO

-- ============================================
-- REPORT 2: Patients Under Doctor's Care
-- ============================================
CREATE OR ALTER PROCEDURE sp_DoctorPatientsList
    @doctor_id INT
AS
BEGIN
    SELECT 
        d.doctor_id,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name,
        p.patient_id,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        p.phone,
        
        MAX(a.appointment_date) AS last_appointment_date,
        
        -- Subquery to get latest disease name correctly linked
        (SELECT TOP 1 dis.disease_name 
         FROM Diagnosis diag 
         JOIN Disease dis ON diag.disease_id = dis.disease_id
         WHERE diag.patient_id = p.patient_id AND diag.doctor_id = d.doctor_id
         ORDER BY diag.diagnosis_date DESC) AS last_diagnosis,
         
        COUNT(DISTINCT a.appointment_id) AS total_appointments
        
    FROM Doctor d
    INNER JOIN Appointment a ON d.doctor_id = a.doctor_id
    INNER JOIN Patient p ON a.patient_id = p.patient_id
    
    WHERE d.doctor_id = @doctor_id
    GROUP BY d.doctor_id, d.first_name, d.last_name, p.patient_id, p.first_name, p.last_name, p.phone
    ORDER BY last_appointment_date DESC;
END;
GO

-- ============================================
-- REPORT 3: Abnormal Test Results
-- Updated for Test_Catalog join
-- ============================================
CREATE OR ALTER PROCEDURE sp_AbnormalTestResults
AS
BEGIN
    SELECT 
        p.patient_id,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        tc.test_name,
        tc.test_category,
        mt.test_date,
        mt.result_value,
        mt.normal_range,
        mt.result_status,
        mt.findings,
        CONCAT(d.first_name, ' ', d.last_name) AS interpreted_by
        
    FROM Medical_Test mt
    JOIN Test_Catalog tc ON mt.test_catalog_id = tc.test_catalog_id
    JOIN Patient p ON mt.patient_id = p.patient_id
    LEFT JOIN Doctor d ON mt.interpreted_by = d.doctor_id
    
    WHERE mt.result_status IN ('Abnormal', 'Critical')
    ORDER BY mt.test_date DESC;
END;
GO

-- ============================================
-- REPORT 4: Active Medications for Patient
-- ============================================
CREATE OR ALTER PROCEDURE sp_ActiveMedicationsForPatient
    @patient_id INT
AS
BEGIN
    SELECT 
        p.patient_id,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        m.medication_name,
        m.strength,
        pm.dosage,
        pm.frequency,
        pm.end_date,
        CONCAT(d.first_name, ' ', d.last_name) AS prescriber,
        DATEDIFF(DAY, GETDATE(), pm.end_date) AS days_remaining
        
    FROM Patient p
    JOIN Prescription pr ON p.patient_id = pr.patient_id
    JOIN Prescription_Medication pm ON pr.prescription_id = pm.prescription_id
    JOIN Medication m ON pm.medication_id = m.medication_id
    JOIN Doctor d ON pr.doctor_id = d.doctor_id
    
    WHERE p.patient_id = @patient_id 
    AND pm.end_date >= CAST(GETDATE() AS DATE)
    ORDER BY pm.end_date ASC;
END;
GO

-- ============================================
-- REPORT 5: Treatment Plan Progress
-- Fixed: Links Treatment_Plan -> Diagnosis -> Disease
-- ============================================
CREATE OR ALTER PROCEDURE sp_TreatmentPlanProgress
    @plan_id INT
AS
BEGIN
    SELECT 
        tp.plan_id,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        dis.disease_name, -- Fetched via Diagnosis
        tp.treatment_type,
        tp.start_date,
        tp.end_date,
        tp.session_duration AS planned_sessions,
        COUNT(ts.session_id) AS completed_sessions,
        
        CASE 
            WHEN tp.session_duration > 0 
            THEN CAST(COUNT(ts.session_id) * 100.0 / tp.session_duration AS DECIMAL(5,2))
            ELSE 0 
        END AS progress_percent,
        
        tp.status
        
    FROM Treatment_Plan tp
    JOIN Patient p ON tp.patient_id = p.patient_id
    -- JOIN CHAIN FIX:
    JOIN Diagnosis diag ON tp.diagnosis_id = diag.diagnosis_id
    JOIN Disease dis ON diag.disease_id = dis.disease_id
    
    LEFT JOIN Treatment_Sessions ts ON tp.plan_id = ts.plan_id AND ts.status = 'Completed'
    
    WHERE tp.plan_id = @plan_id
    GROUP BY tp.plan_id, p.first_name, p.last_name, dis.disease_name, tp.treatment_type, tp.start_date, tp.end_date, tp.session_duration, tp.status;
END;
GO

-- ============================================
-- REPORT 6: Common Diseases Statistics
-- ============================================
CREATE OR ALTER PROCEDURE sp_CommonDiseasesStatistics
AS
BEGIN
    SELECT TOP 20
        dis.disease_name,
        dis.category,
        COUNT(diag.diagnosis_id) AS total_diagnoses,
        
        SUM(CASE WHEN diag.severity = 'Severe' THEN 1 ELSE 0 END) AS severe_cases,
        CAST(AVG(CAST(DATEDIFF(YEAR, p.date_of_birth, GETDATE()) AS FLOAT)) AS DECIMAL(5,1)) AS avg_patient_age
        
    FROM Disease dis
    LEFT JOIN Diagnosis diag ON dis.disease_id = diag.disease_id
    LEFT JOIN Patient p ON diag.patient_id = p.patient_id
    
    GROUP BY dis.disease_id, dis.disease_name, dis.category
    ORDER BY total_diagnoses DESC;
END;
GO

-- ============================================
-- REPORT 7: Patients with Specific Disease
-- ============================================
CREATE OR ALTER PROCEDURE sp_PatientsWithSpecificDisease
    @disease_id INT
AS
BEGIN
    SELECT 
        dis.disease_name,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        p.gender,
        diag.diagnosis_date,
        diag.severity,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name
        
    FROM Disease dis
    JOIN Diagnosis diag ON dis.disease_id = diag.disease_id
    JOIN Patient p ON diag.patient_id = p.patient_id
    JOIN Doctor d ON diag.doctor_id = d.doctor_id
    
    WHERE dis.disease_id = @disease_id
    ORDER BY diag.diagnosis_date DESC;
END;
GO

-- ============================================
-- REPORT 8: Doctor's Workload Report
-- ============================================
CREATE OR ALTER PROCEDURE sp_DoctorWorkloadReport
AS
BEGIN
    SELECT 
        d.doctor_id,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name,
        d.specialization,
        
        COUNT(DISTINCT a.appointment_id) AS total_appointments,
        COUNT(DISTINCT diag.diagnosis_id) AS total_diagnoses,
        COUNT(DISTINCT tp.plan_id) AS active_treatment_plans,
        
        -- Recent Activity
        MAX(a.appointment_date) AS last_activity
        
    FROM Doctor d
    LEFT JOIN Appointment a ON d.doctor_id = a.doctor_id
    LEFT JOIN Diagnosis diag ON d.doctor_id = diag.doctor_id
    LEFT JOIN Treatment_Plan tp ON d.doctor_id = tp.doctor_id AND tp.status = 'Active'
    
    GROUP BY d.doctor_id, d.first_name, d.last_name, d.specialization
    ORDER BY total_appointments DESC;
END;
GO

-- ============================================
-- REPORT 9: Expired Prescriptions
-- ============================================
CREATE OR ALTER PROCEDURE sp_ExpiredPrescriptions
    @days_ahead INT = 7
AS
BEGIN
    SELECT 
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        m.medication_name,
        pm.end_date,
        DATEDIFF(DAY, GETDATE(), pm.end_date) AS days_until_expiry,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name
        
    FROM Prescription_Medication pm
    JOIN Prescription pr ON pm.prescription_id = pr.prescription_id
    JOIN Medication m ON pm.medication_id = m.medication_id
    JOIN Patient p ON pr.patient_id = p.patient_id
    JOIN Doctor d ON pr.doctor_id = d.doctor_id
    
    WHERE DATEDIFF(DAY, GETDATE(), pm.end_date) BETWEEN 0 AND @days_ahead
    ORDER BY pm.end_date ASC;
END;
GO

-- ============================================
-- REPORT 10: Diagnosis Accuracy Report
-- Checks if diagnosis is confirmed by Lab Tests
-- ============================================
CREATE OR ALTER PROCEDURE sp_DiagnosisAccuracyReport
    @doctor_id INT = NULL
AS
BEGIN
    WITH DiagnosisStatus AS (
        SELECT 
            diag.doctor_id,
            diag.diagnosis_id,
            CASE 
                WHEN EXISTS (
                    SELECT 1 FROM Medical_Test mt 
                    WHERE mt.patient_id = diag.patient_id 
                    AND mt.test_date >= diag.diagnosis_date
                    AND mt.result_status IN ('Abnormal', 'Critical')
                ) THEN 1 
                ELSE 0 
            END AS is_confirmed
        FROM Diagnosis diag
    )
    SELECT 
        d.doctor_id,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name,
        d.specialization,
        
        COUNT(ds.diagnosis_id) AS total_diagnoses,
        
        SUM(ds.is_confirmed) AS confirmed_cases,
        
        CAST(
            SUM(ds.is_confirmed) * 100.0 / NULLIF(COUNT(ds.diagnosis_id), 0)
        AS DECIMAL(5,2)) AS confirmation_rate
        
    FROM Doctor d
    LEFT JOIN DiagnosisStatus ds ON d.doctor_id = ds.doctor_id
    
    WHERE @doctor_id IS NULL OR d.doctor_id = @doctor_id
    GROUP BY d.doctor_id, d.first_name, d.last_name, d.specialization
    ORDER BY confirmation_rate DESC;
END;
GO

-- ============================================
-- REPORT 11: Laboratory & Imaging Workload
-- ============================================
CREATE OR ALTER PROCEDURE sp_TestingFacilityWorkload
AS
BEGIN
    SELECT 
        tc.test_name,
        tc.test_category,
        COUNT(mt.test_id) AS total_performed,
        
        SUM(CASE WHEN mt.result_status = 'Abnormal' THEN 1 ELSE 0 END) AS abnormal_count,
        CAST(AVG(DATEDIFF(DAY, mt.test_date, GETDATE())) AS DECIMAL(5,1)) AS avg_days_since_test
        
    FROM Test_Catalog tc
    LEFT JOIN Medical_Test mt ON tc.test_catalog_id = mt.test_catalog_id
    GROUP BY tc.test_name, tc.test_category
    ORDER BY total_performed DESC;
END;
GO

-- ============================================
-- REPORT 12: Active Treatment Plans
-- Fixed: Links Treatment_Plan -> Diagnosis -> Disease
-- ============================================
CREATE OR ALTER PROCEDURE sp_ActiveTreatmentPlans
AS
BEGIN
    SELECT 
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        dis.disease_name, -- Fixed link
        tp.treatment_type,
        tp.end_date,
        DATEDIFF(DAY, GETDATE(), tp.end_date) AS days_remaining,
        tp.progress_notes,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor_name
        
    FROM Treatment_Plan tp
    JOIN Patient p ON tp.patient_id = p.patient_id
    JOIN Doctor d ON tp.doctor_id = d.doctor_id
    -- JOIN CHAIN FIX:
    JOIN Diagnosis diag ON tp.diagnosis_id = diag.diagnosis_id
    JOIN Disease dis ON diag.disease_id = dis.disease_id
    
    WHERE tp.status = 'Active'
    ORDER BY tp.end_date ASC;
END;
GO

-- ============================================
-- REPORT 13: Specialists & Interpreter Performance
-- ============================================
CREATE OR ALTER PROCEDURE sp_SpecialistPerformance
AS
BEGIN
    SELECT 
        CONCAT(d.first_name, ' ', d.last_name) AS specialist_name,
        d.specialization,
        COUNT(mt.test_id) AS tests_interpreted,
        MAX(mt.test_date) AS last_interpretation
        
    FROM Doctor d
    JOIN Medical_Test mt ON d.doctor_id = mt.interpreted_by
    GROUP BY d.doctor_id, d.first_name, d.last_name, d.specialization
    ORDER BY tests_interpreted DESC;
END;
GO

-- ============================================
-- REPORT 14: Severe Cases Alert
-- ============================================
CREATE OR ALTER PROCEDURE sp_SevereCasesAlert
AS
BEGIN
    SELECT 
        p.patient_id,
        CONCAT(p.first_name, ' ', p.last_name) AS patient_name,
        dis.disease_name,
        diag.severity,
        diag.notes,
        CONCAT(d.first_name, ' ', d.last_name) AS doctor,
        DATEDIFF(DAY, diag.diagnosis_date, GETDATE()) AS days_since_diagnosis
        
    FROM Diagnosis diag
    JOIN Patient p ON diag.patient_id = p.patient_id
    JOIN Disease dis ON diag.disease_id = dis.disease_id
    JOIN Doctor d ON diag.doctor_id = d.doctor_id
    
    WHERE diag.severity = 'Severe'
    ORDER BY diag.diagnosis_date DESC;
END;
GO