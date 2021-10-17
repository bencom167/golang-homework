DROP TABLE IF exists users;
CREATE TABLE users (
  id int PRIMARY KEY,
  fullname text NOT null,
  email text,
  phone text,
  age int, 
  sex text
);