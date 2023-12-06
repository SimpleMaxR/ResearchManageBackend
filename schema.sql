-- 创建研究室表
CREATE TABLE Laboratories (
    LabID SERIAL PRIMARY KEY,
    Name VARCHAR(255) UNIQUE NOT NULL,
    -- DirectorID INT,
    OfficeArea FLOAT,
    Address TEXT,
    ResearchDirection TEXT
);

-- 创建办公场地表
CREATE TABLE Offices (
    OfficeID SERIAL PRIMARY KEY,
    LabID INT,
    Area FLOAT,
    Address TEXT,
    FOREIGN KEY (LabID) REFERENCES Laboratories(LabID)
);

-- 创建科研人员表
CREATE TABLE Researchers (
    ResearcherID SERIAL PRIMARY KEY,
    LabID INT,
    Name VARCHAR(255) NOT NULL,
    Gender VARCHAR(10) NOT NULL,
    Title VARCHAR(100) NOT NULL,
    Age INT NOT NULL,
    ResearchDirection TEXT,
    Leader boolean NOT NULL,
    FOREIGN KEY (LabID) REFERENCES Laboratories(LabID)
);

-- 创建研究室主任表
CREATE TABLE Directors (
    DirectorID INT PRIMARY KEY,
    LabID INT,
    StartDate DATE,
    Term INT,
    FOREIGN KEY (DirectorID) REFERENCES Researchers(ResearcherID),
    FOREIGN KEY (LabID) REFERENCES Laboratories(LabID)
);

-- 创建秘书表
CREATE TABLE Secretaries (
    SecretaryID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Gender VARCHAR(10),
    Age INT,
    EmploymentDate DATE,
    Responsibilities TEXT
);

-- 创建秘书服务表
CREATE TABLE SecretaryServices (
    SecretaryID INT,
    LabID INT,
    PRIMARY KEY (SecretaryID, LabID),
    FOREIGN KEY (SecretaryID) REFERENCES Secretaries(SecretaryID),
    FOREIGN KEY (LabID) REFERENCES Laboratories(LabID)
);

-- 创建科研项目表
CREATE TABLE Projects (
    ProjectID SERIAL PRIMARY KEY,
    LeaderID INT,
    Name VARCHAR(255) NOT NULL,
    ResearchContent TEXT,
    TotalFunds FLOAT,
    StartDate DATE,
    EndDate DATE,
    ClientID INT,
--     FOREIGN KEY (ClientID) REFERENCES Clients(ClientID),
    FOREIGN KEY (LeaderID) REFERENCES Researchers(ResearcherID)
);

-- 创建委托方表
CREATE TABLE Clients (
    ClientID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Address TEXT,
    ResponsiblePersonID INT,
    ContactPersonID INT,
    OfficePhone VARCHAR(20),
    MobilePhone VARCHAR(20),
    Email VARCHAR(255)
--     FOREIGN KEY (ResponsiblePersonID) REFERENCES CoWorkers(CoWorkerID)
);

-- 创建合作方表
CREATE TABLE Partners (
    PartnerID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Address TEXT,
    ResponsiblePersonID INT,
    ContactPerson VARCHAR(255),
    OfficePhone VARCHAR(20),
    MobilePhone VARCHAR(20),
    Email VARCHAR(255)
);

-- 创建项目合作方关联表
CREATE TABLE ProjectPartners (
    ProjectID INT,
    PartnerID INT,
    ResponsiblePersonID INT,
    PRIMARY KEY (ProjectID, PartnerID),
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
    FOREIGN KEY (PartnerID) REFERENCES Partners(PartnerID),
    FOREIGN KEY (ResponsiblePersonID) REFERENCES Partners(PartnerID)
);

-- 创建质量监测方表
CREATE TABLE QualityMonitors (
    MonitorID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    Address TEXT,
    ResponsiblePerson VARCHAR(255),
    ContactPersonID INT
);

-- 外部人员表
CREATE TABLE CoWorkers (
    CoWorkerID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    OfficePhone VARCHAR(20),
    MobilePhone VARCHAR(20),
    EmailAddress VARCHAR(255)
);

-- 创建项目参与人员表
CREATE TABLE ProjectResearchers (
    ProjectID INT,
    ResearcherID INT,
    JoinDate DATE,
    Workload FLOAT,
    DisposableFunds FLOAT,
    PRIMARY KEY (ProjectID, ResearcherID),
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
    FOREIGN KEY (ResearcherID) REFERENCES Researchers(ResearcherID)
);

-- 创建子课题表
CREATE TABLE Subtopics (
    SubtopicID SERIAL PRIMARY KEY,
    ProjectID INT,
    LeaderID INT,
    EndDateRequirement DATE,
    DisposableFunds FLOAT,
    TechnicalIndicators TEXT,
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
    FOREIGN KEY (LeaderID) REFERENCES Researchers(ResearcherID)
);

-- 创建科研成果表
CREATE TABLE Achievements (
    AchievementID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    ObtainedDate DATE,
    ContributorID INT,
    Rank INT,
    FOREIGN KEY (ContributorID) REFERENCES Researchers(ResearcherID)
);

-- 创建专利表
CREATE TABLE Patents (
    PatentID INT PRIMARY KEY,
    Type VARCHAR(50),
    FOREIGN KEY (PatentID) REFERENCES Achievements(AchievementID)
);

-- 创建论文表
CREATE TABLE Papers (
    PaperID INT PRIMARY KEY,
    FOREIGN KEY (PaperID) REFERENCES Achievements(AchievementID)
);

-- 创建软件著作权表
CREATE TABLE SoftwareRights (
    SoftwareRightID INT PRIMARY KEY,
    FOREIGN KEY (SoftwareRightID) REFERENCES Achievements(AchievementID)
);

-- 创建项目成果关联表
CREATE TABLE ProjectAchievements (
    ProjectID INT,
    AchievementID INT,
    PRIMARY KEY (ProjectID, AchievementID),
    FOREIGN KEY (ProjectID) REFERENCES Projects(ProjectID),
    FOREIGN KEY (AchievementID) REFERENCES Achievements(AchievementID)
);

-- 创建用户表
CREATE TABLE Users (
    UserID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    Username VARCHAR(255) NOT NULL,
    Password VARCHAR(255) NOT NULL,
    Email VARCHAR(255)
);

-- 创建角色表
CREATE TABLE Roles (
    RoleID SERIAL PRIMARY KEY,
    RoleName VARCHAR(255) NOT NULL
);

-- 创建用户角色关联表
CREATE TABLE UserRoles (
    UserID UUID,
    RoleID INT,
    PRIMARY KEY (UserID, RoleID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID),
    FOREIGN KEY (RoleID) REFERENCES Roles(RoleID)
);

-- 创建设备表
CREATE TABLE Equipment (
    EquipmentID SERIAL PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    PurchaseDate DATE,
    Status VARCHAR(50)
);

-- 创建设备预约表
CREATE TABLE EquipmentReservations (
    ReservationID SERIAL PRIMARY KEY,
    EquipmentID INT,
    ResearcherID INT,
    StartTime TIMESTAMP,
    EndTime TIMESTAMP,
    FOREIGN KEY (EquipmentID) REFERENCES Equipment(EquipmentID),
    FOREIGN KEY (ResearcherID) REFERENCES Researchers(ResearcherID)
);
