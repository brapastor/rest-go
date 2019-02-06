CREATE TABLE users (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	username VARCHAR(64),
	first_name VARCHAR(64),
	last_name VARCHAR(64)
);

INSERT INTO users (username, first_name, last_name) VALUES('brapastor', 'Brayan', 'Pastor');