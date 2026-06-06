-- cursor 1
DECLARE @plan_id INT;
DECLARE @start_date DATE;
DECLARE @session_count INT;
DECLARE @counter INT;
DECLARE @current_date DATE;

DECLARE plan_cursor CURSOR FOR 
SELECT plan_id, start_date, session_count 
FROM Treatment_Plan 
WHERE status = 'Active' 
  AND plan_id NOT IN (SELECT DISTINCT plan_id FROM Treatment_Sessions);

OPEN plan_cursor;
FETCH NEXT FROM plan_cursor INTO @plan_id, @start_date, @session_count;

WHILE @@FETCH_STATUS = 0 -- 0,-1,-2
BEGIN
    SET @counter = 1;
    SET @current_date = @start_date;

    WHILE @counter <= @session_count
    BEGIN
        INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes)
        VALUES (@plan_id, @counter, @current_date, 'Scheduled', 'Auto-generated session');

        SET @current_date = DATEADD(DAY, 7, @current_date);
        SET @counter = @counter + 1;
    END

    FETCH NEXT FROM plan_cursor INTO @plan_id, @start_date, @session_count;
END

CLOSE plan_cursor;
DEALLOCATE plan_cursor;

-- Cursor2
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

WHILE @@FETCH_STATUS = 0
BEGIN
    SET @msg = N'Sending SMS to ' + @patient_name + N' (' + @phone + N'): Reminder for appointment with Dr. ' + @doc_name + N' on ' + CAST(@appt_time AS NVARCHAR(20));
    PRINT @msg; 

    FETCH NEXT FROM reminder_cursor INTO @patient_name, @phone, @appt_time, @doc_name;
END

CLOSE reminder_cursor;
DEALLOCATE reminder_cursor;

-- cursor 3
DECLARE @prescription_id INT;
DECLARE @patient_id INT;
DECLARE @medication_count INT;

DECLARE safety_check_cursor CURSOR FOR 
SELECT prescription_id, patient_id 
FROM Prescription 
WHERE prescribed_date = CAST(GETDATE() AS DATE);

OPEN safety_check_cursor;
FETCH NEXT FROM safety_check_cursor INTO @prescription_id, @patient_id;

WHILE @@FETCH_STATUS = 0
BEGIN
    SELECT @medication_count = COUNT(*) 
    FROM Prescription_Medication 
    WHERE prescription_id = @prescription_id;

    IF @medication_count > 5
    BEGIN
        PRINT 'ALERT: Patient ID ' + CAST(@patient_id AS NVARCHAR) + ' has a prescription with ' + CAST(@medication_count AS NVARCHAR) + ' items. Review needed.';
    END

    FETCH NEXT FROM safety_check_cursor INTO @prescription_id, @patient_id;
END

CLOSE safety_check_cursor;
DEALLOCATE safety_check_cursor;