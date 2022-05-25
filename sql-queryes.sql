CREATE TABLE IF NOT EXISTS users_groups(
	id INT GENERATED ALWAYS AS IDENTITY,
	user_group_name VARCHAR(50) NOT NULL,
	PRIMARY KEY(id)
);

INSERT INTO users_groups(user_group_name) VALUES('common');
INSERT INTO users_groups(user_group_name) VALUES('provider');

CREATE TABLE IF NOT EXISTS users(
	id INT GENERATED ALWAYS AS IDENTITY,
	user_group_id INT,
	uid VARCHAR(255),
	user_name VARCHAR(100) NOT NULL,
	last_login VARCHAR(100),
	email VARCHAR(100) UNIQUE,
	phone VARCHAR(50),
	image_url VARCHAR(500),
	created_at VARCHAR(100),
	active BOOLEAN,
	PRIMARY KEY(id),
	CONSTRAINT fk_users_groups
		FOREIGN KEY(user_group_id)
			REFERENCES users_groups(id)
);

CREATE TABLE IF NOT EXISTS hashs(
	id INT GENERATED ALWAYS AS IDENTITY,
	user_id INT,
	hash VARCHAR(100) NOT NULL,
	PRIMARY KEY(id),
	CONSTRAINT fk_users
		FOREIGN KEY(user_id)
			REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS addresses(
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(100),
	postal_code VARCHAR(50),
	state_uf VARCHAR(8),
	city_id INT,
	district VARCHAR(100),
	street VARCHAR(100),
	number INT,
	complement VARCHAR(200),
	reference_point VARCHAR(200),
	latitude VARCHAR(50),
	longitude VARCHAR(50),
	created_at VARCHAR(50),
	PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_addresses(
	id INT GENERATED ALWAYS AS IDENTITY,
	user_id INT,
	address_id int,
	PRIMARY KEY(id),
	CONSTRAINT fk_users
		FOREIGN KEY(user_id)
			REFERENCES users(id),
	CONSTRAINT fk_addresses
		FOREIGN KEY(address_id)
			REFERENCES addresses(id)
);

CREATE TABLE IF NOT EXISTS provider_types(
	id INT GENERATED ALWAYS AS IDENTITY,
	name VARCHAR(100),
	PRIMARY KEY(id)
);

INSERT IGNORE INTO provider_types(name) VALUES('Diarista');
INSERT IGNORE INTO provider_types(name) VALUES('Pedreiro');
INSERT IGNORE INTO provider_types(name) VALUES('Funileiro');
INSERT IGNORE INTO provider_types(name) VALUES('Encanador');
INSERT IGNORE INTO provider_types(name) VALUES('Serralheiro');
INSERT IGNORE INTO provider_types(name) VALUES('Cuidador');

CREATE TABLE IF NOT EXISTS companies(
	id INT GENERATED ALWAYS AS IDENTITY,
	company_name VARCHAR(100),
	user_id INT,
	provider_type_id INT,
	expiration_date VARCHAR(100),
	description VARCHAR(1000),
	created_at VARCHAR(100),
	active BOOLEAN,
	PRIMARY KEY(id),
	CONSTRAINT fk_users
		FOREIGN KEY(user_id)
			REFERENCES users(id),
	CONSTRAINT fk_provider_types
		FOREIGN KEY(provider_type_id)
			REFERENCES provider_types(id)
);

CREATE TABLE IF NOT EXISTS company_addresses(
	id INT GENERATED ALWAYS AS IDENTITY,
	company_id INT,
	address_id int,
	PRIMARY KEY(id),
	CONSTRAINT fk_companies
		FOREIGN KEY(company_id)
			REFERENCES companies(id),
	CONSTRAINT fk_addresses
		FOREIGN KEY(address_id)
			REFERENCES addresses(id)
);