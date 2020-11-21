-- Get Usernames --
SELECT usename
FROM pg_user;

-- Get Databases --
SELECT datname
FROM pg_database;

-- Create Database Insight_Demo --
CREATE DATABASE Insight_Demo;

-- Create User Insight_Admin --
create user Insight_Admin password 'shinz9474';

-- Create Table Test_Run --
Create table Test_Run (
TestRunID int PRIMARY KEY,
config JSON NOT NULL,
Created_Date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
Modified_date timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Create Trigger for table Test_Run to update modified date --
Create or Replace function update_modified_date()
	returns trigger as
$$	
BEGIN
	new.modified_date := now();
	Return new;
END;
$$
LANGUAGE plpgsql;
-- LANGUAGE plpgsql VOLATILE -- Says the function is implemented in the plpgsql language; VOLATILE says the function has side effects.
-- COST 100;

CREATE TRIGGER update_modified_date
BEFORE UPDATE on test_run
FOR EACH ROW EXECUTE PROCEDURE update_modified_date()

-- Seed data for Test_run --
insert into Test_Run values (11221, '{"tc_name":"Smoke_001", "browsers":[ "Chrome" ], "is_parallelexecution":"false"," is_crossbrowser":"false"}');
insert into Test_Run values (11222, '{"tc_name":"Smoke_002", "browsers":[ "Chrome" ], "is_parallelexecution":"false", "is_crossbrowser":"false"}');
insert into Test_Run values (11223, '{"tc_name":"Smoke_003", "browsers":[ "Chrome" ], "is_parallelexecution":"false", "is_crossbrowser":"false"}');
insert into Test_Run values (11224, '{"tc_name":"Smoke_004", "browsers":[ "Chrome" ], "is_parallelexecution":"false", "is_crossbrowser":"false"}');