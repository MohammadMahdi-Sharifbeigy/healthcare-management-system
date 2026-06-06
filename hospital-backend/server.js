const express = require('express');
const sql = require('mssql');
const cors = require('cors');

const app = express();
app.use(cors());
app.use(express.json());

const config = {
    user: 'app_user',    
    password: '12345',    
    server: 'localhost',
    database: 'db1',
    options: {
        trustServerCertificate: true 
    }
};


app.get('/api/medications', async (req, res) => {
    try {
        let pool = await sql.connect(config);
        const result = await pool.request().execute('sp_GetAllMedications');
        res.json(result.recordset);
    } catch (err) {
        console.error('Error fetching medications:', err);
        res.status(500).send(err.message);
    }
});

// در فایل server.js
app.post('/api/add-patient', async (req, res) => {
    try {
        // ایمیل به لیست ورودی‌ها اضافه شد
        const { id, firstName, lastName, dob, gender, email, phone, address, emergency, blood } = req.body;
        
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('patient_id', sql.Int, id)
            .input('first_name', sql.NVarChar, firstName)
            .input('last_name', sql.NVarChar, lastName)
            .input('date_of_birth', sql.Date, dob)
            .input('gender', sql.NVarChar, gender)
            .input('email', sql.NVarChar, email) // <--- ارسال ایمیل به دیتابیس
            .input('phone', sql.NVarChar, phone)
            .input('address', sql.NVarChar, address)
            .input('emergency_contact', sql.NVarChar, emergency)
            .input('blood_type', sql.NVarChar, blood)
            .execute('sp_AddPatient'); // مطمئن شوید پروسیجر sp_AddPatient در SQL ورودی ایمیل را قبول می‌کند
        
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).json({ status: 'Error', message: err.message });
    }
});

app.get('/api/patient/:id', async (req, res) => {
    try {
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('patient_id', sql.Int, req.params.id)
            .execute('sp_GetPatientById');
        
        if (result.recordset.length === 0) {
            return res.status(404).send('Patient not found');
        }
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).send(err.message);
    }
});

app.post('/api/update-patient', async (req, res) => {
    try {
        const { id, phone, address, emergency } = req.body;
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('patient_id', sql.Int, id)
            .input('phone', sql.NVarChar, phone)
            .input('address', sql.NVarChar, address)
            .input('emergency_contact', sql.NVarChar, emergency)
            .execute('sp_UpdatePatient');
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).json({ status: 'Error', message: err.message });
    }
});

app.post('/api/add-appointment', async (req, res) => {
    try {
        const { patientId, doctorId, date, reason } = req.body;
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('patient_id', sql.Int, patientId)
            .input('doctor_id', sql.Int, doctorId)
            .input('appointment_date', sql.Date, date)
            .input('reason', sql.NVarChar, reason)
            .execute('sp_AddAppointment');
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).json({ status: 'Error', message: err.message });
    }
});

app.post('/api/add-prescription', async (req, res) => {
    try {
        const { patientId, doctorId, date, medications } = req.body;
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('patient_id', sql.Int, patientId)
            .input('doctor_id', sql.Int, doctorId)
            .input('prescribed_date', sql.Date, date)
            .input('MedicationsJson', sql.NVarChar(sql.MAX), JSON.stringify(medications))
            .execute('sp_CreatePrescription_JSON');
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).json({ status: 'Error', message: err.message });
    }
});

app.get('/api/prescription/:id', async (req, res) => {
    try {
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('prescription_id', sql.Int, req.params.id)
            .execute('sp_GetPrescriptionById');
        
        if (!result.recordsets[0] || result.recordsets[0].length === 0) {
            return res.status(404).json({ message: 'نسخه یافت نشد' });
        }


        res.json({ 
            header: result.recordsets[0][0], 
            items: result.recordsets[1] 
        });
    } catch (err) {
        console.error(err);
        res.status(500).send(err.message);
    }
});

app.post('/api/update-prescription', async (req, res) => {
    try {
        const { id, patientId, doctorId, date, medications } = req.body;
        let pool = await sql.connect(config);
        const result = await pool.request()
            .input('prescription_id', sql.Int, id)
            .input('patient_id', sql.Int, patientId)
            .input('doctor_id', sql.Int, doctorId)
            .input('prescribed_date', sql.Date, date)
            .input('MedicationsJson', sql.NVarChar(sql.MAX), JSON.stringify(medications))
            .execute('sp_UpdatePrescription_Full');
        res.json(result.recordset[0]);
    } catch (err) {
        res.status(500).json({ status: 'Error', message: err.message });
    }
});

app.get('/api/report/:id', async (req, res) => {
    try {
        const reportProcedures = {
            1: 'sp_PatientMedicalHistory',
            2: 'sp_DoctorPatientsList',
            3: 'sp_AbnormalTestResults',
            4: 'sp_ActiveMedicationsForPatient',
            5: 'sp_TreatmentPlanProgress',
            6: 'sp_CommonDiseasesStatistics',
            7: 'sp_PatientsWithSpecificDisease',
            8: 'sp_DoctorWorkloadReport',
            9: 'sp_ExpiredPrescriptions',
            10: 'sp_DiagnosisAccuracyReport',
            11: 'sp_TestingFacilityWorkload',
            12: 'sp_ActiveTreatmentPlans',
            13: 'sp_SpecialistPerformance',
            14: 'sp_SevereCasesAlert'
        };

        const reportId = req.params.id;
        const paramValue = req.query.param;
        
        const procedureName = reportProcedures[reportId];
        if (!procedureName) return res.status(404).send('Report not found');

        let pool = await sql.connect(config);
        let request = pool.request();

        if (paramValue) {
            if ([1, 4].includes(+reportId)) request.input('patient_id', sql.Int, paramValue);
            else if ([2, 10].includes(+reportId)) request.input('doctor_id', sql.Int, paramValue);
            else if ([5].includes(+reportId)) request.input('plan_id', sql.Int, paramValue);
            else if ([7].includes(+reportId)) request.input('disease_id', sql.Int, paramValue);
        }

        const result = await request.execute(procedureName);
        res.json(result.recordset);

    } catch (err) {
        console.error('Report Error:', err);
        res.status(500).send(err.message);
    }
});

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});