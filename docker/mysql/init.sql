CREATE DATABASE IF NOT EXISTS socialnetwork;

USE socialnetwork;

CREATE TABLE user(
    id int AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    created TIMESTAMP DEFAULT current_timestamp(),
    removed BOOLEAN DEFAULT 0
) ENGINE=INNODB;

CREATE TABLE follower(
    user_id INT NOT NULL,
    follower_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
    created TIMESTAMP DEFAULT current_timestamp(),
    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;