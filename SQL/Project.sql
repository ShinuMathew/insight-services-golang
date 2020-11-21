USE MYLEARNING;

Create table PROJECTS(

Project_ID int(255) primary key auto_increment,
Project_name varchar(255),
Project_desc varchar(255),
Created_Date datetime,
Modified_Date datetime,
Project_status enum('Active', 'Inactive')
);

Drop trigger if exists Create_Project;
Delimiter $$
Create trigger Insert_Project
before insert on PROJECTS
for each row
begin
	set new.Created_Date=now(), new.Modified_Date=now();
end$$
Delimiter ;

Drop trigger if exists Update_Project;
Delimiter $$
Create trigger Update_Project
before update on PROJECTS
for each row
begin
	set new.Modified_Date=now();
end$$
Delimiter ;

insert into PROJECTS(Project_name, Project_desc, Project_status) values('ESUITE','ESUITE- B2C project', 'Active');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('ZTEST', 'ZTEST- Test project', 'Inactive');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('DUMMY','DUMMY- Dummy project', 'Inactive');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('WEBSERVICES', 'WEBSERVICES- API automation', 'Active');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('EBIZ', 'EBIZ- Online agreements/PCP', 'Active');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('SAMWS', 'SAMWS- Customer service app', 'Active');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('ECOMM', 'ECOMM- B2B project', 'Active');
insert into PROJECTS(Project_name, Project_desc, Project_status) values('PIMSC','PIMCS- PIM and Sitecore', 'Inactive');

Select * from PROJECTS;
