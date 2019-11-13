create table doctors (
	username text primary key,
	password text not null
);

create table pharmacists (
	username text primary key,
	password text not null
);

insert into doctors (username, password) values ('Ben','abc');
insert into doctors (username, password) values ('Farwa','abc');

insert into pharmacists (username, password) values ('Seth','abc');
insert into pharmacists (username, password) values ('Tony','abc');
