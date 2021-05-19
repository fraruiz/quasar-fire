CREATE TABLE `satellites` (
	`id` CHAR(36),
	`name` VARCHAR(36),
	`x` FLOAT DEFAULT '0',
	`y` FLOAT DEFAULT '0',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `messages` (
	`id` CHAR(36),
	`content` VARCHAR(36),
	`distance` FLOAT DEFAULT '0',
	`satelite_id` CHAR(36),
	PRIMARY KEY (`id`),
	FOREIGN KEY (`satelite_id`) REFERENCES `satellites`(`id`)
) ENGINE=InnoDB;

