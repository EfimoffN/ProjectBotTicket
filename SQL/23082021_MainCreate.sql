CREATE DATABASE ticketbotdb;


CREATE TABLE prj_user (
    userid CHARACTER VARYING(32) NOT NULL PRIMARY KEY,
    nameuser CHARACTER VARYING(300) NOT NULL UNIQUE,
    chatid CHARACTER VARYING(32) NOT NULL
);

CREATE TABLE prj_executor(
    executorid CHARACTER VARYING(32) NOT NULL PRIMARY KEY,
    executorname CHARACTER VARYING(32) NOT NULL UNIQUE,
    executorpasword CHARACTER VARYING(32) NOT NULL UNIQUE
);

CREATE TABLE prj_status(
    statusid CHARACTER VARYING(32) NOT NULL PRIMARY KEY,
    statuscode CHARACTER VARYING(32) NOT NULL,
    statusdescription CHARACTER VARYING(128) NOT NULL
);

CREATE TABLE prj_order(
    orderid CHARACTER VARYING(32) NOT NULL UNIQUE,
    ordernumber SERIAL,
    orderdescription text,
    statusid CHARACTER VARYING(32) NOT NULL,
    orderstarttime TIMESTAMP WITHOUT TIME ZONE,
    orderstoptime TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_prj_order PRIMARY KEY (orderid),
    CONSTRAINT pk_prj_status FOREIGN KEY (statusid)
        REFERENCES prj_status (statusid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT
);

CREATE TABLE link_userorderexecutor(
    linkid CHARACTER VARYING(32) NOT NULL UNIQUE,
    userid CHARACTER VARYING(32) NOT NULL,
    orderid CHARACTER VARYING(32) NOT NULL,
    executorid CHARACTER VARYING(32) NOT NULL,

    CONSTRAINT pk_link_userorderexecutor PRIMARY KEY (linkid),
    CONSTRAINT pk_prj_user FOREIGN KEY (userid)
        REFERENCES prj_user (userid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT,
            
    CONSTRAINT pk_prj_order FOREIGN KEY (orderid)
        REFERENCES prj_order  (orderid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT,
            
    CONSTRAINT pk_prj_executor FOREIGN KEY (executorid)
        REFERENCES prj_executor (executorid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT
);
