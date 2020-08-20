CREATE TABLE users (
	username VARCHAR(255),
    password VARCHAR(255),
    
    PRIMARY KEY (username)
);

CREATE TABLE tasks (
	taskid int AUTO_INCREMENT,
    owner varchar(255),
    content varchar(10000),
    completed boolean DEFAULT false,
    
    PRIMARY KEY (taskid),
    FOREIGN KEY (owner) REFERENCES users(username)
);