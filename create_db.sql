-- Created by Vertabelo (http://vertabelo.com)
-- Last modification date: 2018-03-23 18:29:58.1

-- tables
-- Table: Customer
CREATE TABLE Customer (
    id int  NOT NULL,
    name varchar(20)  NOT NULL,
    last_name varchar(20)  NOT NULL,
    email varchar(20)  NOT NULL,
    phone int  NOT NULL,
    CONSTRAINT Customer_pk PRIMARY KEY (id)
);

-- Table: Deliver_List
CREATE TABLE Deliver_List (
    id int  NOT NULL,
    Truck_id int  NOT NULL,
    CONSTRAINT Deliver_List_pk PRIMARY KEY (id)
);

-- Table: Job
CREATE TABLE Job (
    id int  NOT NULL,
    job_name int  NOT NULL,
    CONSTRAINT Job_pk PRIMARY KEY (id)
);

-- Table: Package
CREATE TABLE Package (
    id int  NOT NULL,
    description varchar(200)  NOT NULL,
    Customer_id int  NOT NULL,
    CONSTRAINT Package_pk PRIMARY KEY (id)
);

-- Table: State_Package
CREATE TABLE State_Package (
    id int  NOT NULL,
    Package_id int  NOT NULL,
    Deliver_List_id int  NOT NULL,
    CONSTRAINT State_Package_pk PRIMARY KEY (id)
);

-- Table: Truck
CREATE TABLE Truck (
    id int  NOT NULL,
    truck_number varchar(6)  NOT NULL,
    fk_type_truck int  NULL,
    Type_Truck_id int  NOT NULL,
    CONSTRAINT id PRIMARY KEY (id)
);

-- Table: Type
CREATE TABLE Type (
    id int  NOT NULL,
    type_name varchar(30)  NOT NULL,
    Type_Truck_id int  NOT NULL,
    CONSTRAINT Type_pk PRIMARY KEY (id)
);

-- Table: Type_Package
CREATE TABLE Type_Package (
    id int  NOT NULL,
    Type_id int  NOT NULL,
    Package_id int  NOT NULL,
    CONSTRAINT Type_Package_pk PRIMARY KEY (id)
);

-- Table: Type_Truck
CREATE TABLE Type_Truck (
    type_name varchar(10)  NOT NULL,
    id int  NOT NULL,
    CONSTRAINT Type_Truck_pk PRIMARY KEY (id)
);

-- Table: Worker
CREATE TABLE Worker (
    id int  NOT NULL,
    name varchar(20)  NOT NULL,
    last_name varchar(20)  NOT NULL,
    born date  NOT NULL,
    Job_id int  NOT NULL,
    Truck_id int  NOT NULL,
    CONSTRAINT Worker_pk PRIMARY KEY (id)
);

-- foreign keys
-- Reference: Deliver_List_Truck (table: Deliver_List)
ALTER TABLE Deliver_List ADD CONSTRAINT Deliver_List_Truck
    FOREIGN KEY (Truck_id)
    REFERENCES trucks (id)
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Package_Customer (table: Package)
ALTER TABLE Package ADD CONSTRAINT Package_Customer
    FOREIGN KEY (Customer_id)
    REFERENCES Customer (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: State_Package_Deliver_List (table: State_Package)
ALTER TABLE State_Package ADD CONSTRAINT State_Package_Deliver_List
    FOREIGN KEY (Deliver_List_id)
    REFERENCES Deliver_List (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: State_Package_Package (table: State_Package)
ALTER TABLE State_Package ADD CONSTRAINT State_Package_Package
    FOREIGN KEY (Package_id)
    REFERENCES Package (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Truck_Type_Truck (table: Truck)
ALTER TABLE trucks ADD CONSTRAINT Truck_Type_Truck
    FOREIGN KEY (Type_Truck_id)
    REFERENCES type_trucks (id)
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Type_Package_Package (table: Type_Package)
ALTER TABLE Type_Package ADD CONSTRAINT Type_Package_Package
    FOREIGN KEY (Package_id)
    REFERENCES Package (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Type_Package_Type (table: Type_Package)
ALTER TABLE Type_Package ADD CONSTRAINT Type_Package_Type
    FOREIGN KEY (Type_id)
    REFERENCES Type (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Type_Type_Truck (table: Type)
ALTER TABLE Type ADD CONSTRAINT Type_Type_Truck
    FOREIGN KEY (Type_Truck_id)
    REFERENCES type_trucks (id)
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Worker_Job (table: Worker)
ALTER TABLE workers ADD CONSTRAINT Worker_Job
    FOREIGN KEY (Job_id)
    REFERENCES Job (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: Worker_Truck (table: Worker)
ALTER TABLE workers ADD CONSTRAINT Worker_Truck
    FOREIGN KEY (Truck_id)
    REFERENCES trucks (id)
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- End of file.


