CREATE TABLE Doctors(
    First_Name VARCHAR,
    Last_Name VARCHAR,
    Phone VARCHAR,
    Email VARCHAR,
    Username VARCHAR NOT NULL,
    Password VARCHAR NOT NULL,
    PRIMARY KEY (Username)
);
CREATE TABLE Pharmacists(
    First_Name VARCHAR,
    Last_Name VARCHAR,
    Employee_ID VARCHAR NOT NULL,
    Password VARCHAR NOT NULL,
    PRIMARY KEY (Employee_ID)
    Is_Manager BOOLEAN,
);