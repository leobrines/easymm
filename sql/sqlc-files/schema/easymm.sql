CREATE TABLE users (
	id SERIAL NOT NULL PRIMARY KEY,
	created_at DATE NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE players (
	id SERIAL NOT NULL,
	steam_id VARCHAR(100) NOT NULL,
	user_id INTEGER NOT NULL,
	PRIMARY KEY (id, steam_id, user_id),
	FOREIGN KEY (user_id) REFERENCES users (id)
);

