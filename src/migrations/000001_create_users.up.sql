-- up.sql
CREATE TABLE users (
   `id` INT PRIMARY KEY AUTO_INCREMENT,
   `user_name` VARCHAR(255),
   `email` VARCHAR(255) NOT NULL UNIQUE,
   `password` VARCHAR(255) NOT NULL,
   `image` VARCHAR(255)
);

CREATE TABLE user_auths (
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `password_hash` VARCHAR(255),
    `password_salt` VARCHAR(255),
    `user_id` INT,
    FOREIGN KEY (`user_id`) REFERENCES users(`id`)
);