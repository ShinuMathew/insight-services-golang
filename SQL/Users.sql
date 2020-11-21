Create table USERS(

ID int(255) primary key auto_increment,
USER_ID varchar(255),
Username varchar(255),
User_Password varchar(255),
User_Status enum('Active', 'Inactive')
);

Insert into mylearning.USERS(USER_ID, Username, User_Password, User_Status) value('admin123', 'admin', 'admin', 'Active');
Insert into mylearning.USERS(USER_ID, Username, User_Password, User_Status) value('sneha', 'sneha', 'wilson', 'Inactive');

select User_Status from Users where Username='admin' and User_Password='admin';