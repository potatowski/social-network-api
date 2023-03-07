-- This password is "123456"
INSERT INTO user (name, username, email, password)
values
("user 1", "user_1", "user1@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
("user 2", "user_2", "user2@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy"),
("user 3", "user_3", "user3@gmail.com", "$2a$10$0iGYlKCAYTyJV/vC6nLGgeWFwD6AhSkWLsVRO/.M4lNK8OtIkfggy");

INSERT INTO follower(user_id, follower_id)
values
(1, 2),
(3, 1),
(1, 3);

INSERT INTO post(uuid, title, body, user_id)
VALUES 
("dbc08004-763d-4481-844b-2c23e68c63da", "title 1", "body 1", 1),
("97225a4c-d312-4676-9c16-0456fe5dede2", "title 2", "body 2", 2),
("ad6d4a0c-1e56-40d9-b234-a140d9e88778", "title 3", "body 3", 3);

INSERT INTO post_like(user_id, post_id, `type`)
VALUES
(1, 1, 1),
(2, 1, 1),
(3, 1, 1),
(1, 2, 1),
(2, 2, 1),
(3, 2, 1),
(1, 3, 1),
(2, 3, 1),
(3, 3, 1);

INSERT INTO comment(uuid, body, user_id, post_id)
VALUES
("a0b0c0d0-e1f1-2a3a-4b5b-6c7c8d9e0f1f", "comment 1", 1, 1),
("a0b0c0d0-e1f1-2a3a-4b5b-6c7c8d9e0f2f", "comment 2", 2, 2),
("a0b0c0d0-e1f1-2a3a-4b5b-6c7c8d9e0f3f", "comment 3", 3, 3);