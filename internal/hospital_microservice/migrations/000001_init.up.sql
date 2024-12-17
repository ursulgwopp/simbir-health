CREATE TABLE hospitals (
    id SERIAL PRIMARY KEY,
	name TEXT,
	address TEXT,
	contact_phone TEXT,
	rooms TEXT[],
    is_deleted BOOLEAN DEFAULT FALSE
);