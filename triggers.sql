-- 1. Disease
CREATE OR ALTER TRIGGER trg_Update_Disease ON Disease AFTER UPDATE AS
BEGIN UPDATE Disease SET updated_at = GETDATE() FROM Disease t JOIN inserted i ON t.disease_id = i.disease_id END;
GO

-- 2. Doctor
CREATE OR ALTER TRIGGER trg_Update_Doctor ON Doctor AFTER UPDATE AS
BEGIN UPDATE Doctor SET updated_at = GETDATE() FROM Doctor t JOIN inserted i ON t.doctor_id = i.doctor_id END;
GO

-- 3. Patient
CREATE OR ALTER TRIGGER trg_Update_Patient ON Patient AFTER UPDATE AS
BEGIN UPDATE Patient SET updated_at = GETDATE() FROM Patient t JOIN inserted i ON t.patient_id = i.patient_id END;
GO

-- 4. Medication
CREATE OR ALTER TRIGGER trg_Update_Medication ON Medication AFTER UPDATE AS
BEGIN UPDATE Medication SET updated_at = GETDATE() FROM Medication t JOIN inserted i ON t.medication_id = i.medication_id END;
GO

-- 5. Test_Catalog
CREATE OR ALTER TRIGGER trg_Update_Test_Catalog ON Test_Catalog AFTER UPDATE AS
BEGIN UPDATE Test_Catalog SET updated_at = GETDATE() FROM Test_Catalog t JOIN inserted i ON t.test_catalog_id = i.test_catalog_id END;
GO

-- 6. Diagnosis
CREATE OR ALTER TRIGGER trg_Update_Diagnosis ON Diagnosis AFTER UPDATE AS
BEGIN UPDATE Diagnosis SET updated_at = GETDATE() FROM Diagnosis t JOIN inserted i ON t.diagnosis_id = i.diagnosis_id END;
GO

-- 7. Appointment
CREATE OR ALTER TRIGGER trg_Update_Appointment ON Appointment AFTER UPDATE AS
BEGIN UPDATE Appointment SET updated_at = GETDATE() FROM Appointment t JOIN inserted i ON t.appointment_id = i.appointment_id END;
GO

-- 8. Medical_Test
CREATE OR ALTER TRIGGER trg_Update_Medical_Test ON Medical_Test AFTER UPDATE AS
BEGIN UPDATE Medical_Test SET updated_at = GETDATE() FROM Medical_Test t JOIN inserted i ON t.test_id = i.test_id END;
GO

-- 9. Prescription
CREATE OR ALTER TRIGGER trg_Update_Prescription ON Prescription AFTER UPDATE AS
BEGIN UPDATE Prescription SET updated_at = GETDATE() FROM Prescription t JOIN inserted i ON t.prescription_id = i.prescription_id END;
GO

-- 10. Treatment_Plan
CREATE OR ALTER TRIGGER trg_Update_Treatment_Plan ON Treatment_Plan AFTER UPDATE AS
BEGIN UPDATE Treatment_Plan SET updated_at = GETDATE() FROM Treatment_Plan t JOIN inserted i ON t.plan_id = i.plan_id END;
GO

-- 11. Treatment_Sessions
CREATE OR ALTER TRIGGER trg_Update_Treatment_Sessions ON Treatment_Sessions AFTER UPDATE AS
BEGIN UPDATE Treatment_Sessions SET updated_at = GETDATE() FROM Treatment_Sessions t JOIN inserted i ON t.session_id = i.session_id END;
GO

-- 12. Prescription_Medication
CREATE OR ALTER TRIGGER trg_Update_Prescription_Medication ON Prescription_Medication AFTER UPDATE AS
BEGIN UPDATE Prescription_Medication SET updated_at = GETDATE() FROM Prescription_Medication t JOIN inserted i ON t.prescription_medication_id = i.prescription_medication_id END;
GO

-- Secend senario

CREATE OR ALTER TRIGGER trg_Validate_Appointment
ON Appointment
AFTER INSERT, UPDATE
AS
BEGIN
    IF EXISTS (
        SELECT 1 FROM Appointment a
        JOIN inserted i ON a.doctor_id = i.doctor_id AND a.appointment_date = i.appointment_date
        WHERE a.appointment_id <> i.appointment_id AND a.status IN ('Confirmed', 'Scheduled')
    )
    BEGIN
        RAISERROR ('Validation Error: Doctor is already booked for this time slot.', 16, 1);
        ROLLBACK TRANSACTION;
    END
END;
GO


CREATE OR ALTER TRIGGER trg_Validate_Patient_DOB
ON Patient
AFTER INSERT, UPDATE
AS
BEGIN
    IF EXISTS (SELECT 1 FROM inserted WHERE date_of_birth > GETDATE())
    BEGIN
        RAISERROR ('Validation Error: Date of Birth cannot be in the future.', 16, 1);
        ROLLBACK TRANSACTION;
    END
END;
GO


CREATE OR ALTER TRIGGER trg_Validate_Prescription_Date
ON Prescription_Medication
AFTER INSERT, UPDATE
AS
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM inserted i
        JOIN Prescription p ON i.prescription_id = p.prescription_id
        WHERE i.end_date < p.prescribed_date
    )
    BEGIN
        RAISERROR ('Validation Error: Medication end date cannot be before prescribed date.', 16, 1);
        ROLLBACK TRANSACTION;
    END
END;
GO

-- senario 3

