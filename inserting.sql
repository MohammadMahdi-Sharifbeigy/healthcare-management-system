    INSERT INTO Disease (disease_name, category, icd_code, description) VALUES
    ('Alzheimers Disease', 'Neurodegenerative Disorders', 'G30', 'Neurodegenerative disease leading to memory loss and cognitive decline'),
    ('Parkinsons Disease', 'Movement Disorders', 'G20', 'Progressive movement disorder characterized by tremor and rigidity'),
    ('Epilepsy', 'Seizure Disorders', 'G40', 'Neurological disorder characterized by seizures and loss of consciousness'),
    ('Multiple Sclerosis', 'Neurodegenerative Disorders', 'G35', 'Autoimmune neurodegenerative disease affecting the nervous system'),
    ('Migraine', 'Neurological Headaches', 'G43', 'Severe recurring headache disorder often with neurological symptoms'),
    ('Febrile Seizure', 'Seizure Disorders', 'R56', 'Convulsions triggered by high fever in children'),
    ('Diabetic Neuropathy', 'Peripheral Nerve Diseases', 'G63', 'Nerve damage caused by diabetes affecting peripheral nerves'),
    ('Stroke', 'Cerebrovascular Diseases', 'I63', 'Blockage of blood vessels in the brain'),
    ('Hemophilia', 'Bleeding Disorders', 'D66', 'Blood clotting disorder resulting in excessive bleeding'),
    ('Head Trauma', 'Head Injuries', 'S06', 'Injury to brain tissue from head trauma'),
    ('Coronary Artery Disease', 'Cardiovascular', 'I25.1', 'Narrowing or blockage of the coronary arteries'),
    ('Hypertension', 'Cardiovascular', 'I10', 'High blood pressure');

    -- ALGORITHM EXPLANATION FOR DOCTOR_ID:
    -- Format: 8-digit integer (XXYYYYYY)
    -- XX = Specialization Code:
    --   10 = Neurology
    --   20 = Psychiatry
    --   30 = Radiology
    --   40 = Physiotherapy
    --   50 = Psychology
    --   60 = Laboratory
    --   70 = Cardiology
    -- YYYYYY = Sequential number within specialization

    INSERT INTO Doctor (doctor_id, first_name, last_name, email, phone, license_number, specialization, department) VALUES
    (10000001, 'John', 'Smith', 'john.smith@hospital.com', '09121234567', 'MED001', 'Neurology', 'Neurology'),
    (10000002, 'Sarah', 'Johnson', 'sarah.johnson@hospital.com', '09121234568', 'MED002', 'Neurology', 'Neurology'),
    (20000001, 'Michael', 'Davis', 'michael.davis@hospital.com', '09121234569', 'MED003', 'Psychiatry', 'Psychiatry'),
    (30000001, 'Emily', 'Wilson', 'emily.wilson@hospital.com', '09121234570', 'MED004', 'Radiology', 'Radiology'),
    (40000001, 'James', 'Moore', 'james.moore@hospital.com', '09121234571', 'MED005', 'Physiotherapy', 'Physiotherapy'),
    (10000003, 'Lisa', 'Taylor', 'lisa.taylor@hospital.com', '09121234572', 'MED006', 'Neurology', 'Neurology'),
    (50000001, 'David', 'Anderson', 'david.anderson@hospital.com', '09121234573', 'MED007', 'Psychology', 'Psychology'),
    (60000001, 'Jessica', 'Thomas', 'jessica.thomas@hospital.com', '09121234574', 'MED008', 'Laboratory', 'Laboratory'),
    (10000004, 'Robert', 'Martinez', 'robert.martinez@hospital.com', '09121234575', 'MED009', 'Neurology', 'Neurology'),
    (30000002, 'Jennifer', 'Garcia', 'jennifer.garcia@hospital.com', '09121234576', 'MED010', 'Radiology', 'Radiology'),
    (70000001, 'Reza', 'Karimi', 'reza.karimi@hospital.com', '09121234999', 'MED011', 'Cardiology', 'Cardiology');

    INSERT INTO Patient (patient_id, first_name, last_name, date_of_birth, gender, email, phone, address, emergency_contact, blood_type, medical_history) VALUES
    (15234567, 'Alexander', 'Brown', '1960-05-15', 'Male', 'alexander.brown@email.com', '09121111111', 'Tehran, Enghelab Street', 'Michael Turner - 09121111112', 'O+', 'Diabetes, High Blood Pressure'),
    (20567890, 'Anna', 'Miller', '1975-08-22', 'Female', 'anna.miller@email.com', '09121111113', 'Tehran, Vali-Asr Avenue', 'Jennifer Blake - 09121111114', 'A+', 'No significant history'),
    (18234123, 'Andrew', 'Nelson', '1970-03-10', 'Male', 'andrew.nelson@email.com', '09121111115', 'Tehran, Pahlavi Street', 'Robert King - 09121111116', 'B+', 'Arthritis, High Cholesterol'),
    (22891234, 'Amanda', 'Rodriguez', '1980-11-05', 'Female', 'amanda.rodriguez@email.com', '09121111117', 'Tehran, Shariati Street', 'Sarah Bennett - 09121111118', 'AB+', 'Epilepsy'),
    (16756234, 'Marcus', 'Lewis', '1965-07-18', 'Male', 'marcus.lewis@email.com', '09121111119', 'Tehran, Koosar Street', 'David Thompson - 09121111120', 'O-', 'Asthma, Allergies'),
    (23421567, 'Rachel', 'Lee', '1982-02-14', 'Female', 'rachel.lee@email.com', '09121111121', 'Tehran, Firdousi Street', 'Lisa Patterson - 09121111122', 'A-', 'No significant history'),
    (14123890, 'Christopher', 'Walker', '1958-09-25', 'Male', 'christopher.walker@email.com', '09121111123', 'Tehran, Laleh Zar Street', 'James Stevens - 09121111124', 'B-', 'Parkinsons Disease'),
    (21345678, 'Sophia', 'Hall', '1976-04-08', 'Female', 'sophia.hall@email.com', '09121111125', 'Tehran, Niavaran Street', 'Patricia Johnson - 09121111126', 'AB-', 'Chronic Migraine'),
    (15678912, 'William', 'Allen', '1963-12-30', 'Male', 'william.allen@email.com', '09121111127', 'Tehran, Chalus Street', 'Christopher Martin - 09121111128', 'O+', 'Previous Stroke'),
    (24567890, 'Victoria', 'Young', '1985-06-17', 'Female', 'victoria.young@email.com', '09121111129', 'Tehran, Resalat Street', 'Amanda Foster - 09121111130', 'A+', 'No significant history'),
    (30998877, 'Hassan', 'Rezazadeh', '1955-01-20', 'Male', 'hassan.reza@email.com', '09129998888', 'Tehran, Vanak Square', 'Maryam Reza - 09129997777', 'B+', 'Heart Arrhythmia');

    INSERT INTO Medication (medication_name, generic_name, form, strength, manufacturer, side_effects) VALUES
    ('L-Dopa', 'Levodopa', 'Tablet', '250mg', 'Iran Pharmaceutical Company', 'Nausea, Insomnia, Headache'),
    ('Ramipril', 'Ramipril', 'Tablet', '5mg', 'Pars Pharma Company', 'Dizziness, Fatigue'),
    ('Topiramate', 'Topiramate', 'Capsule', '100mg', 'Pesa Pharma Company', 'Appetite Loss, Weight Loss'),
    ('Metformin', 'Metformin', 'Tablet', '500mg', 'Artan Pharma Company', 'Diarrhea, Bloating'),
    ('Aspirin', 'Acetylsalicylic Acid', 'Tablet', '325mg', 'Dana Pharma Company', 'Bleeding, Stomach Irritation'),
    ('Amitriptyline', 'Amitriptyline', 'Tablet', '10mg', 'Iris Pharma Company', 'Dry Mouth, Drowsiness'),
    ('Diazepam', 'Diazepam', 'Tablet', '2mg', 'Alborz Pharma Company', 'Sleepiness, Weakness'),
    ('Nifedipine', 'Nifedipine', 'Tablet', '20mg', 'Sina Pharma Company', 'Headache, Facial Flushing'),
    ('Sumatriptan', 'Sumatriptan', 'Injection', '20mg', 'Razak Pharma Company', 'Feeling of Warmth, Fatigue'),
    ('Propranolol', 'Propranolol', 'Tablet', '40mg', 'Arvand Pharma Company', 'Fatigue, Cold Extremities'),
    ('Atorvastatin', 'Atorvastatin', 'Tablet', '20mg', 'Abidi Pharma', 'Muscle pain, Digestive issues');


    INSERT INTO Test_Catalog (test_name, test_category, description, normal_range) VALUES
    ('Brain MRI', 'Imaging', 'Magnetic resonance imaging of the brain', 'No lesions'),
    ('CT Scan', 'Imaging', 'Computed tomography scan of the head', 'No injury'),
    ('EEG', 'Laboratory', 'Electroencephalography for brain electrical activity', 'Normal activity'),
    ('Complete Blood Count', 'Laboratory', 'Assessment and counting of blood components', 'Within normal range'),
    ('MMSE Memory Test', 'Cognitive Test', 'Mini-Mental State Examination for memory and concentration', '≥24 points'),
    ('Attention Test', 'Cognitive Test', 'Assessment of concentration ability', '>80%'),
    ('Cerebrospinal Fluid Test', 'Laboratory', 'Analysis of central nervous system fluid', 'No infection'),
    ('CSF Protein Test', 'Laboratory', 'Measurement of protein in cerebrospinal fluid', '<45 mg/dL'),
    ('PET Scan', 'Imaging', 'Positron emission tomography for metabolic imaging', 'Normal metabolism'),
    ('Blood Pressure Measurement', 'Laboratory', 'Measurement of arterial blood pressure', '120/80 mmHg'),
    ('ECG', 'Cardiology', 'Electrocardiogram recording heart activity', 'Normal Sinus Rhythm');


    INSERT INTO Diagnosis (patient_id, disease_id, doctor_id, diagnosis_date, severity, notes) VALUES
    (15234567, 1, 10000001, '2025-01-10', 'Severe', 'Patient shows advanced stages of Alzheimers disease'),
    (20567890, 2, 10000001, '2025-01-12', 'Moderate', 'Mild to moderate Parkinsons symptoms'),
    (18234123, 3, 10000002, '2025-01-08', 'Mild', 'Epilepsy controlled with medication'),
    (22891234, 4, 10000002, '2024-12-20', 'Severe', 'Active MS with frequent relapses'),
    (16756234, 5, 20000001, '2025-01-05', 'Moderate', 'Chronic migraine resistant to treatment'),
    (23421567, 6, 10000001, '2024-11-15', 'Mild', 'Single febrile seizure episode'),
    (14123890, 7, 30000001, '2025-01-02', 'Moderate', 'Diabetic neuropathy in feet'),
    (15678912, 8, 10000001, '2024-10-30', 'Severe', 'Stroke with neurological deficit'),
    (24567890, 9, 40000001, '2025-01-11', 'Mild', 'Hemophilia Type B'),
    (24567890, 10, 10000002, '2024-12-28', 'Severe', 'Head trauma with brain edema'),
    (30998877, 11, 70000001, '2025-01-25', 'Moderate', 'Stable Angina');

    INSERT INTO Appointment (patient_id, doctor_id, appointment_date, reason, status, blood_pressure, heart_rate, duration_minutes, notes) VALUES
    (15234567, 10000001, '2025-01-15', 'Alzheimers Follow-up', 'Confirmed', '140/85', '72', 30, 'Patient attended with caregiver'),
    (20567890, 10000001, '2025-01-16', 'Parkinsons Assessment', 'Completed', '130/80', '68', 25, 'Improvement in symptoms noted'),
    (18234123, 10000002, '2025-01-17', 'Epilepsy Control Check', 'Confirmed', '120/75', '70', 20, 'No seizure episodes'),
    (22891234, 10000002, '2025-01-18', 'MS Evaluation', 'Completed', '125/78', '75', 35, 'Reduced resilience'),
    (16756234, 20000001, '2025-01-19', 'Migraine Treatment', 'Confirmed', '118/76', '66', 25, 'New medications initiated'),
    (23421567, 10000001, '2025-01-20', 'Follow-up Visit', 'Confirmed', '128/82', '71', 15, 'No seizure recurrence'),
    (14123890, 30000001, '2025-01-21', 'Diabetes Control', 'Completed', '135/88', '73', 20, 'Partial control achieved'),
    (15678912, 10000001, '2025-01-22', 'Recurrent Assessment', 'Confirmed', '142/90', '76', 40, 'Residual neurological deficit'),
    (24567890, 40000001, '2025-01-23', 'Medication Renewal', 'Completed', '115/72', '64', 10, 'Normal'),
    (24567890, 10000002, '2025-01-24', 'Head Trauma Follow-up', 'Confirmed', '138/87', '74', 30, 'Imaging required'),
    (30998877, 70000001, '2025-01-26', 'Cardiac Checkup', 'Scheduled', '150/95', '88', 20, 'First visit');


    INSERT INTO Medical_Test (test_catalog_id, patient_id, test_date, result_value, normal_range, findings, result_status, image_url, interpreted_by) VALUES
    (1, 15234567, '2025-01-10', NULL, 'No lesions', 'Lesions in hippocampal region', 'Abnormal', 'http://hospital.com/mri_001.jpg', 30000001),
    (2, 20567890, '2025-01-12', NULL, 'No injury', 'No injury observed', 'Normal', 'http://hospital.com/ct_001.jpg', 30000002),
    (3, 18234123, '2025-01-08', NULL, 'Normal activity', 'Beta rhythm predominant', 'Normal', 'http://hospital.com/eeg_001.jpg', 10000001),
    (4, 22891234, '2025-01-05', 'WBC: 7.2', 'Within normal range', 'White blood cell count normal', 'Normal', NULL, 60000001),
    (5, 16756234, '2024-12-20', '28', '>=24 points', 'Memory relatively preserved', 'Normal', NULL, 50000001),
    (6, 23421567, '2025-01-02', '85%', '>80%', 'Attention relatively good', 'Normal', NULL, 50000001),
    (7, 14123890, '2024-11-15', NULL, 'No infection', 'Minimal cell count', 'Normal', NULL, 60000001),
    (8, 15678912, '2024-10-30', '32', '<45 mg/dL', 'Limited protein levels', 'Normal', NULL, 60000001),
    (9, 24567890, '2025-01-02', NULL, 'Normal metabolism', 'Decreased metabolism in frontal region', 'Abnormal', 'http://hospital.com/pet_001.jpg', 30000001),
    (10, 24567890, '2025-01-20', '122/80', '120/80 mmHg', 'Blood pressure controlled', 'Normal', NULL, 10000001),
    (11, 30998877, '2025-01-25', NULL, 'Normal Sinus', 'ST Segment Depression', 'Abnormal', 'http://hospital.com/ecg_001.jpg', 70000001);


    INSERT INTO Prescription (patient_id, doctor_id, prescribed_date) VALUES
    (15234567, 10000001, '2025-01-10'),
    (20567890, 10000001, '2025-01-12'),
    (18234123, 10000002, '2025-01-08'),
    (22891234, 10000002, '2024-12-20'),
    (16756234, 20000001, '2025-01-05'),
    (23421567, 10000001, '2024-11-15'),
    (14123890, 30000001, '2025-01-02'),
    (15678912, 10000001, '2024-10-30'),
    (24567890, 40000001, '2025-01-11'),
    (24567890, 10000002, '2024-12-28'),
    (30998877, 70000001, '2025-01-25');

    INSERT INTO Prescription_Medication (prescription_id, medication_id, end_date, frequency, dosage, instructions) VALUES
    (1, 1, '2025-02-10', 'Three times daily', '250mg', 'Before meals'),
    (2, 2, '2025-02-12', 'Twice daily', '5mg', 'Morning and evening'),
    (3, 3, '2025-02-08', 'Once daily', '100mg', 'At night'),
    (4, 4, '2025-01-20', 'Twice daily', '500mg', 'With meals'),
    (5, 5, '2025-02-05', 'Every 8 hours', '325mg', 'As needed'),
    (6, 6, '2025-01-15', 'Once daily', '10mg', 'At night'),
    (7, 7, '2025-02-02', 'Three times daily', '2mg', 'As needed'),
    (8, 8, '2024-11-30', 'Twice daily', '20mg', 'Morning and evening'),
    (9, 9, '2025-02-11', 'Single dose', '20mg', 'Intramuscular injection'),
    (10, 10, '2025-01-28', 'Once daily', '40mg', 'Morning'),
    (11, 11, '2025-03-25', 'Once daily', '20mg', 'At night');

    INSERT INTO Treatment_Plan (patient_id, doctor_id, diagnosis_id, treatment_type, start_date, end_date, session_duration, session_count, status, goals, progress_notes) VALUES
    (15234567, 10000001, 1, 'Pharmacotherapy', '2025-01-10', '2025-04-10', 10, 12, 'Active', 'Reduce disruptive behavior', 'Mild improvement'),
    (20567890, 10000001, 2, 'Physiotherapy', '2025-01-12', '2025-03-12', 45,10, 'Active', 'Improve balance and gait', 'Satisfactory progress'),
    (18234123, 10000002, 3, 'Pharmacotherapy', '2025-01-08', '2025-07-08', 20,5, 'Active', 'Seizure control', 'No incidents'),
    (22891234, 10000002, 4, 'Physiotherapy', '2024-12-20', '2025-03-20', 60,8, 'Active', 'Maintain motor skills', 'Stable'),
    (16756234, 20000001, 5, 'Psychological', '2025-01-05', '2025-04-05', 50,7 ,'Active', 'Reduce frequency and severity', 'Relative improvement'),
    (23421567, 10000001, 6, 'Management', '2024-11-15', '2024-12-15', 30, 9,'Completed', 'Prevent recurrence', 'Successful'),
    (14123890, 30000001, 7, 'Physiotherapy', '2025-01-02', '2025-04-02', 30,11, 'Active', 'Reduce pain and improve sensation', 'Slow improvement'),
    (15678912, 10000001, 8, 'Rehabilitation', '2024-10-30', '2025-03-30', 90,2, 'Active', 'Restore function', 'Limited progress'),
    (24567890, 40000001, 9, 'Pharmacotherapy', '2025-01-11', '2025-07-11', 40,1, 'Active', 'Prevent bleeding', 'Good control'),
    (24567890, 10000002, 10, 'Rehabilitation', '2024-12-28', '2025-03-28', 60,4, 'Active', 'Improve cognitive function', 'Progress initiated'),
    (30998877, 70000001, 11, 'Lifestyle Modification', '2025-01-25', '2025-06-25', 50,8, 'Active', 'Weight loss and diet control', 'Plan initiated');

    INSERT INTO Treatment_Sessions (plan_id, session_number, session_date, status, notes) VALUES
    -- Plan 4: Amanda Rodriguez (MS - Physiotherapy) - Started 2024-12-20
    (4, 1, '2024-12-22', 'Completed', 'Initial mobility assessment and stretching exercises'),
    (4, 2, '2024-12-29', 'Completed', 'Strength training for lower limbs, patient reported fatigue'),
    (4, 3, '2025-01-05', 'Completed', 'Balance coordination drills, slight improvement noted'),
    (4, 4, '2025-01-12', 'Completed', 'Hydrotherapy session, patient responded well'),
    (4, 5, '2025-01-19', 'Cancelled', 'Patient called to cancel due to severe weather'),

    -- Plan 8: William Allen (Stroke - Rehabilitation) - Started 2024-10-30
    (8, 1, '2024-11-05', 'Completed', 'Motor skills recovery for right arm'),
    (8, 2, '2024-12-10', 'Completed', 'Speech therapy integration, progress is slow'),
    (8, 3, '2025-01-15', 'Completed', 'Walking with assistance, gait analysis'),

    -- Plan 6: Rachel Lee (Febrile Seizure - Management) - Started 2024-11-15 (Completed Plan)
    (6, 1, '2024-11-20', 'Completed', 'Parental counseling on fever management and seizure protocols'),

    -- Plan 10: Victoria Young (Head Trauma - Rehab) - Started 2024-12-28
    (10, 1, '2025-01-02', 'Completed', 'Cognitive function baseline test and memory exercises'),
    (10, 2, '2025-01-16', 'Completed', 'Attention span training, patient complained of headache'),

    -- Plan 7: Christopher Walker (Diabetic Neuropathy - Physiotherapy) - Started 2025-01-02
    (7, 1, '2025-01-05', 'Completed', 'Foot sensation testing and low-impact cardio'),
    (7, 2, '2025-01-19', 'Completed', 'Nerve stimulation therapy'),

    -- Plan 5: Marcus Lewis (Migraine - Psychological) - Started 2025-01-05
    (5, 1, '2025-01-07', 'Completed', 'Stress trigger identification session'),
    (5, 2, '2025-01-14', 'Completed', 'CBT techniques for pain management'),
    (5, 3, '2025-01-21', 'Completed', 'Relaxation and breathing exercises'),

    -- Plan 2: Anna Miller (Parkinsons - Physiotherapy) - Started 2025-01-12
    (2, 1, '2025-01-14', 'Completed', 'Posture and balance exercises to prevent falls'),
    (2, 2, '2025-01-21', 'Completed', 'Gait training with rhythmic auditory stimulation'),
    (2, 3, '2025-01-28', 'Scheduled', 'Upcoming session for upper body rigidity'),

    -- Plan 11: Hassan Rezazadeh (Cardiac - Lifestyle) - Started 2025-01-25
    (11, 1, '2025-01-27', 'Scheduled', 'Initial consultation with nutritionist for heart-healthy diet');



    -- PATIENT ID MAPPING:
    -- patient_id: 15234567 -> Alexander Brown (treated by John Smith - 10000001)
    -- patient_id: 20567890 -> Anna Miller (treated by John Smith - 10000001)
    -- patient_id: 18234123 -> Andrew Nelson (treated by Sarah Johnson - 10000002)
    -- patient_id: 22891234 -> Amanda Rodriguez (treated by Sarah Johnson - 10000002)
    -- patient_id: 16756234 -> Marcus Lewis (treated by Michael Davis - 20000001)
    -- patient_id: 23421567 -> Rachel Lee (treated by John Smith - 10000001)
    -- patient_id: 14123890 -> Christopher Walker (treated by Emily Wilson - 30000001)
    -- patient_id: 21345678 -> Sophia Hall
    -- patient_id: 15678912 -> William Allen (treated by John Smith - 10000001)
    -- patient_id: 24567890 -> Victoria Young (treated by James Moore - 40000001 & Sarah Johnson - 10000002)

