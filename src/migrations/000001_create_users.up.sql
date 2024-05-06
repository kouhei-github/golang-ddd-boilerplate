-- up.sql
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_name` varchar(255),
  `email` varchar(255) NOT NULL UNIQUE,
  `password` varchar(255) NOT NULL,
  `image` varchar(255)
);

CREATE TABLE `user_auths` (
  `id` int NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `password_hash` varchar(255),
  `password_salt` varchar(255),
  `user_id` int,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);