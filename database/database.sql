-- Database required for the project to work
CREATE DATABASE IF NOT EXISTS todo;

USE `todo`;

CREATE TABLE IF NOT EXISTS users (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS todos (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `content` varchar(255) NOT NULL,
  `checked` BOOLEAN NOT NULL,
  `user_id` int NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);
