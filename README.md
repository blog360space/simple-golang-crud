# Simple Golang CRUD

In this project is about to learn how to do database CRUD operations using Golang and MySQL.
CRUD is an acronym for Create, Read, Update, and Delete.
CRUD operations are basic data manipulation for database.

## Step 1: Prepare and Import MySQL driver into your project
Using Git Bash first install driver for Go's MySQL database package. Run below command and install MySQL driver's

```
go get -u github.com/go-sql-driver/mysql
```

Now create Goblog Database

1. Open PHPMyAdmin/SQLyog or what ever MySQL database management tool that you are using.
2. Create a new database "go_employee"

## Step 2: Creating the Employee Table

```
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `city` varchar(30) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
```

## Step 3: Clone project and try.
Change directory to this project and run command below to start project.
```
go run employee.go
```