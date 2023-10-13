from connectionConfig import getMysqlConnection

# connecting to the mysql server
cursor, db = getMysqlConnection()

# Delete old tables
cursor.execute("DROP TABLE IF EXISTS `logs`;")
cursor.execute("DROP TABLE IF EXISTS `projects`;")
cursor.execute("DROP TABLE IF EXISTS `users`;")
               

# Create users table
cursor.execute("CREATE TABLE `users` ( `userid` VARCHAR(13) NOT NULL, `name` VARCHAR(128) NOT NULL, `surname` VARCHAR(128) NOT NULL, `userType` CHAR(1) NOT NULL, `email` VARCHAR(256), PRIMARY KEY (`userid`));")

# Create logs table
cursor.execute("CREATE TABLE `logs` ( `num` INT NOT NULL AUTO_INCREMENT, `userid` VARCHAR(13) NOT NULL, `projectid` INT NOT NULL, `inTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`num`) );")

# Create projects table
cursor.execute("CREATE TABLE `projects` ( `id` INT NOT NULL AUTO_INCREMENT, `name` VARCHAR(128) NOT NULL, `manager` VARCHAR(13) NOT NULL, `description` VARCHAR(1024) NOT NULL, PRIMARY KEY (`id`) );")

# Create foreign keys
cursor.execute("ALTER TABLE `logs` ADD CONSTRAINT `fk_log_users` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`);")
cursor.execute("ALTER TABLE `logs` ADD CONSTRAINT `fk_log_project` FOREIGN KEY (`projectid`) REFERENCES `projects` (`id`);")
cursor.execute("ALTER TABLE `projects` ADD CONSTRAINT `fk_project_manager` FOREIGN KEY (`manager`) REFERENCES `users` (`userid`);")
