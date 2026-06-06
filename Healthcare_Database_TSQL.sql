    DROP TABLE IF EXISTS Treatment_Sessions;
    DROP TABLE IF EXISTS Prescription_Medication;
    DROP TABLE IF EXISTS Prescription;
    DROP TABLE IF EXISTS Medical_Test;
    DROP TABLE IF EXISTS Treatment_Plan;
    DROP TABLE IF EXISTS Appointment;
    DROP TABLE IF EXISTS Diagnosis;
    DROP TABLE IF EXISTS Patient;
    DROP TABLE IF EXISTS Doctor;
    DROP TABLE IF EXISTS Medication;
    DROP TABLE IF EXISTS Test_Catalog;
    DROP TABLE IF EXISTS Disease;

    CREATE TABLE Disease (
        disease_id INT PRIMARY KEY IDENTITY(1,1),
        disease_name NVARCHAR(150) NOT NULL,
        category NVARCHAR(100) NOT NULL,
        icd_code NVARCHAR(20) UNIQUE NOT NULL,
        description NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1
    );

    CREATE TABLE Doctor (
        doctor_id INT PRIMARY KEY,
        first_name NVARCHAR(100) NOT NULL,
        last_name NVARCHAR(100) NOT NULL,
        email NVARCHAR(150) UNIQUE NOT NULL,
        phone NVARCHAR(20) NOT NULL,
        license_number NVARCHAR(50) UNIQUE NOT NULL,
        specialization NVARCHAR(100) NOT NULL,
        department NVARCHAR(100) NOT NULL,
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1
    );

    CREATE TABLE Patient (
        patient_id INT PRIMARY KEY,
        first_name NVARCHAR(100) NOT NULL,
        last_name NVARCHAR(100) NOT NULL,
        date_of_birth DATE NOT NULL,
        gender NVARCHAR(10) NOT NULL,
        email NVARCHAR(150) UNIQUE NOT NULL,
        phone NVARCHAR(20) NOT NULL,
        address NVARCHAR(MAX) NOT NULL,
        emergency_contact NVARCHAR(150) NOT NULL,
        blood_type NVARCHAR(10) NOT NULL,
        medical_history NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1
    );

    

    CREATE TABLE Medication (
        medication_id INT PRIMARY KEY IDENTITY(1,1),
        medication_name NVARCHAR(150) NOT NULL,
        generic_name NVARCHAR(150) NOT NULL,
        form NVARCHAR(50) NOT NULL,
        strength NVARCHAR(50) NOT NULL,
        manufacturer NVARCHAR(150) NOT NULL,
        side_effects NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1
    );

    CREATE TABLE Test_Catalog (
        test_catalog_id INT PRIMARY KEY IDENTITY(1,1),
        test_name NVARCHAR(150) NOT NULL,
        test_category NVARCHAR(100) NOT NULL,
        description NVARCHAR(MAX),
        normal_range NVARCHAR(100),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1
    );

    CREATE TABLE Diagnosis (
        diagnosis_id INT PRIMARY KEY IDENTITY(1,1),
        patient_id INT NOT NULL,
        disease_id INT NOT NULL,
        doctor_id INT NOT NULL,
        diagnosis_date DATE NOT NULL,
        severity NVARCHAR(50) NOT NULL,
        notes NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
        FOREIGN KEY (disease_id) REFERENCES Disease(disease_id),
        FOREIGN KEY (doctor_id) REFERENCES Doctor(doctor_id),
        UNIQUE(patient_id, disease_id, doctor_id, diagnosis_date)
    );

    CREATE TABLE Appointment (
        appointment_id INT PRIMARY KEY IDENTITY(1,1),
        patient_id INT NOT NULL,
        doctor_id INT NOT NULL,
        appointment_date DATE NOT NULL,
        reason NVARCHAR(255) NOT NULL,
        status NVARCHAR(50) NOT NULL,
        blood_pressure NVARCHAR(25) NOT NULL,
        heart_rate NVARCHAR(25) NOT NULL,
        duration_minutes INT,
        notes NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
        FOREIGN KEY (doctor_id) REFERENCES Doctor(doctor_id)
    );

    CREATE TABLE Medical_Test (
        test_id INT PRIMARY KEY IDENTITY(1,1),
        patient_id INT NOT NULL,
        test_catalog_id INT NOT NULL,
        test_date DATE NOT NULL,
        result_value NVARCHAR(100),
        normal_range NVARCHAR(100),
        findings NVARCHAR(MAX),
        result_status NVARCHAR(50),
        image_url NVARCHAR(MAX),
        interpreted_by INT,
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
        FOREIGN KEY (test_catalog_id) REFERENCES Test_Catalog(test_catalog_id),
        FOREIGN KEY (interpreted_by) REFERENCES Doctor(doctor_id)
    );
    CREATE TABLE Prescription (
        prescription_id INT PRIMARY KEY IDENTITY(1,1),
        patient_id INT NOT NULL,
        doctor_id INT NOT NULL,
        prescribed_date DATE NOT NULL,
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
        FOREIGN KEY (doctor_id) REFERENCES Doctor(doctor_id)
    );

    CREATE TABLE Treatment_Plan (
        plan_id INT PRIMARY KEY IDENTITY(1,1),
        patient_id INT NOT NULL,
        doctor_id INT NOT NULL,
        diagnosis_id INT NOT NULL,
        treatment_type NVARCHAR(150) NOT NULL,
        start_date DATE NOT NULL,
        end_date DATE,
        session_duration INT,
        session_count INT NOT NULL,
        status NVARCHAR(50) NOT NULL,
        goals NVARCHAR(MAX) NOT NULL,
        progress_notes NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        is_active BIT DEFAULT 1,
        FOREIGN KEY (patient_id) REFERENCES Patient(patient_id),
        FOREIGN KEY (doctor_id) REFERENCES Doctor(doctor_id),
        FOREIGN KEY (diagnosis_id) REFERENCES Diagnosis(diagnosis_id)
    );

    CREATE TABLE Treatment_Sessions (
        session_id INT PRIMARY KEY IDENTITY(1,1),
        plan_id INT NOT NULL,
        session_number INT,     
        session_date DATE NOT NULL,
        status NVARCHAR(50), 
        notes NVARCHAR(MAX),
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (plan_id) REFERENCES Treatment_Plan(plan_id)
    );

    CREATE TABLE Prescription_Medication (
        prescription_medication_id INT PRIMARY KEY IDENTITY(1,1),
        prescription_id INT NOT NULL,
        medication_id INT NOT NULL,
        end_date DATE NOT NULL,
        frequency NVARCHAR(50) NOT NULL,
        dosage NVARCHAR(10) NOT NULL,
        instructions NVARCHAR(MAX) NOT NULL,
        created_at DATETIME DEFAULT GETDATE(),
        updated_at DATETIME,
        FOREIGN KEY (prescription_id) REFERENCES Prescription(prescription_id),
        FOREIGN KEY (medication_id) REFERENCES Medication(medication_id),
        UNIQUE(prescription_id, medication_id)
    );

    IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Audit_Log' AND xtype='U')
    CREATE TABLE Audit_Log (
        log_id INT PRIMARY KEY IDENTITY(1,1),
        table_name NVARCHAR(50),
        record_id INT,
        action_type NVARCHAR(20),
        changed_by NVARCHAR(100),
        change_date DATETIME DEFAULT GETDATE(),
        old_value NVARCHAR(MAX),
        new_value NVARCHAR(MAX)
    );

    -- Patient indexes
    CREATE INDEX idx_patient_email ON Patient(email);
    CREATE INDEX idx_patient_phone ON Patient(phone);
    CREATE INDEX idx_patient_date_of_birth ON Patient(date_of_birth);

    -- Doctor indexes
    CREATE INDEX idx_doctor_email ON Doctor(email);
    CREATE INDEX idx_doctor_license ON Doctor(license_number);
    CREATE INDEX idx_doctor_specialization ON Doctor(specialization);

    -- Diagnosis indexes
    CREATE INDEX idx_diagnosis_patient ON Diagnosis(patient_id);
    CREATE INDEX idx_diagnosis_disease ON Diagnosis(disease_id);
    CREATE INDEX idx_diagnosis_date ON Diagnosis(diagnosis_date);

    -- Appointment indexes
    CREATE INDEX idx_appointment_patient ON Appointment(patient_id);
    CREATE INDEX idx_appointment_doctor ON Appointment(doctor_id);
    CREATE INDEX idx_appointment_date ON Appointment(appointment_date);

    -- Test indexes
    CREATE INDEX idx_medical_test_catalog ON Medical_Test(test_catalog_id);
    CREATE INDEX idx_medical_test_date ON Medical_Test(test_date);
    CREATE INDEX idx_medical_test_patient ON Medical_Test(patient_id); 

    -- Prescription indexes
    CREATE INDEX idx_prescription_patient ON Prescription(patient_id);
    CREATE INDEX idx_prescription_doctor ON Prescription(doctor_id);

    -- Treatment Plan indexes
    CREATE INDEX idx_treatment_plan_patient ON Treatment_Plan(patient_id);
    CREATE INDEX idx_treatment_plan_doctor ON Treatment_Plan(doctor_id);
    CREATE INDEX idx_treatment_plan_diagnosis ON Treatment_Plan(diagnosis_id);

    -- Treatment Sessions
CREATE INDEX idx_treatment_sessions_plan ON Treatment_Sessions(plan_id);
CREATE INDEX idx_treatment_sessions_date ON Treatment_Sessions(session_date);