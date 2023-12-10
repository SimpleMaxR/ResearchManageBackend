create table Laboratories (
    LabID integer not null,
    Name varchar(255) not null,
    DirectorID integer not null,
    OfficeArea double precision not null,
    Address text not null,
    ResearchDirection text,
    primary key (LabID),
    unique (Name)
);

create table Researchers (
    ResearcherID integer not null,
    LabID integer not null,
    ResearchNumber integer not null,
    Name varchar(255) not null,
    Gender varchar(10) not null,
    Title varchar(100) not null,
    Age integer not null,
    EmailAddress varchar(255) not null,
    ResearchDirection text not null,
    Leader boolean not null,
    primary key (ResearcherID),
    foreign key (LabID) references Laboratories
);

create table Offices (
    OfficeID integer not null,
    LabID integer not null,
    Area double precision not null,
    Address text not null,
    ManagerID integer not null,
    primary key (OfficeID),
    foreign key (LabID) references Laboratories,
    foreign key (ManagerID) references Researchers
);

create table Directors (
    DirectorID integer not null,
    LabID integer not null,
    StartDate date not null,
    Term integer not null,
    primary key (DirectorID),
    foreign key (DirectorID) references Researchers,
    foreign key (LabID) references Laboratories
);

create table Secretaries (
    SecretaryID integer not null,
    Name varchar(255) not null,
    Gender varchar(10) not null,
    Age integer not null,
    MobilePhone varchar(20) not null,
    EmailAddress varchar(255) not null,
    primary key (SecretaryID)
);

create table SecretaryServices (
    SecretaryID integer not null,
    LabID integer not null,
    EmploymentDate date not null,
    Responsibilities text not null,
    primary key (SecretaryID, LabID),
    foreign key (SecretaryID) references Secretaries,
    foreign key (LabID) references Laboratories
);

create table Leaders (
    LeaderID integer not null,
    Name varchar(255) not null,
    OfficePhone varchar(20),
    MobilePhone varchar(20),
    EmailAddress varchar(255),
    primary key (LeaderID)
);

create table Clients (
    ClientID integer not null,
    Name varchar(255) not null,
    Address text not null,
    LeaderID integer not null,
    OfficePhone varchar(20) not null,
    primary key (ClientID),
    foreign key (LeaderID) references Leaders
);

comment on column Clients.LeaderID is '委托方负责人';

create table Partners (
    PartnerID integer not null,
    Name varchar(255) not null,
    Address text not null,
    LeaderID integer not null,
    OfficePhone varchar(20) not null,
    primary key (PartnerID),
    foreign key (LeaderID) references Leaders
);

comment on column Partners.LeaderID is '合作方负责人';

create table QualityMonitors (
    MonitorID integer not null,
    Name varchar(255) not null,
    Address text not null,
    LeaderID integer not null,
    primary key (MonitorID),
    foreign key (LeaderID) references Leaders
);

create table Projects (
    ProjectID integer not null,
    PeojectLeader integer not null,
    Name varchar(255) not null,
    ResearchContent text,
    TotalFunds double precision not null,
    StartDate date not null,
    EndDate date not null,
    QualityMonitorsID integer not null,
    ClientID integer not null,
    primary key (ProjectID),
    foreign key (PeojectLeader) references Researchers,
    foreign key (QualityMonitorsID) references QualityMonitors,
    foreign key (ClientID) references Clients
);

create table ProjectPartners (
    ProjectID integer not null,
    PartnerID integer not null,
    primary key (ProjectID, PartnerID),
    foreign key (ProjectID) references Projects,
    foreign key (PartnerID) references Partners
);

create table Contacts (
    ContactID integer not null,
    Name varchar(255) not null,
    OfficePhone varchar(20) not null,
    MobilePhone varchar(20) not null,
    EmailAddress varchar(255) not null,
    BaseClient integer,
    BasePartners integer,
    BaseQM integer,
    primary key (ContactID),
    foreign key (BaseClient) references Clients,
    foreign key (BasePartners) references Partners,
    foreign key (BaseQM) references QualityMonitors
);

create table ProjectResearchers (
    ProjectID integer not null,
    ResearcherID integer not null,
    JoinDate date not null,
    Workload double precision not null,
    DisposableFunds double precision not null,
    primary key (ProjectID, ResearcherID),
    foreign key (ProjectID) references Projects,
    foreign key (ResearcherID) references Researchers
);

create table Subtopics (
    SubtopicID integer not null,
    ProjectID integer not null,
    LeaderID integer not null,
    EndDateRequirement date not null,
    DisposableFunds double precision not null,
    TechnicalIndicators text not null,
    primary key (SubtopicID),
    foreign key (ProjectID) references Projects,
    foreign key (LeaderID) references Researchers
);

comment on column Subtopics.EndDateRequirement is 'DDL日期';

create table Achievements (
    AchievementID integer not null,
    Name varchar(255) not null,
    ObtainedDate date not null,
    ContributorID integer not null,
    BaseProject integer not null,
    BaseSubtopic integer,
    Rank integer not null,
    primary key (AchievementID),
    foreign key (ContributorID) references Researchers,
    foreign key (BaseProject) references Projects,
    foreign key (BaseSubtopic) references Subtopics
);

create table Patents (
    PatentID integer not null,
    Type varchar(50) not null,
    primary key (PatentID),
    foreign key (PatentID) references Achievements
);

create table Papers (
    PaperID integer not null,
    primary key (PaperID),
    foreign key (PaperID) references Achievements
);

create table SoftwareRights (
    SoftwareRightID integer not null,
    primary key (SoftwareRightID),
    foreign key (SoftwareRightID) references Achievements
);

create table ProjectAchievements (
    ProjectID integer not null,
    AchievementID integer not null,
    primary key (ProjectID, AchievementID),
    foreign key (ProjectID) references Projects,
    foreign key (AchievementID) references Achievements
);

create table Users (
    UserID integer not null,
    Username varchar(255) not null,
    Password varchar(255) not null,
    RoleID integer not null,
    primary key (UserID)
);

create table Equipment (
    EquipmentID integer not null,
    Name varchar(255) not null,
    PurchaseDate date,
    Available boolean not null,
    primary key (EquipmentID)
);

create table EquipmentReservations (
    ReservationID integer not null,
    EquipmentID integer,
    ResearcherID integer,
    StartTime timestamp,
    EndTime timestamp,
    primary key (ReservationID),
    foreign key (EquipmentID) references Equipment,
    foreign key (ResearcherID) references Researchers
);