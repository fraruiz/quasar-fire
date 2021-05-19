CREATE TABLE `satellites` (
	`id` CHAR(36),
	`name` VARCHAR(36),
	`x` FLOAT DEFAULT '0',
	`y` FLOAT DEFAULT '0',
	PRIMARY KEY (`id`)
) ENGINE=InnoDB;
