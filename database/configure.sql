
CREATE TABLE `bitacoraDB`.`users` (
    `userid` VARCHAR(13) NOT NULL,
    `name` VARCHAR(128) NOT NULL,
    `surname` VARCHAR(128) NOT NULL,
    `userType` CHAR(1) NOT NULL,
    `email` VARCHAR(256),
    PRIMARY KEY (`userid`)
);

CREATE TABLE `bitacoraDB`.`logs` (
    `num` INT NOT NULL AUTO_INCREMENT,
    `userid` VARCHAR(13) NOT NULL,
    `projectid` INT NOT NULL,
    `day` DATE NOT NULL,
    `inTime` TIME NOT NULL,
    `outTime` TIME,
    PRIMARY KEY (`num`)
);

CREATE TABLE `bitacoraDB`.`projects` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(128) NOT NULL,
    `manager` VARCHAR(13) NOT NULL,
    `description` VARCHAR(1024) NOT NULL,
    PRIMARY KEY (`id`)
);

ALTER TABLE `bitacoraDB`.`logs` ADD CONSTRAINT `fk_log_users` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`);
ALTER TABLE `bitacoraDB`.`logs` ADD CONSTRAINT `fk_log_project` FOREIGN KEY (`projectid`) REFERENCES `bitacoraDB`.`projects` (`id`);
