CREATE DATABASE info;
use info;

CREATE TABLE info(
    ID int not null AUTO_INCREMENT,
    FirstName varchar(100) NOT NULL,
    Surname varchar(100) NOT NULL,
    Company varchar(100), 
    Domicile varchar(100),
    Notes varchar(255),
    Telephone varchar(100),
    Email varchar(100),
    PRIMARY KEY (ID)
);