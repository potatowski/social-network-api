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