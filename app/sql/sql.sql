CREATE DATABASE IF NOT EXISTS socialnetwork;

USE socialnetwork;

CREATE TABLE users(
    id int auto_increment primary key,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp()
    removed BOOLEAN DEFAULT false
) ENGINE=INNODB;