CREATE OR ALTER TRIGGER trg_Audit_Patient
ON Patient
AFTER UPDATE
AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO Audit_Log (table_name, record_id, action_type, changed_by, old_value, new_value)
    SELECT 'Patient', i.patient_id, 'UPDATE', SYSTEM_USER, 
           CONCAT('History: ', d.medical_history), CONCAT('History: ', i.medical_history)
    FROM inserted i JOIN deleted d ON i.patient_id = d.patient_id
    WHERE i.medical_history <> d.medical_history;
END;
GO

CREATE OR ALTER TRIGGER trg_Audit_Diagnosis
ON Diagnosis
AFTER UPDATE
AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO Audit_Log (table_name, record_id, action_type, changed_by, old_value, new_value)
    SELECT 'Diagnosis', i.diagnosis_id, 'UPDATE', SYSTEM_USER, 
           CONCAT('Severity: ', d.severity), CONCAT('Severity: ', i.severity)
    FROM inserted i JOIN deleted d ON i.diagnosis_id = d.diagnosis_id
    WHERE i.severity <> d.severity;
END;
GO

CREATE OR ALTER TRIGGER trg_Audit_Doctor
ON Doctor
AFTER UPDATE
AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO Audit_Log (table_name, record_id, action_type, changed_by, old_value, new_value)
    SELECT 'Doctor', i.doctor_id, 'UPDATE', SYSTEM_USER, 
           CONCAT('Active: ', d.is_active), CONCAT('Active: ', i.is_active)
    FROM inserted i JOIN deleted d ON i.doctor_id = d.doctor_id
    WHERE i.is_active <> d.is_active;
END;
GO

-- senario 4 business logic

CREATE OR ALTER TRIGGER trg_Logic_CompleteTreatmentPlan
ON Treatment_Sessions
AFTER INSERT, UPDATE
AS
BEGIN
    SET NOCOUNT ON;
    DECLARE @plan_id INT = (SELECT DISTINCT plan_id FROM inserted);

    IF (SELECT COUNT(*) FROM Treatment_Sessions WHERE plan_id = @plan_id AND status = 'Completed') >= 
       (SELECT session_count FROM Treatment_Plan WHERE plan_id = @plan_id)
    BEGIN
        UPDATE Treatment_Plan 
        SET status = 'Completed', progress_notes = ISNULL(progress_notes, '') + ' [System: All sessions completed]'
        WHERE plan_id = @plan_id AND status = 'Active';
    END
END;
GO

CREATE OR ALTER TRIGGER trg_Logic_CriticalTestAlert
ON Medical_Test
AFTER INSERT, UPDATE
AS
BEGIN
    SET NOCOUNT ON;
    IF EXISTS (SELECT 1 FROM inserted WHERE result_status = 'Critical')
    BEGIN
        UPDATE Patient
        SET medical_history = ISNULL(medical_history, '') + ' [ALERT: Critical Lab Result on ' + CAST(GETDATE() AS NVARCHAR(20)) + ']'
        FROM Patient p
        JOIN inserted i ON p.patient_id = i.patient_id
        WHERE i.result_status = 'Critical';
    END
END;
GO

--==========
DECLARE @TableName NVARCHAR(255);
DECLARE @TriggerName NVARCHAR(255);
DECLARE @SQL NVARCHAR(MAX);

DECLARE TableCursor CURSOR FOR 
SELECT TABLE_NAME 
FROM INFORMATION_SCHEMA.TABLES 
WHERE TABLE_TYPE = 'BASE TABLE' 
  AND TABLE_NAME NOT IN ('Audit_Log', 'sysdiagrams');

OPEN TableCursor;
FETCH NEXT FROM TableCursor INTO @TableName;

WHILE @@FETCH_STATUS = 0
BEGIN
    SET @TriggerName = 'trg_Audit_Auto_' + @TableName;
    
    SET @SQL = '
    CREATE OR ALTER TRIGGER ' + @TriggerName + '
    ON [' + @TableName + ']
    AFTER INSERT, UPDATE, DELETE
    AS
    BEGIN
        SET NOCOUNT ON;
        DECLARE @Action NVARCHAR(20); -- تغییر سایز به 20 جهت اطمینان
        
        -- تشخیص نوع عملیات
        IF EXISTS (SELECT * FROM inserted) AND EXISTS (SELECT * FROM deleted)
            SET @Action = ''UPDATE'';
        ELSE IF EXISTS (SELECT * FROM inserted)
            SET @Action = ''INSERT'';
        ELSE IF EXISTS (SELECT * FROM deleted)
            SET @Action = ''DELETE'';
        ELSE
            RETURN; -- هیچ ردیفی تحت تاثیر قرار نگرفته

        -- درج در جدول لاگ (نام ستون‌ها اصلاح شد)
        INSERT INTO Audit_Log (table_name, action_type, changed_by, old_value, new_value)
        SELECT 
            ''' + @TableName + ''', 
            @Action, 
            SYSTEM_USER,
            (SELECT * FROM deleted FOR JSON AUTO), -- ذخیره در old_value
            (SELECT * FROM inserted FOR JSON AUTO); -- ذخیره در new_value
    END;';

    PRINT 'Creating trigger for table: ' + @TableName;
    EXEC sp_executesql @SQL;

    FETCH NEXT FROM TableCursor INTO @TableName;
END

CLOSE TableCursor;
DEALLOCATE TableCursor;
PRINT 'Done! All tables are now audited.';