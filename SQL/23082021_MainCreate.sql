CREATE DATABASE ticketbotdb;



CREATE TABLE prj_user (
    userid CHARACTER VARYING(32) NOT NULL UNIQUE,
    nameuser CHARACTER VARYING(300) NOT NULL,
    chatid CHARACTER VARYING(32) NOT NULL UNIQUE,

    CONSTRAINT pk_prj_user PRIMARY KEY (userid)
)
WITH(
	OIDS=FALSE
);

ALTER TABLE prj_user OWNER TO postgres;
COMMENT ON TABLE prj_user IS 'Таблица пользователей';
COMMENT ON TABLE prj_user.userid IS 'Идентификатор пользователя';
COMMENT ON TABLE prj_user.nameuser IS 'Имя пользовтаеля в tlg';
COMMENT ON TABLE prj_user.chatid IS 'Id бота';



CREATE TABLE prj_executor(
    executorid CHARACTER VARYING(32) NOT NULL UNIQUE,
    executorname CHARACTER VARYING(32) NOT NULL UNIQUE,
    executorpasword CHARACTER VARYING(32) NOT NULL UNIQUE,

    CONSTRAINT pk_prj_executor PRIMARY KEY (executorid)
)
WITH(
	OIDS=FALSE
);

ALTER TABLE prj_executor OWNER TO postgres;
COMMENT ON TABLE prj_executor IS 'Таблица исполнителей';
COMMENT ON TABLE prj_executor.executorid IS 'Идентификатор исполнителя';
COMMENT ON TABLE prj_executor.executorname IS 'Имя исполнителя';
COMMENT ON TABLE prj_executor.executorpasword IS 'Пароль исполнителя';




CREATE TABLE prj_status(
    statusid CHARACTER VARYING(32) NOT NULL UNIQUE,
    statuscode CHARACTER VARYING(32) NOT NULL,
    statusdescription CHARACTER VARYING(128) NOT NULL,

    CONSTRAINT pk_prj_status PRIMARY KEY (statusid)
);

ALTER TABLE prj_status OWNER TO postgres;
COMMENT ON TABLE prj_status IS 'Таблица статусов';
COMMENT ON TABLE prj_status.statusid IS 'Идентификатор статуса';
COMMENT ON TABLE prj_status.statuscode IS 'Код статуса';
COMMENT ON TABLE prj_status.statusdescription IS 'Описание статуса';





CREATE SEQUENCE rank_number_seq;
CREATE TABLE prj_order(
    orderid CHARACTER VARYING(32) NOT NULL UNIQUE,
    ordernumber INTEGER NOT NULL default nextval('rank_number_seq'),
    orderdescription text,
    statusid CHARACTER VARYING(32) NOT NULL,
    orderstarttime TIMESTAMP WITHOUT TIME ZONE,
    orderstoptime TIMESTAMP WITHOUT TIME ZONE,

    CONSTRAINT pk_prj_order PRIMARY KEY (orderid),
    CONSTRAINT pk_prj_status FOREIGN KEY (statusid)
        REFERENCES prj_status (statusid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT
);


ALTER SEQUENCE rank_number_seq owned by prj_order.ordernumber;

ALTER TABLE prj_order OWNER TO postgres;
COMMENT ON TABLE prj_order IS 'Таблица заказов';
COMMENT ON TABLE prj_order.orderid IS 'Идентификатор заказа';
COMMENT ON TABLE prj_order.ordernumber IS 'Номер заказа';
COMMENT ON TABLE prj_order.orderdescription IS 'Описание заказа';
COMMENT ON TABLE prj_order.statusid IS 'Статус заказа';
COMMENT ON TABLE prj_order.orderstarttime IS 'Время принятия заказ';
COMMENT ON TABLE prj_order.orderstoptime IS 'Номер завершения заказа';





CREATE TABLE link_userorderexecutor(
    linkid CHARACTER VARYING(32) NOT NULL UNIQUE,
    userid CHARACTER VARYING(32) NOT NULL,
    orderid CHARACTER VARYING(32) NOT NULL,
    executorid CHARACTER VARYING(32) NOT NULL,

    CONSTRAINT pk_link_userorderexecutor PRIMARY KEY (linkid),
    CONSTRAINT pk_prj_user FOREIGN KEY (userid)
        REFERENCES prj_user (statusid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT,
            
    CONSTRAINT pk_prj_order FOREIGN KEY (orderid)
        REFERENCES prj_order  (orderid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT,
            
    CONSTRAINT pk_prj_executor FOREIGN KEY (executorid)
        REFERENCES prj_executor (executorid ) MATCH SIMPLE 
        ON UPDATE RESTRICT ON DELETE RESTRICT
);

ALTER TABLE link_userorderexecutor OWNER TO postgres;
COMMENT ON TABLE link_userorderexecutor IS 'Таблица зависимостей заказов';
COMMENT ON TABLE link_userorderexecutor.linkid IS 'Идентификатор зависимости';
COMMENT ON TABLE link_userorderexecutor.userid IS 'Идентификатор пользователя';
COMMENT ON TABLE link_userorderexecutor.orderid IS 'Идентификатор заказа';
COMMENT ON TABLE link_userorderexecutor.executorid IS 'Идентификатор исполнителя';