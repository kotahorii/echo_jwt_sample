DROP TABLE IF EXISTS users;

CREATE TABLE users (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(60) NOT NULL,
  `password` VARCHAR(60) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users (`name`, `password`) VALUES ("test", "test")