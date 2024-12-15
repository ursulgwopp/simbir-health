CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
	last_name TEXT,
	first_name TEXT,
	username TEXT UNIQUE,
	hash_password TEXT,
	roles TEXT[] DEFAULT '{}',
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE blacklist (
    token TEXT
);