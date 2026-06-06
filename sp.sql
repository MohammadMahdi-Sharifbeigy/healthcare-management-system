CREATE OR ALTER PROCEDURE sp_AddPatient
    @patient_id INT,
    @first_name NVARCHAR(100),
    @last_name NVARCHAR(100),
    @date_of_birth DATE,
    @gender NVARCHAR(10),
    @email NVARCHAR(150), 
    @phone NVARCHAR(20),
    @address NVARCHAR(MAX),
    @emergency_contact NVARCHAR(150),
    @blood_type NVARCHAR(10)
AS
BEGIN
    INSERT INTO Patient (patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type)
    VALUES (@patient_id, @first_name, @last_name, @date_of_birth, @gender, @email, @phone, @address, @emergency_contact, @blood_type);
    
    SELECT 'Success' AS status, 'Patient added successfully' AS message;
END;
GO

CREATE OR ALTER PROCEDURE sp_AddAppointment
    @patient_id INT,
    @doctor_id INT,
    @appointment_date DATE,
    @reason NVARCHAR(255)
AS
BEGIN
    IF EXISTS (SELECT 1 FROM Appointment WHERE doctor_id = @doctor_id AND appointment_date = @appointment_date)
    BEGIN
        SELECT 'Error' AS status, 'Doctor is already booked for this date.' AS message;
        RETURN;
    END

    INSERT INTO Appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate)
    VALUES (@patient_id, @doctor_id, @appointment_date, @reason, 'Scheduled', 'N/A', 'N/A');
    
    SELECT 'Success' AS status, 'Appointment scheduled successfully' AS message;
END;
GO

CREATE OR ALTER PROCEDURE sp_UpdatePatient
    @patient_id INT,
    @phone NVARCHAR(20),
    @address NVARCHAR(MAX),
    @emergency_contact NVARCHAR(150)
AS
BEGIN
    UPDATE Patient
    SET phone = @phone,
        address = @address,
        emergency_contact = @emergency_contact,
        updated_at = GETDATE()
    WHERE patient_id = @patient_id;
    
    SELECT 'Success' AS status, 'Patient details updated successfully' AS message;
END;
GO

CREATE OR ALTER PROCEDURE sp_GetPatientById
    @patient_id INT
AS
BEGIN
    SELECT * FROM Patient WHERE patient_id = @patient_id;
END;
GO

CREATE OR ALTER PROCEDURE sp_CreatePrescription_JSON
    @patient_id INT,
    @doctor_id INT,
    @prescribed_date DATE,
    @MedicationsJson NVARCHAR(MAX) 
AS
BEGIN
    SET NOCOUNT ON;
    BEGIN TRANSACTION;
    BEGIN TRY
        DECLARE @new_presc_id INT;
        INSERT INTO Prescription (patient_id, doctor_id, prescribed_date)
        VALUES (@patient_id, @doctor_id, @prescribed_date);
        SET @new_presc_id = SCOPE_IDENTITY();

        INSERT INTO Prescription_Medication (prescription_id, medication_id, end_date, frequency, dosage, instructions)
        SELECT @new_presc_id, med_id, DATEADD(DAY, days, @prescribed_date), freq, dose, inst
        FROM OPENJSON(@MedicationsJson)
        WITH (med_id INT '$.med_id', days INT '$.days', freq NVARCHAR(50) '$.freq', dose NVARCHAR(10) '$.dose', inst NVARCHAR(MAX) '$.inst');

        COMMIT TRANSACTION;
        SELECT 'Success' AS status, CAST(@new_presc_id AS NVARCHAR) AS message;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        SELECT 'Error' AS status, ERROR_MESSAGE() AS message;
    END CATCH;
END;
GO

CREATE OR ALTER PROCEDURE sp_UpdatePrescription_Full
    @prescription_id INT,
    @patient_id INT,
    @doctor_id INT,
    @prescribed_date DATE,
    @MedicationsJson NVARCHAR(MAX)
AS
BEGIN
    SET NOCOUNT ON;
    BEGIN TRANSACTION;
    BEGIN TRY
        UPDATE Prescription 
        SET patient_id = @patient_id, doctor_id = @doctor_id, prescribed_date = @prescribed_date, updated_at = GETDATE()
        WHERE prescription_id = @prescription_id;

        DELETE FROM Prescription_Medication WHERE prescription_id = @prescription_id;

        INSERT INTO Prescription_Medication (prescription_id, medication_id, end_date, frequency, dosage, instructions)
        SELECT @prescription_id, med_id, DATEADD(DAY, days, @prescribed_date), freq, dose, inst
        FROM OPENJSON(@MedicationsJson)
        WITH (med_id INT '$.med_id', days INT '$.days', freq NVARCHAR(50) '$.freq', dose NVARCHAR(10) '$.dose', inst NVARCHAR(MAX) '$.inst');

        COMMIT TRANSACTION;
        SELECT 'Success' AS status, 'Prescription updated successfully' AS message;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        SELECT 'Error' AS status, ERROR_MESSAGE() AS message;
    END CATCH;
END;
GO


CREATE OR ALTER PROCEDURE sp_GetAllMedications
AS
BEGIN
    SELECT medication_id, medication_name FROM Medication ORDER BY medication_name;
END;
GO

CREATE OR ALTER PROCEDURE sp_GetPrescriptionById
    @prescription_id INT
AS
BEGIN
    SELECT * FROM Prescription WHERE prescription_id = @prescription_id;

    SELECT 
        pm.prescription_medication_id,
        pm.medication_id,
        m.medication_name,
        pm.dosage,
        pm.frequency,
        pm.instructions,
        DATEDIFF(DAY, (SELECT prescribed_date FROM Prescription WHERE prescription_id = @prescription_id), pm.end_date) as duration_days
    FROM Prescription_Medication pm
    JOIN Medication m ON pm.medication_id = m.medication_id
    WHERE pm.prescription_id = @prescription_id;
END;
GO