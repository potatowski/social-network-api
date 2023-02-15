CREATE DATABASE IF NOT EXISTS socialnetwork;

USE socialnetwork;

CREATE TABLE user(
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp()
    removed BOOLEAN DEFAULT false
) ENGINE=INNODB;

CREATE TABLE follower(
    user_id INT NOT NULL,
    FOREING KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    follower_id INT NOT NULL,
    FOREING KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
    created TIMESTAMP DEFAULT current_timestamp()
    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE post(
    uuid VARCHAR(36) NOT NULL PRIMARY KEY,
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    body VARCHAR(255) NOT NULL,
    likes INT DEFAULT 0,
    user_id INT NOT NULL,
    FOREING KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    text VARCHAR(255) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;