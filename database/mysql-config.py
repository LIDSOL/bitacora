# Python implementation to create a Database in MySQL
import mysql.connector

# connecting to the mysql server

db = mysql.connector.connect(
    host = "127.0.0.1",
    user = "bitacoraU",
    passwd = "test-passwd",
    database= "bitacoraDB"
)

cursor = db.cursor()

# Delete old tables
cursor.execute("DROP TABLE IF EXISTS `bitacoraDB`.`logs`;")
cursor.execute("DROP TABLE IF EXISTS `bitacoraDB`.`users`;")
cursor.execute("DROP TABLE IF EXISTS `bitacoraDB`.`projects`;")
               

# Create users table
cursor.execute("CREATE TABLE `bitacoraDB`.`users` ( `userid` VARCHAR(13) NOT NULL, `name` VARCHAR(128) NOT NULL, `surname` VARCHAR(128) NOT NULL, `userType` CHAR(1) NOT NULL, `email` VARCHAR(256), PRIMARY KEY (`userid`));")

# Create logs table
cursor.execute("CREATE TABLE `bitacoraDB`.`logs` ( `num` INT NOT NULL AUTO_INCREMENT, `userid` VARCHAR(13) NOT NULL, `projectid` INT NOT NULL, `inTime` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP, PRIMARY KEY (`num`) );")

# Create projects table
cursor.execute("CREATE TABLE `bitacoraDB`.`projects` ( `id` INT NOT NULL AUTO_INCREMENT, `name` VARCHAR(128) NOT NULL, `manager` VARCHAR(13) NOT NULL, `description` VARCHAR(1024) NOT NULL, PRIMARY KEY (`id`) );")

# Create foreign keys
cursor.execute("ALTER TABLE `bitacoraDB`.`logs` ADD CONSTRAINT `fk_log_users` FOREIGN KEY (`userid`) REFERENCES `users` (`userid`);")
cursor.execute("ALTER TABLE `bitacoraDB`.`logs` ADD CONSTRAINT `fk_log_project` FOREIGN KEY (`projectid`) REFERENCES `bitacoraDB`.`projects` (`id`);")
