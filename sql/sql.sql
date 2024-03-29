CREATE DATABASE IF NOT EXISTS webdev;

USE webdev;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
  id INT auto_increment PRIMARY KEY,
  user_name VARCHAR (50) NOT NULL,
  nick VARCHAR (50) NOT NULL UNIQUE,
  email VARCHAR (50) NOT NULL UNIQUE,
  user_password VARCHAR (200) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ()
) engine = innodb;