CREATE TABLE Doctors(
    First_Name varchar,
    Last_Name varchar,
    Phone varchar,
    Email varchar,
    Username varchar primary key,
    Password varchar NOT NULL);
CREATE TABLE Pharmacists(
    First_Name varchar,
    Last_Name varchar,
    Employee_ID varchar primary key,
    Password varchar NOT NULL,
    Is_Manager varchar);
INSERT INTO Doctors VALUES('Young', 'Farwa' ,'989-989-9898','thefarwacist@gmail.com','theverybest');
INSERT INTO Pharmacists('Bruce', 'Banner', '55', 'true');
