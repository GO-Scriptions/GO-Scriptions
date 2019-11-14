CREATE TABLE doctors (
    docUsername varchar primary key,
    firstName varchar,
    lastName varchar,
    cocPassword varchar);

CREATE TABLE Pharmacists (
    First_Name varchar,
    Last_Name varchar,
    Username varchar primary key,
    Pharm_Password varchar,
    Is_Manager varchar);

CREATE TABLE Inventory (
    Drug_Name varchar primary key,
    Amt_On_Hand int,
    Cost_Per_Mg decimal,
    Supplier varchar);

CREATE TABLE Prescriptions (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Presc_Status varchar,
    Date_Prescribed date);

CREATE TABLE Prescription_History (
    Presc_ID varchar,
    Doc_Prescribing varchar,
    Drug_Name varchar,
    Amount int,
    Patient_First varchar,
    Patient_Last varchar,
    Cost decimal,
    Date_Prescribed date);
INSERT INTO doctors (docUsername, firstName, lastName, docPassword) values ('drFarwa', 'Young', 'Farwa', 'thefarwacist');
INSERT INTO doctors values ('PPotts','Pepper', 'Potts', 'mrsironman');
INSERT INTO Pharmacists values ('Bruce', 'Banner', 'MrGreen', 'hulksmash','true');
INSERT INTO Inventory values ('Ibuprofen', '5500', '1.25', 'Meditech');
INSERT INTO Prescriptions values ('459056', 'Young Farwa', 'Amoxicillin', '500', 'Tony', 'Stark', '3.25', '2019 -10-31');
