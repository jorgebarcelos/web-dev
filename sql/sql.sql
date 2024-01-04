CREATE DATABASE IF NOT EXISTS webdev;
USE webdev;

DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id int auto_increment primary key,
    user_name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    user_password  varchar(50) not null unique,
    created_at timestamp default current_timestamp()
) ENGINE=INNODB;