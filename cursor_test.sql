INSERT INTO Treatment_Plan (patient_id, doctor_id, diagnosis_id, treatment_type, start_date, session_count, status, goals)
VALUES (15234567, 40000001, 1, 'Manual Therapy', GETDATE(), 5, 'Active', 'Test Auto Generation');

DECLARE @new_plan_id INT = SCOPE_IDENTITY(); 

DECLARE @plan_id INT;
DECLARE @start_date DATE;
DECLARE @session_count INT;
DECLARE @counter INT;
DECLARE @current_date DATE;

DECLARE plan_cursor CURSOR FOR 
SELECT plan_id, start_date, session_count 
FROM Treatment_Plan 
WHERE status = 'Active' 
  AND plan_id = @new_plan_id 
  AND plan_id NOT IN (SELECT DISTINCT plan_id FROM Treatment_Sessions);

OPEN plan_cursor;
FETCH NEXT FROM plan_cursor INTO @plan_id, @start_date, @session_count;

WHILE @@FETCH_STATUS = 0
BEGIN
    SET @counter = 1;
    SET @current_date = @start_date;

    WHILE @counter <= @session_count
    BEGIN
        INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes)
        VALUES (@plan_id, @counter, @current_date, 'Scheduled', 'Auto-generated session via Cursor');

        SET @current_date = DATEADD(DAY, 7, @current_date); 
        SET @counter = @counter + 1;
    END

    FETCH NEXT FROM plan_cursor INTO @plan_id, @start_date, @session_count;
END

CLOSE plan_cursor;
DEALLOCATE plan_cursor;

SELECT * FROM Treatment_Sessions WHERE plan_id = @new_plan_id;





-- =============================================
-- TEST 2: APPOINTMENT REMINDERS (Check 'Messages' Tab)
-- =============================================

INSERT INTO Appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate)
VALUES (20567890, 10000001, DATEADD(DAY, 1, CAST(GETDATE() AS DATE)), 'Checkup Reminder Test', 'Confirmed', '120/80', '80');

DECLARE @patient_name NVARCHAR(200);
DECLARE @phone NVARCHAR(20);
DECLARE @appt_time DATE;
DECLARE @doc_name NVARCHAR(200);
DECLARE @msg NVARCHAR(MAX);

DECLARE reminder_cursor CURSOR FOR 
SELECT 
    CONCAT(p.first_name, ' ', p.last_name),
    p.phone,
    a.appointment_date,
    CONCAT(d.first_name, ' ', d.last_name)
FROM Appointment a
JOIN Patient p ON a.patient_id = p.patient_id
JOIN Doctor d ON a.doctor_id = d.doctor_id
WHERE a.appointment_date = DATEADD(DAY, 1, CAST(GETDATE() AS DATE)) 
AND a.status = 'Confirmed';

OPEN reminder_cursor;
FETCH NEXT FROM reminder_cursor INTO @patient_name, @phone, @appt_time, @doc_name;

PRINT '       SENDING SMS REMINDERS (SIMULATION)         ';

WHILE @@FETCH_STATUS = 0
BEGIN
    SET @msg = N'>> SMS TO: ' + @patient_name + N' (' + @phone + N')' + CHAR(13) + 
               N'   MSG: Reminder for appointment with Dr. ' + @doc_name + N' on ' + CAST(@appt_time AS NVARCHAR(20));
    PRINT @msg; 
    PRINT '--------------------------------------------------';

    FETCH NEXT FROM reminder_cursor INTO @patient_name, @phone, @appt_time, @doc_name;
END

CLOSE reminder_cursor;
DEALLOCATE reminder_cursor;



-- =============================================
-- TEST 3: SAFETY CHECK (Polypharmacy)
-- =============================================

INSERT INTO Prescription (patient_id, doctor_id, prescribed_date)
VALUES (15234567, 10000001, CAST(GETDATE() AS DATE));

DECLARE @new_presc_id INT = SCOPE_IDENTITY();

INSERT INTO Prescription_Medication (prescription_id, medication_id, end_date, frequency, dosage, instructions) VALUES
(@new_presc_id, 1, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test'),
(@new_presc_id, 2, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test'),
(@new_presc_id, 3, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test'),
(@new_presc_id, 4, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test'),
(@new_presc_id, 5, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test'),
(@new_presc_id, 6, DATEADD(DAY, 30, GETDATE()), 'Daily', '10mg', 'Test');

DECLARE @prescription_id INT;
DECLARE @patient_id INT;
DECLARE @medication_count INT;

DECLARE safety_check_cursor CURSOR FOR 
SELECT prescription_id, patient_id 
FROM Prescription 
WHERE prescribed_date = CAST(GETDATE() AS DATE);

OPEN safety_check_cursor;
FETCH NEXT FROM safety_check_cursor INTO @prescription_id, @patient_id;

PRINT '           SAFETY AUDIT LOG                       ';

WHILE @@FETCH_STATUS = 0
BEGIN
    SELECT @medication_count = COUNT(*) 
    FROM Prescription_Medication 
    WHERE prescription_id = @prescription_id;

    IF @medication_count > 5
    BEGIN
        PRINT 'ALERT !!! Patient ID ' + CAST(@patient_id AS NVARCHAR) + 
              ' has a prescription (ID: ' + CAST(@prescription_id AS NVARCHAR) + 
              ') with ' + CAST(@medication_count AS NVARCHAR) + ' items. REVIEW NEEDED.';
    END
    ELSE
    BEGIN
         PRINT 'OK: Prescription ID ' + CAST(@prescription_id AS NVARCHAR) + ' is safe.';
    END

    FETCH NEXT FROM safety_check_cursor INTO @prescription_id, @patient_id;
END

CLOSE safety_check_cursor;
DEALLOCATE safety_check_cursor;