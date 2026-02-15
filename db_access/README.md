## [Accessing a relational database](https://go.dev/doc/tutorial/database-access)

This tutorial introduces the basics of accessing a relational database with Go and the database/sql package in its standard library.

The [database/sql](https://pkg.go.dev/database/sql) package you’ll be using includes types and functions for connecting to databases, executing transactions, canceling an operation in progress, and more. For more details on using the package, see [Accessing databases](https://go.dev/doc/database/).

In this tutorial, you’ll create a database, then write code to access the database. Your example project will be a repository of data about vintage jazz records.

In this tutorial, you’ll progress through the following sections:

Create a folder for your code.
Set up a database.
Import the database driver.
Get a database handle and connect.
Query for multiple rows.
Query for a single row.
Add data.

### Set up a database
In this step, you’ll create the database you’ll be working with. You’ll use the CLI for the DBMS itself to create the database and table, as well as to add data.

You’ll be creating a database with data about vintage jazz recordings on vinyl.

The code here uses the [MySQL CLI](https://dev.mysql.com/doc/refman/8.0/en/mysql.html), but most DBMSes have their own CLI with similar features.

1. Open a new command prompt.

At the command line, log into your DBMS, as in the following example for MySQL.
```bash
$ mysql -u root -p
Enter password:
mysql>
```

2. At the mysql command prompt, create a database.

```bash
mysql> create database recordings;
```

3. Change to the database you just created so you can add tables.

```bash
mysql> use recordings;
Database changed
```

4. Execute create-tables.sql script for adding tables.
Into the file, paste the following SQL code, then save the file.

From the mysql command prompt, run the script you just created.

You’ll use the source command in the following form:

```bash
mysql> source /path/to/create-tables.sql

Query OK, 0 rows affected, 1 warning (0.01 sec)

Query OK, 0 rows affected (0.02 sec)

Query OK, 4 rows affected (0.01 sec)
Records: 4  Duplicates: 0  Warnings: 0


mysql> select * from album;
+----+---------------+----------------+-------+
| id | title         | artist         | price |
+----+---------------+----------------+-------+
|  1 | Blue Train    | John Coltrane  | 56.99 |
|  2 | Giant Steps   | John Coltrane  | 63.99 |
|  3 | Jeru          | Gerry Mulligan | 17.99 |
|  4 | Sarah Vaughan | Sarah Vaughan  | 34.98 |
+----+---------------+----------------+-------+
4 rows in set (0.00 sec)


```


5. Create a dedicated MySQL user for your app:
```bash
sudo mysql -u root
```

Then in MySQL:
```bash
CREATE USER 'gouser'@'localhost' IDENTIFIED BY 'gopassword';
CREATE DATABASE recordings;
GRANT ALL PRIVILEGES ON recordings.* TO 'gouser'@'localhost';
FLUSH PRIVILEGES;
EXIT;
```
Then set the env vars:
```bash
export DBUSER=gouser
export DBPASS=gopassword
go run .


Connected!
```

