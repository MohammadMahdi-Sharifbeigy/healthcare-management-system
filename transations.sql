CREATE OR ALTER PROCEDURE sp_CreateTreatmentPlanWithSessions
    @patient_id INT,
    @doctor_id INT,
    @diagnosis_id INT,
    @treatment_type NVARCHAR(150),
    @start_date DATE,
    @session_count INT,
    @goals NVARCHAR(MAX)
AS
BEGIN
    SET NOCOUNT ON;
    
    BEGIN TRANSACTION;

    BEGIN TRY
        DECLARE @new_plan_id INT;
        
        INSERT INTO Treatment_Plan (patient_id, doctor_id, diagnosis_id, treatment_type, start_date, session_count, status, goals)
        VALUES (@patient_id, @doctor_id, @diagnosis_id, @treatment_type, @start_date, @session_count, 'Active', @goals);
        
        SET @new_plan_id = SCOPE_IDENTITY(); 
        DECLARE @counter INT = 1;
        DECLARE @current_date DATE = @start_date;

        WHILE @counter <= @session_count
        BEGIN
            INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes)
            VALUES (@new_plan_id, @counter, @current_date, 'Scheduled', 'Auto-generated session');

            -- IF @counter = 3 RAISERROR('Simulated Error!', 16, 1);

            SET @current_date = DATEADD(DAY, 7, @current_date); 
            SET @counter = @counter + 1;
        END

        COMMIT TRANSACTION;
        PRINT 'Success: Plan and ' + CAST(@session_count AS NVARCHAR) + ' sessions created successfully.';
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;
            
        PRINT 'Error: Operation failed. Changes rolled back.';
        PRINT 'Error Message: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO


--===============
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

        INSERT INTO Prescription_Medication (
            prescription_id, 
            medication_id, 
            end_date, 
            frequency, 
            dosage, 
            instructions
        )
        SELECT 
            @new_presc_id, 
            Med.medication_id, 
            DATEADD(DAY, Med.duration_days, @prescribed_date), 
            Med.frequency, 
            Med.dosage, 
            Med.instructions
        FROM OPENJSON(@MedicationsJson) 
        WITH (
            medication_id INT '$.med_id',
            duration_days INT '$.days',
            frequency NVARCHAR(50) '$.freq',
            dosage NVARCHAR(10) '$.dose',
            instructions NVARCHAR(MAX) '$.inst'
        ) AS Med;

        COMMIT TRANSACTION;
        PRINT 'Success: Prescription ID ' + CAST(@new_presc_id AS NVARCHAR) + ' created successfully.';
    END TRY
    BEGIN CATCH
        IF @@TRANCOUNT > 0
            ROLLBACK TRANSACTION;

        PRINT 'Error: Transaction Failed! Rolled back.';
        PRINT 'Error Message: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO

--test:
DECLARE @MyJsonList NVARCHAR(MAX);

SET @MyJsonList = N'[
    {
        "med_id": 1,
        "days": 30,
        "freq": "Daily",
        "dose": "250mg",
        "inst": "Take with water"
    },
    {
        "med_id": 4,
        "days": 10,
        "freq": "Twice Daily",
        "dose": "500mg",
        "inst": "After meals"
    }
]';

EXEC sp_CreatePrescription_JSON 
    @patient_id = 15234567, 
    @doctor_id = 10000001, 
    @prescribed_date = '2025-02-05',
    @MedicationsJson = @MyJsonList;