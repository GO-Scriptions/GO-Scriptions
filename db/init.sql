CREATE TABLE Doctors (
    First_Name varchar,
    Last_Name varchar,
    Phone varchar,
    Email varchar,
    Username varchar primary key,
    Doc_Password varchar NOT NULL
    );

CREATE TABLE Pharmacists (
    First_Name varchar,
    Last_Name varchar,
    Employee_ID varchar primary key,
    Pharm_Password varchar NOT NULL,
    Is_Manager varchar);

CREATE TABLE Inventory (
    Drug_Name varchar primary key,
    Amt_On_Hand int,
    Cost_Per_Mg decimal,
    Supplier varchar
);

CREATE TABLE Prescriptions (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Pateient_DOB date,
    Cost decimal,
    Presc_Status varchar,
    Date_Prescribed date,
);

CREATE TABLE Prescription_History (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Patient_DOB date,
    Cost decimal,
    Date_Prescribed date,
);

INSERT INTO Doctors (First_Name, Last_Name, Phone, Email, Username, Doc_Password) values ('Young', 'Farwa' ,'989-989-9898','thefarwacist@gmail.com','drFarwa','theverybest');
INSERT INTO Pharmacists values ('Bruce', 'Banner', '55', 'hulksmash','true');
INSERT INTO Inventory values ('Ibuprofen', '5500', '1.25', 'Meditech');
INSERT INTO Prescriptions values ('459056', 'Young Farwa', 'Amoxicillin', '500', 'Ben', 'Bonson', '2001-01-01', '3.25', '2019 -10-31');