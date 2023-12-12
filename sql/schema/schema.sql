create table laboratories
(
	lab_id integer generated by default as identity,
	name varchar(255) not null,
	office_area double precision not null,
	address text not null,
	research_direction text not null,
	primary key (lab_id),
	unique (name)
);

create table researchers
(
	researcherid integer generated by default as identity,
	lab_id integer not null,
	researcher_number varchar(10) not null,
	name varchar(255) not null,
	gender varchar(10) not null,
	title varchar(100) not null,
	age integer not null,
	emailaddress varchar(255) not null,
	leader boolean not null,
	startdate date,
	term integer,
	researchdirection text not null,
	primary key (researcherid),
	unique (researchnumber),
	foreign key (lab_id) references laboratories
);

create table offices
(
	officeid integer generated by default as identity,
	lab_id integer not null,
	area double precision not null,
	address text not null,
	managerid integer not null,
	primary key (officeid),
	foreign key (lab_id) references laboratories,
	foreign key (managerid) references researchers
);

create table secretaries
(
	secretaryid integer generated by default as identity,
	name varchar(255) not null,
	gender varchar(10) not null,
	age integer not null,
	mobilephone varchar(20) not null,
	emailaddress varchar(255) not null,
	primary key (secretaryid)
);

create table secretaryservices
(
	secretaryid integer,
	lab_id integer not null,
	employmentdate date not null,
	responsibilities text not null,
	primary key (secretaryid, lab_id),
	foreign key (secretaryid) references secretaries,
	foreign key (lab_id) references laboratories
);

create table leaders
(
	leaderid integer generated by default as identity,
	name varchar(255) not null,
	officephone varchar(20),
	mobilephone varchar(20),
	emailaddress varchar(255),
	primary key (leaderid)
);

create table clients
(
	clientid integer generated by default as identity,
	name varchar(255) not null,
	address text not null,
	leaderid integer not null,
	officephone varchar(20) not null,
	primary key (clientid),
	foreign key (leaderid) references leaders
);

comment on column clients.leaderid is '委托方负责人';

create table partners
(
	partnerid integer generated by default as identity,
	name varchar(255) not null,
	address text not null,
	leaderid integer not null,
	officephone varchar(20) not null,
	primary key (partnerid),
	foreign key (leaderid) references leaders
);

comment on column partners.leaderid is '合作方负责人';

create table qualitymonitors
(
	monitorid integer generated by default as identity,
	name varchar(255) not null,
	address text not null,
	leaderid integer not null,
	primary key (monitorid),
	foreign key (leaderid) references leaders
);

create table projects
(
	projectid integer generated by default as identity,
	peojectleader integer not null,
	name varchar(255) not null,
	researchcontent text,
	totalfunds double precision not null,
	startdate date not null,
	enddate date not null,
	qualitymonitorsid integer not null,
	clientid integer not null,
	primary key (projectid),
	foreign key (peojectleader) references researchers,
	foreign key (qualitymonitorsid) references qualitymonitors,
	foreign key (clientid) references clients
);

create table projectpartners
(
	projectid integer not null,
	partnerid integer not null,
	primary key (projectid, partnerid),
	foreign key (projectid) references projects,
	foreign key (partnerid) references partners
);

create table contacts
(
	contactid integer generated by default as identity,
	name varchar(255) not null,
	officephone varchar(20) not null,
	mobilephone varchar(20) not null,
	emailaddress varchar(255) not null,
	baseclient integer,
	basepartners integer,
	baseqm integer,
	primary key (contactid),
	foreign key (baseclient) references clients,
	foreign key (basepartners) references partners,
	foreign key (baseqm) references qualitymonitors
);

create table projectresearchers
(
	projectid integer generated by default as identity,
	researcherid integer not null,
	joindate date not null,
	workload double precision not null,
	disposablefunds double precision not null,
	primary key (projectid, researcherid),
	foreign key (projectid) references projects,
	foreign key (researcherid) references researchers
);

create table subtopics
(
	subtopicid integer generated by default as identity,
	projectid integer not null,
	leaderid integer not null,
	enddaterequirement date not null,
	disposablefunds double precision not null,
	technicalindicators text not null,
	primary key (subtopicid),
	foreign key (projectid) references projects,
	foreign key (leaderid) references researchers
);

comment on column subtopics.enddaterequirement is 'DDL日期';

create table achievements
(
	achievementid integer generated by default as identity,
	name varchar(255) not null,
	obtaineddate date not null,
	contributorid integer not null,
	baseproject integer not null,
	basesubtopic integer,
	rank integer not null,
	primary key (achievementid),
	foreign key (contributorid) references researchers,
	foreign key (baseproject) references projects,
	foreign key (basesubtopic) references subtopics
);

create table patents
(
	patentid integer generated by default as identity,
	type varchar(50) not null,
	primary key (patentid),
	foreign key (patentid) references achievements
);

create table papers
(
	paperid integer generated by default as identity,
	primary key (paperid),
	foreign key (paperid) references achievements
);

create table softwarerights
(
	softwarerightid integer generated by default as identity,
	primary key (softwarerightid),
	foreign key (softwarerightid) references achievements
);

create table projectachievements
(
	projectid integer not null,
	achievementid integer not null,
	primary key (projectid, achievementid),
	foreign key (projectid) references projects,
	foreign key (achievementid) references achievements
);

create table users
(
	userid varchar(10) not null,
	username varchar(255) not null,
	password varchar(255) not null,
	roleid integer not null,
	primary key (username)
);

create table equipment
(
	equipmentid integer generated by default as identity,
	name varchar(255) not null,
	purchasedate date,
	available boolean not null,
	primary key (equipmentid)
);

create table equipmentreservations
(
	reservationid integer generated by default as identity,
	equipmentid integer,
	researcherid integer,
	starttime timestamp,
	endtime timestamp,
	primary key (reservationid),
	foreign key (equipmentid) references equipment,
	foreign key (researcherid) references researchers
);

