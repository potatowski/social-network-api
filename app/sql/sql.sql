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
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    follower_id INT NOT NULL,
    FOREIGN KEY (follower_id) REFERENCES user(id) ON DELETE CASCADE,
    created TIMESTAMP DEFAULT current_timestamp()
    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;

CREATE TABLE post(
    uuid VARCHAR(36) NOT NULL UNIQUE,
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(50) NOT NULL,
    body VARCHAR(255) NOT NULL,
    user_id INT NOT NULL,
    removed BOOLEAN DEFAULT false,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    created TIMESTAMP DEFAULT current_timestamp()
) ENGINE=INNODB;
CREATE INDEX idx_post_uuid ON post (uuid);

CREATE TABLE post_like(
    user_id INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id) ON DELETE CASCADE,
    post_id INT NOT NULL,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
    `type` INT NOT NULL DEFAULT 1,
    PRIMARY KEY (user_id, post_id)
) ENGINE=INNODB;