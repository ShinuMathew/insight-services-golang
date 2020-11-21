use mylearning;

-- TEST_CASES table created
Drop table if exists Test_cases;
create table TEST_CASES(
	TC_ID int auto_increment primary key,
    TC_NAME varchar(255),
    TC_DESC varchar(255),
    CREATED_DATE datetime,
	MODIFIED_DATE datetime,
    TC_STATUS enum('Active', 'Inactive', 'InProgress', 'Depreciated')
	);

Drop trigger Update_Test_case_Date;

-- Insert trigger
DROP TRIGGER IF EXISTS Update_TCDate_On_Insert;
Delimiter $$
Create trigger Update_TCDate_On_Insert
	Before insert on TEST_CASES 
    for each row
    Begin
		set new.CREATED_DATE=now(), new.MODIFIED_DATE=now();
	end$$
Delimiter ;

-- Update trigger
DROP TRIGGER IF EXISTS Update_TCDate_On_Update;
Delimiter $$
Create trigger  Update_TCDate_On_Update
	Before update on TEST_CASES 
    for each row
    Begin
		set new.MODIFIED_DATE=now();
	end$$
Delimiter ;
	
-- Row insert into TEST_CASES
insert into mylearning.Test_cases(TC_NAME, TC_DESC, TC_STATUS) values('Smoke004', 'Smoke004_ProPay_Payment', 'Active');

update mylearning.Test_cases set TC_DESC = 'Smoke001_CrediCard_Payment';

update Test_cases set TC_NAME='', TC_DESC='', TC_STATUS='';

select * from mylearning.test_cases;

delete from mylearning.test_cases where tc_status ='Active';
