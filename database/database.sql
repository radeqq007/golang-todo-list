-- Database required for the project to work

CREATE DATABASE todo;

USE `todo`;

CREATE TABLE users (
  `id` int PRIMARY KEY,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
);

CREATE TABLE todos (
  `id` int PRIMARY KEY,
  `content` varchar(255) NOT NULL,
  `user_id` int NOT NULL,
  FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
);
