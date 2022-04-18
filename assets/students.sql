CREATE TABLE `students` (
	`id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` VARCHAR(128) NOT NULL,
	`age` int(11) NOT NULL,
	`sex` TINYINT NOT NULL,
	`phone` VARCHAR(11) NOT NULL,
	`photo` VARCHAR(256) NOT NULL,
	PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;