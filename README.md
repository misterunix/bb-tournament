# bb-tournament
Tournament system for basic bots

**Many things are hard coded at the moment**



CREATE TABLE robots (ID INTEGER NOT NULL PRIMARY KEY,OwnerID INTEGER,Filename TEXT,FilenameHash TEXT,Code TEXT,CodeHash TEXT,Count INTEGER,Points REAL,Win INTEGER,Tie INTEGER,Loss INTEGER);


CREATE TABLE match2 (ID INTEGER NOT NULL PRIMARY KEY,r1 INTEGER,r2 INTEGER,w1 INTEGER,w2 INTEGER,t1 INTEGER,t2 INTEGER,l1 INTEGER,l2 INTEGER,p1 REAL,p2 REAL)

CREATE TABLE match3 (ID INTEGER NOT NULL PRIMARY KEY,r1 INTEGER,r2 INTEGER,r3 INTEGER,w1 INTEGER,w2 INTEGER,w3 INTEGER,t1 INTEGER,t2 INTEGER,t3 INTEGER,l1 INTEGER,l2 INTEGER,l3 INTEGER,p1 REAl,p2 REAL,p3 REAL)

CREATE TABLE match4 (ID INTEGER NOT NULL PRIMARY KEY,r1 INTEGER,r2 INTEGER,r3 INTEGER,r4 INTEGER,w1 INTEGER,w2 INTEGER,w3 INTEGER,w4 INTEGER,t1 INTEGER,t2 INTEGER,t3 INTEGER,t4 INTEGER,l1 INTEGER,l2 INTEGER,l3 INTEGER,l4 INTEGER,p1 REAl,p2 REAL,p3 REAL,p4 REAL)

