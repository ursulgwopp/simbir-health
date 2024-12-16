CREATE TABLE accounts (
    id SERIAL PRIMARY KEY,
	last_name TEXT DEFAULT '',
	first_name TEXT DEFAULT '',
	username TEXT UNIQUE,
	hash_password TEXT,
	roles TEXT[] DEFAULT '{"User"}',
    is_deleted BOOLEAN DEFAULT FALSE
);

CREATE TABLE blacklist (
    token TEXT
);

INSERT INTO accounts (username, hash_password, roles) VALUES ('admin', '73646b612f31332f3366306f642c402f326c73646f64d033e22ae348aeb5660fc2140aec35850c4da997', ARRAY['Admin']);
INSERT INTO accounts (username, hash_password, roles) VALUES ('manager', '73646b612f31332f3366306f642c402f326c73646f641a8565a9dc72048ba03b4156be3e569f22771f23', ARRAY['Manager']);
INSERT INTO accounts (username, hash_password, roles) VALUES ('doctor', '73646b612f31332f3366306f642c402f326c73646f641f0160076c9f42a157f0a8f0dcc68e02ff69045b', ARRAY['Doctor']);
INSERT INTO accounts (username, hash_password, roles) VALUES ('user', '73646b612f31332f3366306f642c402f326c73646f6412dea96fec20593566ab75692c9949596833adc9', ARRAY['User']);
