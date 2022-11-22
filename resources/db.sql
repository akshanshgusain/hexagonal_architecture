CREATE
DATABASE banking;

DROP TABLE IF EXISTS "customers";
CREATE TABLE "customers"
(
    "customer_id"   bigserial PRIMARY KEY,
    "name"          varchar(100) NOT NULL,
    "date_of_birth" date         NOT NULL,
    "city"          varchar(100) NOT NULL,
    "zipcode"       varchar(10)  NOT NULL,
    "status"        int          NOT NULL DEFAULT '1'
);

INSERT INTO "customers"
VALUES (2000, 'Steve', '1978-12-15', 'Delhi', '110075', 1),
       (2001, 'Arian', '1988-05-21', 'Newburgh, NY', '12550', 1),
       (2002, 'Hadley', '1988-04-30', 'Englewood, NJ', '07631', 1),
       (2003, 'Ben', '1988-01-04', 'Manchester, NH', '03102', 0),
       (2004, 'Nina', '1988-05-14', 'Clarkston, MI', '48348', 1),
       (2005, 'Osman', '1988-11-08', 'Hyattsville, MD', '20782', 0);


DROP TABLE IF EXISTS "accounts";
CREATE TABLE "accounts"
(
    "account_id"   int            NOT NULL PRIMARY KEY,
    "customer_id"  int            NOT NULL,
    "opening_date" timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "account_type" varchar(10)    NOT NULL,
    "amount"       decimal(10, 2) NOT NULL,
    "status"       int            NOT NULL DEFAULT '1',
    CONSTRAINT accounts_FK
        FOREIGN KEY (customer_id)
            REFERENCES customers (customer_id)

);
INSERT INTO "accounts"
VALUES (95470, 2000, '2020-08-22 10:20:06', 'saving', 6823.23, 1),
       (95471, 2002, '2020-08-09 10:27:22', 'checking', 3342.96, 1),
       (95472, 2001, '2020-08-09 10:35:22', 'saving', 7000, 1),
       (95473, 2001, '2020-08-09 10:38:22', 'saving', 5861.86, 1);


DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions
(
    transaction_id   int            NOT NULL PRIMARY KEY,
    account_id       int            NOT NULL,
    amount           decimal(10, 2) NOT NULL,
    transaction_type varchar(10)    NOT NULL,
    transaction_date timestamp      NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT transactions_FK
        FOREIGN KEY (account_id)
            REFERENCES accounts (account_id)

);


DROP TABLE IF EXISTS users;
CREATE TABLE users
(
    username    varchar(20) NOT NULL PRIMARY KEY,
    password    varchar(20) NOT NULL,
    role        varchar(20) NOT NULL,
    customer_id int                  DEFAULT NULL,
    created_on  timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP

);
INSERT INTO users
VALUES ('admin', 'abc123', 'admin', NULL, '2020-08-09 10:27:22'),
       ('2001', 'abc123', 'user', 2001, '2020-08-09 10:27:22'),
       ('2000', 'abc123', 'user', 2000, '2020-08-09 10:27:22');

DROP TABLE IF EXISTS refresh_token_store;

CREATE TABLE refresh_token_store
(
    refresh_token varchar(300) NOT NULL PRIMARY KEY,
    created_on    TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);