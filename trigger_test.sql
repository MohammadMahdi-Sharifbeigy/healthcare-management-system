PRINT '==================================================';
PRINT 'TEST 1: AUTOMATION TRIGGER (updated_at)';
PRINT '==================================================';

DECLARE @InitialTime DATETIME;
SELECT @InitialTime = updated_at FROM Patient WHERE patient_id = 15234567;
PRINT 'Old Updated_At: ' + ISNULL(CAST(@InitialTime AS NVARCHAR), 'NULL');

WAITFOR DELAY '00:00:02';

UPDATE Patient 
SET address = 'Tehran, New Address For Test' 
WHERE patient_id = 15234567;

DECLARE @NewTime DATETIME;
SELECT @NewTime = updated_at FROM Patient WHERE patient_id = 15234567;
PRINT 'New Updated_At: ' + CAST(@NewTime AS NVARCHAR);

IF @NewTime > @InitialTime
    PRINT '>>> TEST RESULT: PASSED (Date updated automatically)';
ELSE
    PRINT '>>> TEST RESULT: FAILED (Date did not change)';
GO


PRINT '==================================================';
PRINT 'TEST 2: VALIDATION TRIGGERS';
PRINT '==================================================';

PRINT '--- Testing Appointment Overlap ---';
BEGIN TRY
    INSERT INTO Appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate)
    VALUES (20567890, 10000001, '2025-01-15', 'Overlap Test', 'Confirmed', '120/80', '80');
    
    PRINT '>>> TEST RESULT: FAILED (System allowed duplicate booking!)';
END TRY
BEGIN CATCH
    PRINT '>>> TEST RESULT: PASSED (Error Caught: ' + ERROR_MESSAGE() + ')';
END CATCH;

PRINT '--- Testing Future DOB ---';
BEGIN TRY
    INSERT INTO Patient (patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type)
    VALUES (99999999, 'Test', 'User', DATEADD(YEAR, 1, GETDATE()), 'Male', 'future@test.com', '090000000', 'Tehran', 'None', 'O+');

    PRINT '>>> TEST RESULT: FAILED (System allowed future DOB!)';
END TRY
BEGIN CATCH
    PRINT '>>> TEST RESULT: PASSED (Error Caught: ' + ERROR_MESSAGE() + ')';
END CATCH;
GO

PRINT '==================================================';
PRINT 'TEST 3: AUDITING TRIGGERS';
PRINT '==================================================';

PRINT 'Updating Patient Medical History...';
UPDATE Patient 
SET medical_history = 'Diabetes, High Blood Pressure, Updated Log Test'
WHERE patient_id = 15234567;

PRINT 'Updating Doctor Status...';
UPDATE Doctor
SET is_active = 0
WHERE doctor_id = 10000004;

PRINT '--- Checking Audit_Log Table ---';
SELECT TOP 5 
    log_id, 
    table_name, 
    action_type, 
    old_value, 
    new_value, 
    change_date 
FROM Audit_Log 
ORDER BY log_id DESC;

GO

PRINT '==================================================';
PRINT 'TEST 4: BUSINESS LOGIC TRIGGERS';
PRINT '==================================================';

INSERT INTO Treatment_Plan (patient_id, doctor_id, diagnosis_id, treatment_type, start_date, session_count, status, goals, progress_notes)
VALUES (15234567, 40000001, 1, 'Test Logic Therapy', GETDATE(), 2, 'Active', 'Testing Auto Complete', 'Started');

DECLARE @NewPlanID INT = SCOPE_IDENTITY();
PRINT 'Created Plan ID: ' + CAST(@NewPlanID AS NVARCHAR);

INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes)
VALUES (@NewPlanID, 1, GETDATE(), 'Completed', 'First Session');

DECLARE @Status1 NVARCHAR(50);
SELECT @Status1 = status FROM Treatment_Plan WHERE plan_id = @NewPlanID;
PRINT 'Status after 1 session (Expected: Active): ' + @Status1;

INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes)
VALUES (@NewPlanID, 2, DATEADD(DAY, 7, GETDATE()), 'Completed', 'Second Session');

DECLARE @Status2 NVARCHAR(50);
DECLARE @Notes NVARCHAR(MAX);
SELECT @Status2 = status, @Notes = progress_notes FROM Treatment_Plan WHERE plan_id = @NewPlanID;

PRINT 'Status after 2 sessions (Expected: Completed): ' + @Status2;
PRINT 'Notes: ' + @Notes;

IF @Status2 = 'Completed'
    PRINT '>>> TEST RESULT: PASSED (Plan auto-completed)';
ELSE
    PRINT '>>> TEST RESULT: FAILED (Plan is still active)';
GO

--- TEST audit log
PRINT 'Testing Triggers...';

INSERT INTO Patient (patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type)
VALUES (999, 'Test', 'Robot', '2000-01-01', 'Male', 'robot@test.com', '000', 'Server', 'None', 'O+');

UPDATE Doctor 
SET phone = '09120000000' 
WHERE doctor_id = 10000001;

INSERT INTO Appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate)
VALUES (999, 10000001, GETDATE(), 'Test Delete', 'Scheduled', '12/8', '80');

DELETE FROM Appointment WHERE patient_id = 999;

DELETE FROM Patient WHERE patient_id = 999;

SELECT * FROM Audit_Log ORDER BY log_id DESC;