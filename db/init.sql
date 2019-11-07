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

INSERT INTO Doctors (First_Name, Last_Name, Phone, Email, Username, Doc_Password) values ('Young', 'Farwa' ,'989-989-9898','thefarwacist@gmail.com','drFarwa','theverybest');
INSERT INTO Pharmacists values ('Bruce', 'Banner', '55', 'hulksmash','true');
