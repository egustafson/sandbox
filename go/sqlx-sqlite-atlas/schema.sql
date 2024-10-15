-- Create "users" table
CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar NULL,
  PRIMARY KEY (`id`)
);

-- create "blog_posts" table
CREATE TABLE `blog_posts` (
  `id` int NOT NULL,
  `title` varchar(100) NULL,
  `body` text NULL,
  `author_id` int NULL,
  PRIMARY KEY (`id`),
  CONSTRAINT `author_fk` FOREIGN KEY (`author_id`) REFERENCES `users` (`id`)
 );
