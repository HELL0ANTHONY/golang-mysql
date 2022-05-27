-- new database
CREATE DATABASE go_users;

-- create employees table
CREATE TABLE employees(
  id INT NOT NULL AUTO_INCREMENT, 
  name VARCHAR(30) NOT NULL,
  email VARCHAR(100) NOT NULL,
  PRIMARY KEY (id)
);

-- insert data to table
INSERT INTO employees (name, email) VALUES
  ("Maria", "maria@email.com"),
  ("Jose", "jose@email.com");
